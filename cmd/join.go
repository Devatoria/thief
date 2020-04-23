package cmd

import (
	"fmt"
	"os"

	"github.com/Devatoria/thief/cgroup"
	"github.com/Devatoria/thief/container"
	"github.com/spf13/cobra"
)

// Join command
var Join = &cobra.Command{
	Use:     "join <container ID>",
	Short:   "joins the given container cgroups",
	Example: "thief --runtime containerd --socket /run/containerd/containerd.sock join --cpu abcdef123456",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// get flags and args
		sysfsPath, _ := cmd.Flags().GetString("sysfs-path")
		runtime, _ := cmd.Flags().GetString("runtime")
		socket, _ := cmd.Flags().GetString("socket")
		cpu, _ := cmd.Flags().GetBool("cpu")
		id := args[0]

		// check cgroup flags
		if !cpu {
			fmt.Println("no cgroup to join, please specify at least one cgroup to join")
			os.Exit(1)
		}

		// create container
		c, err := container.New(runtime, socket, id)
		if err != nil {
			fmt.Printf("error creating container socket: %v\n", err)
			os.Exit(1)
		}

		// retrieve cgroup path
		path, err := c.CgroupPath()
		if err != nil {
			fmt.Printf("error getting container cgroup path: %v\n", err)
			os.Exit(1)
		}

		// join cgroup
		driver := cgroup.NewDriver(sysfsPath, path, cpu)
		if err := driver.Join(); err != nil {
			fmt.Printf("error joining cgroup: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("successfully joined %s container cgroups at path %s\n", id, path)
	},
}

func init() {
	Join.Flags().Bool("cpu", false, "join cpu cgroup")
}
