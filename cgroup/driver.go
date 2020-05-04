package cgroup

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

// Driver is a cgroup driver to join or execute
// into a given container's cgroup
type Driver interface {
	Join(path string, cgroups []Kind) error
	Run(path string, cgroups []Kind, command string, args []string) error
	Attach(path string, cgroups []Kind, pid int) error
	Exit() error
}

type driver struct {
	sysfs string
}

// Config is the cgroup driver configuration
type Config struct {
	Sysfs string
}

// NewDriver returns a cgroup driver from the given configuration
func NewDriver(config Config) Driver {
	return driver{
		sysfs: config.Sysfs,
	}
}

// getFullPath returns the cgroup.procs file full path
// according to the given sysfs base and cgroup name
func (d driver) getFullPath(base, path string, cg Kind) string {
	return fmt.Sprintf("%s/%s/%s/cgroup.procs", base, cg, strings.TrimPrefix(path, "/"))
}

// writeCgroupFile writes the given pid into the given cgroup file
func (d driver) writeCgroupFile(base, path string, cg Kind, pid int) error {
	fullPath := d.getFullPath(base, path, cg)

	// open given cgroup file
	file, err := os.OpenFile(fullPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error opening cgroup procs file (%s): %w", fullPath, err)
	}

	// write pid into file
	if _, err := file.WriteString(strconv.Itoa(pid)); err != nil {
		return fmt.Errorf("error writing to cgroup procs file (%s): %w", fullPath, err)
	}

	// force sync
	if err := file.Sync(); err != nil {
		return fmt.Errorf("error syncing cgroup procs file written content (%s): %w", fullPath, err)
	}

	// close the file explicitly here
	if err := file.Close(); err != nil {
		return fmt.Errorf("error closing cgroup procs file (%s): %w", fullPath, err)
	}

	return nil
}

// join adds the given pid to the given cgroups procs files
func (d driver) join(base, path string, cgroups []Kind, pid int) error {
	for _, cgroup := range cgroups {
		if err := d.writeCgroupFile(base, path, cgroup, pid); err != nil {
			return fmt.Errorf("error joining %s cgroup: %w", cgroup, err)
		}
	}

	return nil
}

// Join adds the thief process parent's pid to the
// container's cgroup
func (d driver) Join(path string, cgroups []Kind) error {
	// get parent pid
	ppid := os.Getppid()

	// join enabled cgroups
	return d.join(d.sysfs, path, cgroups, ppid)
}

// Run adds the thief process to the
// container's cgroup and then executes the given command
func (d driver) Run(path string, cgroups []Kind, command string, args []string) error {
	// get thread group ID
	pgid, err := syscall.Getpgid(os.Getpid())
	if err != nil {
		return fmt.Errorf("error getting thread group ID: %w", err)
	}

	// join enabled cgroups
	if err := d.join(d.sysfs, path, cgroups, pgid); err != nil {
		return err
	}

	// execute the given command
	cmd := exec.Command(command, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error executing command %s: %w", command, err)
	}

	return nil
}

// Attach attaches the given PID to the given cgroups
func (d driver) Attach(path string, cgroups []Kind, pid int) error {
	return d.join(d.sysfs, path, cgroups, pid)
}

// Exit re-adds the thief process parent's pid to the
// system base sysfs cgroups, for all cgroups
func (d driver) Exit() error {
	// get parent pid
	ppid := os.Getppid()

	// parse PID 1 cgroup file to retrieve base cgroups paths
	// and join them back
	for _, kind := range all {
		path, err := getCgroupPath(1, kind)
		if err != nil {
			return fmt.Errorf("error getting PID 1 cgroup path: %w", err)
		}

		if err := d.join(d.sysfs, path, []Kind{kind}, ppid); err != nil {
			return fmt.Errorf("error joining PID 1 cgroup path: %w", err)
		}
	}

	return nil
}
