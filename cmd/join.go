package main

import (
	"fmt"
	"os"

	"github.com/Devatoria/thief/cgroup"
	"github.com/Devatoria/thief/container"
	"github.com/spf13/cobra"
)

var joinCmd = &cobra.Command{
	Use:     "join <container ID>",
	Short:   "joins the given container cgroups",
	Example: "thief --runtime containerd --socket /run/containerd/containerd.sock join --cpu abcdef123456",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// get flags and args
		runtime, _ := cmd.Flags().GetString("runtime")
		socket, _ := cmd.Flags().GetString("socket")
		cpu, _ := cmd.Flags().GetBool("cpu")
		id := args[0]

		// check cgroup flags
		if !cpu {
			fmt.Println("No cgroup to join, please specify at least one cgroup to join")
			os.Exit(1)
		}

		// create container
		c, err := container.New(runtime, socket, id)
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

		cgroups := []cgroup.Kind{}
		if cpu {
			cgroups = append(cgroups, cgroup.CPU)
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
	joinCmd.Flags().Bool("cpu", false, "join cpu cgroup")
}
