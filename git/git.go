package git

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

var tshockurl string = "https://github.com/Pryaxis/TShock/releases"
var regexurl string = "\\/Pryaxis\\/TShock\\/releases\\/download\\/(.+).zip"

func GetVersion() string {
	reg, err := regexp.Compile(regexurl)
	if err != nil {
		fmt.Println("Failed to compile regex expression")
		fmt.Println(err)
		return ""
	}
	resp, err := http.Get(tshockurl)
	if err != nil {
		fmt.Println("Failed to request TShock version information")
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to parce body information")
		fmt.Println(err)
		return ""
	}

	var urllist []string
	for _, v := range reg.FindAllString(string(b), -1) {
		urllist = append(urllist, "https://github.com"+v)
	}

	return urllist[0]
}

func Download(url string, out string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(out, bytes, 0644)
	if err != nil {
		return err
	}

	return nil
}
