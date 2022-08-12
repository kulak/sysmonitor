package features

import (
	"fmt"

	"gitlab.com/nest-machine/sysmonitor/core"
)

func RsyncReport(conf *core.Config) []core.Group {
	var groups []core.Group
	for _, pair := range conf.RsyncDirs {
		var cmdArgs []string
		app := "rsync"
		if pair.UseSudo {
			cmdArgs = append(cmdArgs, "-n", "rsync")
			app = "sudo"
		}
		cmdArgs = append(cmdArgs, conf.RsyncArgs...)
		cmdArgs = append(cmdArgs, pair.Src)
		cmdArgs = append(cmdArgs, pair.Dst)
		groups = append(groups, core.ExecReport(app, cmdArgs, fmt.Sprintf("Backup %s to %s", pair.Src, pair.Dst)))
	}
	return groups
}
