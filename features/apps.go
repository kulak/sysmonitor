package features

import (
	"gitlab.com/nest-machine/sysmonitor/core"
)

func AppsReport(conf *core.Config) []core.Group {
	var groups []core.Group
	for _, app := range conf.Apps {
		group := core.ExecReport(app.App, app.Args, app.Description)
		groups = append(groups, group)
	}
	return groups
}
