package main

func journalReport() Group {
	cmdArgs := []string{"-p", "0..3", "--since", "24 hour ago", "--no-pager"}
	return execReport("journalctl", cmdArgs, "Journal Errors in Last 24 Hours")
}
