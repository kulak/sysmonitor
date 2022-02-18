package main

import "fmt"

func btrfsReport(conf *Config) []Group {
	var groups []Group
	for _, eachBtrfsDevice := range conf.BtrfsDevices {
		cmdArgs := []string{"device", "stats", eachBtrfsDevice}
		title := fmt.Sprintf("BTRFS Disk %s Health", eachBtrfsDevice)
		group := execReport("btrfs", cmdArgs, title)
		groups = append(groups, group)
	}
	return groups
}
