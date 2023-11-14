package workers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GETRequest(url string, xrelayer string, authkey string) (*http.Response, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("X-Relayer", xrelayer)
	req.Header.Add("Authorization", "Bearer "+authkey)

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		log.Println("Unable to authenticate to server! StatusCode: ", resp.StatusCode)
		return nil, err
	}

	return resp, nil
}

func ReadResponseBody(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func POSTRequestWithoutBody(url string, xrelayer string, authkey string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Relayer", xrelayer)
	req.Header.Add("Authorization", "Bearer "+authkey)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 204 {
		return nil, fmt.Errorf("server returned no-content")
	}

	if resp.StatusCode == 401 {
		return nil, fmt.Errorf("unable to authenticate to the server. an error occured")
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("internal server error")
	}

	return resp, nil
}

func EncodeGetJobAssignmentReqBody(chainid ChainIds) ([]byte, error) {
	body := struct {
		Network string `json:"network"`
	}{
		Network: string(chainid),
	}

	out, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}
	return out, nil
}

func POSTRequestWithBody(url string, xrelayer string, authkey string, body []byte) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Relayer", xrelayer)
	req.Header.Add("Authorization", "Bearer "+authkey)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 204 {
		return nil, fmt.Errorf("server returned no-content")
	}

	if resp.StatusCode == 401 {
		return nil, fmt.Errorf("unable to authenticate to the server. an error occured")
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("internal server error")
	}

	return resp, nil
}
