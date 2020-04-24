package utils

import "github.com/Devatoria/thief/cgroup"

// AppendCgroups creates a slice of the enabled cgroups
func AppendCgroups(blkio, cpu, cpuset, devices, freezer, hugetlb, memory, net, perfevent, pids bool) []cgroup.Kind {
	kinds := []cgroup.Kind{}

	if blkio {
		kinds = append(kinds, cgroup.Blkio)
	}
	if cpu {
		kinds = append(kinds, cgroup.CPU)
	}
	if cpuset {
		kinds = append(kinds, cgroup.CPUSet)
	}
	if devices {
		kinds = append(kinds, cgroup.Devices)
	}
	if freezer {
		kinds = append(kinds, cgroup.Freezer)
	}
	if hugetlb {
		kinds = append(kinds, cgroup.HugeTLB)
	}
	if memory {
		kinds = append(kinds, cgroup.Memory)
	}
	if net {
		kinds = append(kinds, cgroup.Net)
	}
	if perfevent {
		kinds = append(kinds, cgroup.PerfEvent)
	}
	if pids {
		kinds = append(kinds, cgroup.PIDs)
	}

	return kinds
}
