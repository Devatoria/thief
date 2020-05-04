package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Devatoria/thief/container"
	"github.com/spf13/cobra"
)

var attachCmd = &cobra.Command{
	Use:     "attach <container ID> <PID>",
	Short:   "attach the given PID to the given container cgroups",
	Example: "thief --runtime containerd --socket /run/containerd/containerd.sock attach --cpu abcdef123456 666",
	Args: func(cmd *cobra.Command, args []string) error {
		// check args count
		if err := cobra.ExactArgs(2)(cmd, args); err != nil {
			return err
		}

		// check PID format
		if _, err := strconv.Atoi(args[1]); err != nil {
			return fmt.Errorf("Invalid PID %s, PID must be an integer", args[1])
		}

		return nil
	},
	PreRun: initCgroups,
	Run: func(cmd *cobra.Command, args []string) {
		// get args
		id := args[0]
		pid, _ := strconv.Atoi(args[1])

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
		if err := driver.Attach(path, cgroups, pid); err != nil {
			fmt.Printf("Error attaching PID %d in the cgroup: %v\n", pid, err)
			os.Exit(1)
		}

		fmt.Printf("Successfully attached PID %d to %s container cgroups\n", pid, id)
	},
}

func init() {
	attachCmd.Flags().BoolVar(&blkio, "blkio", false, "blkio cgroup")
	attachCmd.Flags().BoolVar(&cpu, "cpu", false, "cpu cgroup")
	attachCmd.Flags().BoolVar(&cpuset, "cpuset", false, "cpuset cgroup")
	attachCmd.Flags().BoolVar(&devices, "devices", false, "devices cgroup")
	attachCmd.Flags().BoolVar(&freezer, "freezer", false, "freezer cgroup")
	attachCmd.Flags().BoolVar(&hugetlb, "hugetlb", false, "hugetlb cgroup")
	attachCmd.Flags().BoolVar(&memory, "memory", false, "memory cgroup")
	attachCmd.Flags().BoolVar(&net, "net", false, "net cgroup")
	attachCmd.Flags().BoolVar(&perfevent, "perfevent", false, "perfevent cgroup")
	attachCmd.Flags().BoolVar(&pids, "pids", false, "pids cgroup")
}
