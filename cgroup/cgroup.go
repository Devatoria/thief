package cgroup

// Kind represents a cgroup name
type Kind string

const (
	sysfs      = "/sys/fs/cgroup"
	CPU   Kind = "cpu,cpuacct"
)

var (
	all = []Kind{
		CPU,
	}
)
