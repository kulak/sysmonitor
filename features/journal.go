package features

import "gitlab.com/nest-machine/sysmonitor/core"

func JournalReport() core.Group {
	cmdArgs := []string{"-p", "0..3", "--since", "24 hour ago", "--no-pager"}
	return core.ExecReport("journalctl", cmdArgs, "Journal Errors in Last 24 Hours")
}
