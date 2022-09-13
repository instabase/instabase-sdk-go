# Instabase Apps Golang SDK

Use this SDK to access apps on apps.instabase.com.

## Configuration

You can create config object from environment variables or supply your own config.

### Environment variable based config

You will need to set the following environment variables:

1. `IB_API_TOKEN` : Your API token from Instabase Apps website.
2. `IB_ROOT_URL` : Root URL of Instabase App website.
3. `IB_VERBOSE_LOG` (Not mandatory): Valid values are true and false. Set to true to turn on verbose logging.

### Manual config

You can create your own config object with custom logger using the config object in SDK.

Example:

```
    cfg := config.Config{
		Logger:   &customLogger,
		APIToken: customToken,
		RootURL:  customIBURL,
	}
```

## Getting Started

Example below demonstrates usage of the SDK to extract contents of an image using Instabase's Bank Statements application (version 3.0.0):

```go
package main
import (
	"fmt"

	"github.com/instabase/instabase-sdk-go/api"
	"github.com/instabase/instabase-sdk-go/config"
)
func main() {
	config, _ := config.NewAppConfigFromEnv()
	extractor := api.IBAPIExtractor{Config: config}
	fileURL, appName, appVersion := "https://apps.instabase.com/static/assets/images/cloud-developers/us-bs/sample-us-bs-1.jpeg", "US Bank Statements", "3.0.0"
	resp, _ := extractor.ExtractFromInputURL(fileURL, appName, appVersion)
	fmt.Println(resp)
}
```
