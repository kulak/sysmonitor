package features

import (
	"fmt"

	"gitlab.com/nest-machine/sysmonitor/core"
)

func BtrfsReport(conf *core.Config) []core.Group {
	var groups []core.Group
	for _, eachBtrfsDevice := range conf.BtrfsDevices {
		cmdArgs := []string{"device", "stats", eachBtrfsDevice}
		title := fmt.Sprintf("BTRFS Disk %s Health", eachBtrfsDevice)
		group := core.ExecReport("btrfs", cmdArgs, title)
		groups = append(groups, group)
	}
	return groups
}
