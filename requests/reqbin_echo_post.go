package requests

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// ReqbinEchoPost sends super simple POST exho request to "https://reqbin.com/echo/post/json"
func ReqbinEchoPost() error {
	reqBody, err := json.Marshal(map[string]string{})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "https://reqbin.com/echo/post/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	_, err = client.Do(req)

	return err
}
