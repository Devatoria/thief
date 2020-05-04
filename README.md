# thief: a tool to steal your containers resources

Thief is a tool allowing to execute processes into a container cgroups instead of the actual ones, stealing its allowed resources.

This project has been made for ludic purposes only and is not expected to be used for another way than that. It manipulates your processes cgroups directly.

## Example

A quick example running the `sysbench` tool on CPU. First shot is from one container with no 1 CPU, second shot is from a container cgroup having only 0.5 CPU.

```
/# sysbench cpu run
CPU speed:
    events per second:  1135.99

General statistics:
    total time:                          10.0012s
    total number of events:              11363

Threads fairness:
    events (avg/stddev):           11363.0000/0.00
    execution time (avg/stddev):   9.8310/0.00

/# thief --sysfs-path /mnt/cgroup --runtime docker run --cpu 48b8916a8d82 sysbench cpu run
CPU speed:
    events per second:   516.93

General statistics:
    total time:                          10.0074s
    total number of events:              5174

Threads fairness:
    events (avg/stddev):           5174.0000/0.00
    execution time (avg/stddev):   9.8624/0.00
```

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

### Run a bash shell in a container CPU cgroup from the host

```
~# thief run --cpu 2c9eac1a0147e449208872330685c575933ddd2148888e8cdab899bbb7c14287 bash
~# 
```

### Attach a running process to a container CPU cgroup

```
~# thief attach --cpu 662c9eac1a0147e449208872330685c575933ddd2148888e8cdab899bbb7c14287 6
Successfully attached PID 666 to 62c9eac1a0147e449208872330685c575933ddd2148888e8cdab899bbb7c14287 container cgroups
```

## Notes

* You need to execute `thief` as root to be able to manipulate cgroups
* If executed in a container, this container needs to:
  * Mount the host `/sys/fs/cgroup` path
  * Be privileged

## TODOs

### Commands

- [x] Join command to attach the current process to a cgroup
- [x] Exit command to re-attach the current process to the main cgroup
- [x] Run command to run a command in a cgroup
- [x] Attach command to attach an existing process to a cgroup

### Other

- [ ] Tests
- [ ] CI
- [ ] Releases
