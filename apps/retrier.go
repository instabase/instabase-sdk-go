package apps

import (
	"fmt"
	"time"
)

// retryFunc retries the function supplied using exponential backoff mechanism
// The maximum retries are supplied by the user. If the function execution is not
// successful after maximum retries, an error is returned
func (i *Client) retryFunc(totalCount int, fun func() error) error {
	waitTime := 1 * time.Second
	for ; totalCount > 0; totalCount-- {
		err := fun()
		if err != nil {
			i.verboseLog("Error while executing function. %v retries left",
				totalCount)
			time.Sleep(waitTime)
			waitTime *= 2
			continue
		}
		return nil
	}
	return fmt.Errorf("Max retries reached")
}
