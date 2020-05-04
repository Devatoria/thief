package main

import (
	"fmt"
	"os"

	"github.com/Devatoria/thief/container"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:     "run <container ID>",
	Short:   "runs the given command in the given container cgroups",
	Example: "thief --runtime containerd --socket /run/containerd/containerd.sock run --cpu abcdef123456 sysbench cpu run",
	Args:    cobra.MinimumNArgs(2),
	PreRun:  initCgroups,
	Run: func(cmd *cobra.Command, args []string) {
		// get args
		id := args[0]
		commandName := args[1]
		commandArgs := args[2:]

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

		// run command in enabled cgroups
		if err := driver.Run(path, cgroups, commandName, commandArgs); err != nil {
			fmt.Printf("Error running command in the cgroup: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Successfully executed %s command in %s container cgroups\n", commandName, id)
	},
}

func init() {
	runCmd.Flags().BoolVar(&blkio, "blkio", false, "blkio cgroup")
	runCmd.Flags().BoolVar(&cpu, "cpu", false, "cpu cgroup")
	runCmd.Flags().BoolVar(&cpuset, "cpuset", false, "cpuset cgroup")
	runCmd.Flags().BoolVar(&devices, "devices", false, "devices cgroup")
	runCmd.Flags().BoolVar(&freezer, "freezer", false, "freezer cgroup")
	runCmd.Flags().BoolVar(&hugetlb, "hugetlb", false, "hugetlb cgroup")
	runCmd.Flags().BoolVar(&memory, "memory", false, "memory cgroup")
	runCmd.Flags().BoolVar(&net, "net", false, "net cgroup")
	runCmd.Flags().BoolVar(&perfevent, "perfevent", false, "perfevent cgroup")
	runCmd.Flags().BoolVar(&pids, "pids", false, "pids cgroup")
}
