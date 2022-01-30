package main

import "fmt"

func journalReport() Group {
	cmdArgs := []string{"-p", "0..3", "--since", "24 hour ago", "--no-pager"}
	return execReport("journalctl", cmdArgs, "Journal Errors in Last 24 Hours")
}

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
