package cgroup

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	procs = "cgroup.procs"
)

type Driver interface {
	Enter() error
}

type cgroup struct {
	sysfs string
	path  string
	cpu   bool
}

func NewDriver(sysfs, path string, cpu bool) Driver {
	return cgroup{
		sysfs: sysfs,
		path:  path,
		cpu:   cpu,
	}
}

func (c cgroup) Enter() error {
	// get parent pid
	ppid := os.Getppid()

	// write parent pid into file
	path := fmt.Sprintf("%s/cpu/%s/%s", c.sysfs, strings.TrimPrefix(c.path, "/"), procs)
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error opening cgroup procs file (%s): %w", path, err)
	}

	defer file.Close()

	if _, err := file.WriteString(strconv.Itoa(ppid)); err != nil {
		return fmt.Errorf("error writing to cgroup procs file (%s): %w", path, err)
	}

	return nil
}
