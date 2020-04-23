package main

import (
	"fmt"
	"os"

	"github.com/Devatoria/thief/cgroup"
	"github.com/spf13/cobra"
)

var (
	sysfsPath string
	driver    cgroup.Driver
)

var rootCmd = &cobra.Command{
	Use:   "thief",
	Short: "thief is a CLI allowing to enter or execute a command into a container cgroups",
}

func init() {
	cobra.OnInitialize(initDriver)

	rootCmd.PersistentFlags().StringVar(&sysfsPath, "sysfs-path", "/sys/fs/cgroup", "cgroup path if not the usual one (for instance when mounted)")
	rootCmd.PersistentFlags().String("runtime", "containerd", "container runtime (docker or containerd)")
	rootCmd.PersistentFlags().String("socket", "/run/containerd/containerd.sock", "docker or containerd daemon socket path")
}

func initDriver() {
	driverCfg := cgroup.Config{
		Sysfs: sysfsPath,
	}

	driver = cgroup.NewDriver(driverCfg)
}

func main() {
	rootCmd.AddCommand(joinCmd)
	rootCmd.AddCommand(exitCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
