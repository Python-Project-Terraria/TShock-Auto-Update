package tshock

import (
	"tshockau/git"
	"tshockau/logger"
	"tshockau/ziper"
)

func Update() {
	s := git.GetVersion()
	git.Download(s, "test.zip")
	ziper.Unpack("test.zip", ".")
	logger.Info("TShock update complete")
}
