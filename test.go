package main

import (
	"fmt"

	"github.com/instabase/instabase-sdk-go/api"
	"github.com/instabase/instabase-sdk-go/config"
)

func main() {
	appConfig, _ := config.NewAppConfigFromEnv()
	extractor := api.IBAPIExtractor{Config: appConfig}
	fileURL, appName, appVersion := "https://apps.instabase.com/static/assets/images/cloud-developers/us-bs/sample-us-bs-1.jpeg", "US Bank Statements", "3.0.0"
	resp, _ := extractor.ExtractFromInputURL(fileURL, appName, appVersion)
	fmt.Println(resp)
}
