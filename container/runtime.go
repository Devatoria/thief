package container

// Runtime represents a container runtime like docker or containerd
type Runtime string

const (
	// DockerRuntime is the docker runtime
	DockerRuntime Runtime = "docker"

	// ContainerdRuntime is the containerd runtime
	ContainerdRuntime Runtime = "containerd"
)
