package git

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

var tshockurl string = "https://github.com/Pryaxis/TShock/releases"
var regexurl string = "\\/Pryaxis\\/TShock\\/releases\\/download\\/v[0-9].[0-9].[0-9][0-9]\\/TShock[0-9].[0-9].[0-9][0-9]_Terraria[0-9].[0-9].[0-9].[0-9].zip"

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
		urllist = append(urllist, "https://github.com/Pryaxis/TShock/releases/download"+v+".zip")
	}

	return urllist[0]
}
