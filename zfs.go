package main

func zfsReport() Group {
	cmdArgs := []string{"list"}
	group := execReport("zpool", cmdArgs, "ZFS Pool List")
	return group
}
