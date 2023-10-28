//go:build darwin && !debug
// +build darwin,!debug

package platform

import (
	"path"
)

func initVariables() {
	settingsDirectory := "/Library/Application Support/RadicalVPN"

	serviceFile = path.Join(settingsDirectory, "service.txt")
	settingsFile = path.Join(settingsDirectory, "settings.json")

	wireguardPath = path.Join("")
	wireguardQuickPath = path.Join("")
	wireguardConfigPath = path.Join(settingsDirectory, "radicalvpn.conf")
}
