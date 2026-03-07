package restlogger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func SendLog(url string, path string, log_data *LogRow) (error) {
    send_data, err := json.Marshal(log_data)
	if err != nil {
		return err
	}
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(send_data))
	if err != nil {
		return err
	}
    req.Header.Set("Content-Type", "application/json")
	req.URL.Path = path

    client := &http.Client{}
    resp, err := client.Do(req)
	if err != nil {
        return err
    }
	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)	
		return fmt.Errorf("Error write log: %s. Status: %v", string(body), resp.StatusCode)
	} else {
		return nil
	}
}