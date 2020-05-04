package main

import (
	"fmt"
	"os"

	"github.com/Devatoria/thief/container"
	"github.com/spf13/cobra"
)

var joinCmd = &cobra.Command{
	Use:     "join <container ID>",
	Short:   "joins the given container cgroups",
	Example: "thief --runtime containerd --socket /run/containerd/containerd.sock join --cpu abcdef123456",
	Args:    cobra.ExactArgs(1),
	PreRun:  initCgroups,
	Run: func(cmd *cobra.Command, args []string) {
		// get args
		id := args[0]

		// create container
		c, err := container.New(cRuntime, socket, id)
		if err != nil {
			fmt.Printf("Error creating container socket: %v\n", err)
			os.Exit(1)
		}

		// retrieve cgroup path
		path, err := c.CgroupPath()
		if err != nil {
			fmt.Printf("Error getting container cgroup path: %v\n", err)
			os.Exit(1)
		}

		// join cgroups
		if err := driver.Join(path, cgroups); err != nil {
			fmt.Printf("Error joining cgroup: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Successfully joined %s container cgroups\n", id)
	},
}

func init() {
	joinCmd.Flags().BoolVar(&blkio, "blkio", false, "blkio cgroup")
	joinCmd.Flags().BoolVar(&cpu, "cpu", false, "cpu cgroup")
	joinCmd.Flags().BoolVar(&cpuset, "cpuset", false, "cpuset cgroup")
	joinCmd.Flags().BoolVar(&devices, "devices", false, "devices cgroup")
	joinCmd.Flags().BoolVar(&freezer, "freezer", false, "freezer cgroup")
	joinCmd.Flags().BoolVar(&hugetlb, "hugetlb", false, "hugetlb cgroup")
	joinCmd.Flags().BoolVar(&memory, "memory", false, "memory cgroup")
	joinCmd.Flags().BoolVar(&net, "net", false, "net cgroup")
	joinCmd.Flags().BoolVar(&perfevent, "perfevent", false, "perfevent cgroup")
	joinCmd.Flags().BoolVar(&pids, "pids", false, "pids cgroup")
}
