package apps

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/instabase/instabase-sdk-go/config"
)

func TestExtractFromBytes(t *testing.T) {
	appConfig, _ := config.NewAppConfigFromEnv()
	client := Client{Config: appConfig}
	fileName := "test_data_bank_statement.pdf"
	currentPath, err := os.Getwd()
	if err != nil {
		t.Errorf("Error while getting path: %s", err)
	}
	bytesread, err := os.ReadFile(path.Join(currentPath, "../", "testdata", fileName))
	if err != nil {
		t.Errorf("Error while reading file: %s", err)
	}
	appName, appVersion := "US Bank Statements", "3.0.0"
	resp, err := client.ExtractFromBytes(bytesread, fileName, appName, appVersion)
	if err != nil {
		t.Errorf("Error while extracting information: %s", err)
	}
	fmt.Println(resp)
	if resp.Error != "" {
		t.Errorf("Error while extracting information: %s", resp.Error)
	}

}
