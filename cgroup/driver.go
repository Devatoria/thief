package cgroup

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Driver is a cgroup driver to join or execute
// into a given container's cgroup
type Driver interface {
	Join(path string, cgroups []Kind) error
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
	return fmt.Sprintf("%s/%s/%s/tasks", base, cg, strings.TrimPrefix(path, "/"))
}

// writeCgroupFile writes the given pid into the given cgroup file
func (d driver) writeCgroupFile(base, path string, cg Kind, pid int) error {
	fullPath := d.getFullPath(base, path, cg)

	// open given cgroup file
	file, err := os.OpenFile(fullPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error opening cgroup procs file (%s): %w", fullPath, err)
	}

	defer file.Close()

	// write pid into file
	if _, err := file.WriteString(strconv.Itoa(pid)); err != nil {
		return fmt.Errorf("error writing to cgroup procs file (%s): %w", fullPath, err)
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
