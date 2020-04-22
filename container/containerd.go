package container

import (
	"context"
	"fmt"

	containerdlib "github.com/containerd/containerd"
)

type containerd struct {
	container
}

func (c *containerd) CgroupPath() (string, error) {
	// create client
	client, err := containerdlib.New(c.socket, containerdlib.WithDefaultNamespace("k8s.io"))
	if err != nil {
		return "", fmt.Errorf("error creating containerd client with socket %s: %w", c.socket, err)
	}

	defer client.Close()

	// load container
	con, err := client.LoadContainer(context.Background(), c.id)
	if err != nil {
		return "", fmt.Errorf("error loading containerd container %s: %w", c.id, err)
	}

	// get container spec
	spec, err := con.Spec(context.Background())
	if err != nil {
		return "", fmt.Errorf("error getting containerd container spec: %w", err)
	}

	return spec.Linux.CgroupsPath, nil
}
