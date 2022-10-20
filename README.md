# Instabase SDK

Use this SDK to access Instabase applications.

## Configuration

You can create config object from environment variables or supply config object at runtime.

### Environment variable based config

You will need to set the following environment variables:

1. `IB_API_TOKEN` : Your API token from Instabase website.
2. `IB_ROOT_URL` : Root URL of Instabase website.
3. `IB_VERBOSE_LOG` (optional): Valid values are true and false. Set to true to turn on verbose logging.

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

Example below demonstrates usage of the SDK to extract contents of an input file using Instabase's Bank Statements application (version 3.0.0):

### Input file represented by remote URL

```go
package main
import (
	"fmt"

	"github.com/instabase/instabase-sdk-go/apps"
	"github.com/instabase/instabase-sdk-go/config"
)
func main() {
	config, _ := config.NewAppConfigFromEnv()
	client := apps.Client{Config: config}
	fileURL, appName, appVersion := "<FILE_URL_HERE>", "US Bank Statements", "3.0.0"
	resp, _ := client.ExtractFromInputURL(fileURL, appName, appVersion)
	fmt.Println(resp)
}
```

### Input file represented by file path

```go
package main

import (
	"fmt"
	"os"

	"github.com/instabase/instabase-sdk-go/apps"
	"github.com/instabase/instabase-sdk-go/config"
)

func main() {
	appConfig, _ := config.NewAppConfigFromEnv()
	client := apps.Client{Config: appConfig}
	fileName := "<FILE_PATH_HERE>"
	bytesread, _ := os.ReadFile(fileName)
	appName, appVersion := "US Bank Statements", "3.0.0"
	resp, _ := client.ExtractFromBytes(bytesread, fileName, appName, appVersion)
	fmt.Println(resp)
}

```
