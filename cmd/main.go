package main

import (
	"fmt"
	"os"

	"github.com/Devatoria/thief/cgroup"
	"github.com/Devatoria/thief/container"
	"github.com/spf13/cobra"
)

var (
	// flags
	sysfsPath string
	runtime   string
	socket    string

	// shared
	driver   cgroup.Driver
	cRuntime container.Runtime
)

var rootCmd = &cobra.Command{
	Use:   "thief",
	Short: "thief is a CLI allowing to enter or execute a command into a container cgroups",
}

func init() {
	cobra.OnInitialize(initDriver, checkRuntime)

	rootCmd.PersistentFlags().StringVar(&sysfsPath, "sysfs-path", "/sys/fs/cgroup", "cgroup path if not the usual one (for instance when mounted)")
	rootCmd.PersistentFlags().StringVar(&runtime, "runtime", "containerd", "container runtime (docker or containerd)")
	rootCmd.PersistentFlags().StringVar(&socket, "socket", "", "docker or containerd daemon socket path (defaults /run/containerd/containerd.sock (containerd) or /var/run/docker.sock (docker))")
}

func initDriver() {
	driverCfg := cgroup.Config{
		Sysfs: sysfsPath,
	}

	driver = cgroup.NewDriver(driverCfg)
}

func checkRuntime() {
	cRuntime = container.Runtime(runtime)
	switch cRuntime {
	case container.ContainerdRuntime:
		if socket == "" {
			socket = "/run/containerd/containerd.sock"
		}
	case container.DockerRuntime:
		if socket == "" {
			socket = "/var/run/docker.sock"
		}
	default:
		fmt.Printf("Unexpected runtime %s, expecting either docker or containerd\n", runtime)
		os.Exit(1)
	}
}

func main() {
	rootCmd.AddCommand(joinCmd)
	rootCmd.AddCommand(exitCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
