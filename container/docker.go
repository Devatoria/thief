package container

type docker struct {
	container
}

func (d *docker) CgroupPath() (string, error) {
	return "", nil
}
