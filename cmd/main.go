package main

import (
	"fmt"
	"os"

	"github.com/Devatoria/thief/cgroup"
	"github.com/Devatoria/thief/container"
	"github.com/spf13/cobra"
)

var (
	// global flags
	sysfsPath string
	runtime   string
	socket    string

	// shared flags
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

	// shared
	driver   cgroup.Driver
	cRuntime container.Runtime
	cgroups  []cgroup.Kind
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

// initCgorups builds enabled cgroups list and
// ensures at least one has been enabled
func initCgroups(cmd *cobra.Command, args []string) {
	cgroups = cgroup.AppendCgroups(blkio, cpu, cpuset, devices, freezer, hugetlb, memory, net, perfevent, pids)
	if len(cgroups) == 0 {
		fmt.Println("No cgroup given, please specify at least one cgroup")
		os.Exit(1)
	}
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
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(exitCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
