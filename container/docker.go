package container

import (
	"context"
	"fmt"

	"github.com/docker/docker/client"
)

type docker struct {
	container
}

func (d *docker) CgroupPath() (string, error) {
	// initialize client with api version negociation
	c, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())
	if err != nil {
		return "", fmt.Errorf("error initializing docker client: %w", err)
	}

	// inspect the given container
	con, err := c.ContainerInspect(context.Background(), d.id)
	if err != nil {
		return "", fmt.Errorf("error inspecting containerd %s: %w", d.id, err)
	}

	return con.HostConfig.CgroupParent, nil
}
