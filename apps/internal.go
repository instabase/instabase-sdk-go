package apps

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
)

type apiResponse struct {
	APIVersion    string   `json:"api_version"`
	Err           string   `json:"error"`
	InputFileName string   `json:"input_file_name"`
	Records       []Record `json:"records"`
}

type apiRequest struct {
	InputFileContent []byte `json:"input_file_content"`
	InputFileName    string `json:"input_file_name"`
	Name             string `json:"name"`
	Version          string `json:"version"`
}

func (i *Client) buildHeaders() (http.Header, error) {
	authToken := i.Config.APIToken
	if authToken == "" {
		err := "API Token not supplied"
		i.log("Error while initializing headers: %s", err)
		return nil, fmt.Errorf(err)
	}

	return map[string][]string{
		"Authorization": {fmt.Sprintf("Bearer %s", authToken)},
		"Content-Type":  {"application/json"},
	}, nil
}

func (i *Client) buildRunURL(version string) (string, error) {

	parsedURL, err := url.Parse(i.Config.RootURL)
	if err != nil {
		err := fmt.Sprintf("Invalid API URL %s", err)
		i.log("Error while creating API endpoint URL: %s", err)
		return "", fmt.Errorf(err)
	}
	parsedURL.Path = path.Join(parsedURL.Path, "api",
		version, "/cloud/app/run")
	parsedPath := parsedURL.String()
	return parsedPath, nil
}

func (i *Client) verboseLog(format string, args ...interface{}) {
	if !i.Config.VerboseLog {
		return
	}
	i.Config.Logger.Printf(format, args...)
}

func (i *Client) log(format string, args ...interface{}) {
	i.Config.Logger.Printf(format, args...)
}
