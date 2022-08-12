package features

import "gitlab.com/nest-machine/sysmonitor/core"

func ZfsReport() core.Group {
	cmdArgs := []string{"list"}
	group := core.ExecReport("zpool", cmdArgs, "ZFS Pool List")
	return group
}
