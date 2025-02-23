package config

import (
	"os/user"
	"path/filepath"
	"runtime"

	"github.com/adrg/xdg"
)

func GetLocations() (res []string) {
	res = append(res, "./cli.json")

	xdgPath, err := xdg.ConfigFile(filepath.Join("WB2-CLI", "cli.json"))
	if err == nil {
		res = append(res, xdgPath)
	}

	if usr, err := user.Current(); err == nil {
		res = append(res, filepath.Join(usr.HomeDir, ".WB2-CLI", "cli.json"))
	}

	if runtime.GOOS != "windows" {
		res = append(res, "/etc/WB2-CLI/cli.json")
	}

	return
}
