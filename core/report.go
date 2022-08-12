package core

import (
	"fmt"
	"os/exec"
	"strings"
)

func ExecReport(app string, cmdArgs []string, title string) Group {
	var msgs []Message
	cmd := exec.Command(app, cmdArgs...)
	var out []byte
	var err error
	if out, err = cmd.CombinedOutput(); err != nil {
		msgs = append(msgs, Msg(err.Error(), errorLvl, p))
		if len(out) > 0 {
			msgs = append(msgs, Msg(string(out), errorLvl, code))
		}
	} else {
		msgs = append(msgs, Msg(string(out), infoLvl, code))
	}
	return Group{
		Title:       title,
		Description: fmt.Sprintf("%s %s", app, strings.Join(quote(cmdArgs), " ")),
		Msgs:        msgs,
	}
}

func quote(strs []string) []string {
	var rv []string
	for _, each := range strs {
		rv = append(rv, fmt.Sprintf("'%s'", each))
	}
	return rv
}
