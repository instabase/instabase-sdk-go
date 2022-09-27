package apps

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/instabase/instabase-sdk-go/config"
)

// Client exposes all the API extraction capabilities of Instabase
type Client struct {
	Config *config.Config
}

// ExtractFromBytes extracts information for an input file represented by
// bytesContent and fileName
func (i *Client) ExtractFromBytes(bytesContent []byte,
	fileName, appName, appVersion string) (*Response, error) {
	if appName == "" {
		return nil, fmt.Errorf("Missing a required field: app name")
	}
	if appVersion == "" {
		return nil, fmt.Errorf("Missing a required field: app version")
	}

	requestURL, err := i.buildRunURL("v1")
	if err != nil {
		return nil, err
	}
	requestBody := apiRequest{
		InputFileContent: bytesContent,
		InputFileName:    fileName,
		Name:             appName,
		Version:          appVersion,
	}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		i.log("Error while building headers: %s", err)
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, requestURL,
		bytes.NewReader(jsonBody))
	if err != nil {
		i.log("Error while creating request: %s", err)
		return nil, err
	}

	req.Header, err = i.buildHeaders()
	if err != nil {
		i.log("Error while building headers: %s", err)
		return nil, err
	}
	client := http.Client{
		Timeout: 60 * time.Second,
	}
	var res *http.Response
	var body []byte
	var parsedResp apiResponse

	err = i.retryFunc(5, func() error {
		res, err = client.Do(req)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		if res.StatusCode < 200 || res.StatusCode >= 400 {
			return fmt.Errorf("Invalid status code %v: %v", res.StatusCode, res.Body)
		}
		body, err = ioutil.ReadAll(res.Body)
		if err := json.Unmarshal(body, &parsedResp); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		i.log("Error while fetching response from API: %s", err)
		return nil, err
	}

	result := Response{
		Error:         parsedResp.Err,
		Records:       parsedResp.Records,
		APIVersion:    parsedResp.APIVersion,
		InputFileName: parsedResp.InputFileName,
	}
	return &result, nil
}

// ExtractFromInputURL extracts information for provided file located at fileURL
func (i *Client) ExtractFromInputURL(fileURL, appName,
	appVersion string) (*Response, error) {
	if fileURL == "" {
		return nil, fmt.Errorf("%s is an invalid file URL", fileURL)
	}
	parsedURL, err := url.Parse(fileURL)
	if err != nil {
		err := fmt.Sprintf("Invalid file URL %s", err)
		i.log("Error while parsing file URL: %s", err)
		return nil, fmt.Errorf(err)
	}

	resp, err := http.Get(parsedURL.String())
	if err != nil {
		i.log("Error while fetching file using provided URL: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		i.log("Error while reading API response %v", err)
	}
	_, filePath := path.Split(fileURL)
	return i.ExtractFromBytes(body, filePath, appName, appVersion)
}
