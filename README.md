# thief

Thief is a tool allowing to execute processes into a container cgroups instead of the actual ones, stealing its allowed resources.

## Usage

### Global flags

* `--sysfs-path` is the mounting point of the cgroups files, defaulting to `/sys/fs/cgroup`
* `--runtime` is the runtime to use, either `docker` or `containerd`
* `--socket` is the runtime socket path
  * `containerd` runtime default is `/run/containerd/containerd.sock`
  * `docker` runtime default is `/var/run/docker.sock`

### Join a container CPU cgroup from the host

```
~# thief join --cpu 2c9eac1a0147e449208872330685c575933ddd2148888e8cdab899bbb7c14287

Successfully joined 2c9eac1a0147e449208872330685c575933ddd2148888e8cdab899bbb7c14287 container cgroups
```

### Join a container CPU cgroup from another container

The host `/sys/fs/cgroup` directory has been mounted in the container running thief in the `/mnt/cgroup` mount point.

```
~# thief --sysfs-path /mnt/cgroup join --cpu 2c9eac1a0147e449208872330685c575933ddd2148888e8cdab899bbb7c14287

Successfully joined 2c9eac1a0147e449208872330685c575933ddd2148888e8cdab899bbb7c14287 container cgroups
```

### Join back root cgroups

The `exit` subcommand re-attaches the current process to the same cgroups as the PID `1` process.

```
~# thief exit

Exited successfully!
```

## Notes

* You need to execute `thief` as root to be able to manipulate cgroups
* If executed in a container, this container needs to:
  * Mount the host `/sys/fs/cgroup` path
  * Be privileged
