package container

import "fmt"

// Container is an interface to abstract container
// related needed methods
type Container interface {
	CgroupPath() (string, error)
}

type container struct {
	id     string
	socket string
}

func New(runtime Runtime, socket, id string) (Container, error) {
	c := container{
		id:     id,
		socket: socket,
	}

	switch runtime {
	case ContainerdRuntime:
		return &containerd{c}, nil
	case DockerRuntime:
		return &docker{c}, nil
	default:
		return nil, fmt.Errorf("unexpected runtime: %s", runtime)
	}
}
