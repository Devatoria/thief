package main

import (
	"fmt"
	"os"

	"github.com/Devatoria/thief/cmd/utils"
	"github.com/Devatoria/thief/container"
	"github.com/spf13/cobra"
)

var (
	blkio     bool
	cpu       bool
	cpuset    bool
	devices   bool
	freezer   bool
	hugetlb   bool
	memory    bool
	net       bool
	perfevent bool
	pids      bool
)

var joinCmd = &cobra.Command{
	Use:     "join <container ID>",
	Short:   "joins the given container cgroups",
	Example: "thief --runtime containerd --socket /run/containerd/containerd.sock join --cpu abcdef123456",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// get flags and args
		id := args[0]

		// check cgroup flags
		if !cpu {
			fmt.Println("No cgroup to join, please specify at least one cgroup to join")
			os.Exit(1)
		}

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

		// build enabled cgroups list
		cgroups := utils.AppendCgroups(blkio, cpu, cpuset, devices, freezer, hugetlb, memory, net, perfevent, pids)

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
