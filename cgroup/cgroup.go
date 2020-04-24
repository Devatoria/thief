package cgroup

// Kind represents a cgroup name
type Kind string

const (
	sysfs = "/sys/fs/cgroup"

	// kinds
	Blkio     Kind = "blkio"
	CPU       Kind = "cpu,cpuacct"
	CPUSet    Kind = "cpuset"
	Devices   Kind = "devices"
	Freezer   Kind = "freezer"
	HugeTLB   Kind = "hugetlb"
	Memory    Kind = "memory"
	Net       Kind = "net_cls,net_prio"
	PerfEvent Kind = "perf_event"
	PIDs      Kind = "pids"
)

var (
	all = []Kind{
		Blkio,
		CPU,
		CPUSet,
		Devices,
		Freezer,
		HugeTLB,
		Memory,
		Net,
		PerfEvent,
		PIDs,
	}
)
