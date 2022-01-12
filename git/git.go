package git

import (
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"tshockau/logger"
)

var tshockurl string = "https://github.com/Pryaxis/TShock/releases"
var regexurl string = "\\/Pryaxis\\/TShock\\/releases\\/download\\/(.+).zip"
var versionreg string = "v[0-9]?[0-9].[0-9]?[0-9].[0-9]?[0-9]"

func GetVersion() string {
	logger.Info("Requesting TShock version")
	reg, err := regexp.Compile(regexurl)
	if err != nil {
		logger.Error("Regex failed: " + err.Error())
		return ""
	}
	vreg, err := regexp.Compile(versionreg)
	if err != nil {
		logger.Error("Version Regex failed: " + err.Error())
		return ""
	}
	logger.Info("Polling GitHub...")
	resp, err := http.Get(tshockurl)
	if err != nil {
		logger.Error("Request failed: " + err.Error())
		return ""
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Parse failed: " + err.Error())
		return ""
	}

	var urllist []string
	logger.Info("Reading responce...")
	for _, v := range reg.FindAllString(string(b), -1) {
		urllist = append(urllist, "https://github.com"+v)
	}

	logger.Info("Current version " + vreg.FindString(urllist[0]))

	return urllist[0]
}

func Download(url string, out string) error {
	logger.Info("Downloading " + url + " as " + out)
	resp, err := http.Get(url)
	if err != nil {
		logger.Error("Request failed: " + err.Error())
		return err
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Request failed: " + err.Error())
		return err
	}

	logger.Info("Writing file to disk...")
	err = ioutil.WriteFile(out, bytes, 0644)
	if err != nil {
		logger.Error("Request failed: " + err.Error())
		return err
	}

	return nil
}
