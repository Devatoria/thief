package main

import (
	"fmt"
	"os"

	"github.com/Devatoria/thief/cmd"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "thief",
	Short: "thief is a CLI allowing to enter or execute a command into a container cgroups",
}

func init() {
	rootCmd.PersistentFlags().String("sysfs-path", "/sys/fs/cgroup", "cgroup path if not the usual one (for instance when mounted)")
	rootCmd.PersistentFlags().String("runtime", "", "container runtime (docker or containerd)")
	rootCmd.PersistentFlags().String("socket", "", "docker or containerd daemon socket path")

	rootCmd.MarkPersistentFlagRequired("runtime")
	rootCmd.MarkPersistentFlagRequired("socket")
}

func main() {
	rootCmd.AddCommand(cmd.Join)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
