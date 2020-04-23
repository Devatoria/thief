package cgroup

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// getCgroupPath parses the given pid cgroup file and
// returns the path of the given cgroup
func getCgroupPath(pid int, kind Kind) (string, error) {
	entries, err := parseCgroupFile(pid)
	if err != nil {
		return "", fmt.Errorf("error getting cgroup path for PID %d and kind %s: %w", pid, kind, err)
	}

	entry := ""
	ok := false

	if entry, ok = entries[string(kind)]; !ok {
		return "", fmt.Errorf("error getting cgroup path for PID %d: unknown kind %s", pid, kind)
	}

	return entry, nil
}

// parseCgroupFile reads the cgroup file for the given pid
// and builds a map kind => path like:
// cpu => /kubepods/besteffort/pod7180b...
func parseCgroupFile(pid int) (map[string]string, error) {
	entries := map[string]string{}

	// open cgroup file for the given pid
	file, err := os.Open(fmt.Sprintf("/proc/%d/cgroup", pid))
	if err != nil {
		return nil, fmt.Errorf("error opening PID %d cgroup file: %w", pid, err)
	}

	defer file.Close()

	// read each line to fill the map
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		// format is id:kind:path
		params := strings.Split(sc.Text(), ":")
		if len(params) != 3 {
			return nil, fmt.Errorf("unexpected cgroup file format for PID %d: %s", pid, strings.Join(params, ":"))
		}

		// fill map entry
		kind := params[1]
		path := params[2]
		entries[kind] = path
	}

	// check errors
	if err := sc.Err(); err != nil {
		return nil, fmt.Errorf("error reading PID %d cgroup file: %w", pid, err)
	}

	return entries, nil
}
