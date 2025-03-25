package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var torrents []Torrents
var torrent Torrent

func req_torrents() error {
	reqBody := []byte(`{"action":"list"}`)

	resp, err := request(reqBody)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&torrents); err != nil {
		return err
	}

	return nil
}

func req_add(magnetUrl, title string) error {
	reqBody := []byte(fmt.Sprintf(`{"action": "add","link": "%s","title": "%s","category": "","poster": "","save_to_db": true}`, magnetUrl, title))

	resp, err := request(reqBody)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var status Status
	if err := json.NewDecoder(resp.Body).Decode(&status); err != nil {
		return err
	}
	fmt.Println(status.Stat)

	return nil
}

// magnet:?xt=urn:btih:35a100b8e0d90d98cdcaa0e369aee47619df2e02&dn=Total.Drama.S01.1080p.NF.WEBRip.DDP2.0.x264-NTb%5Brartv%5D
func req_delete(torrentHash string) error {
	reqBody := []byte(`{"action":"rem","hash":"` + torrentHash + `"}`)

	resp, err := request(reqBody)
	if err != nil {
		return err
	}
	resp.Body.Close()

	return nil
}

func request(reqBody []byte) (*http.Response, error) {
	req, err := http.NewRequest("POST", server_url+"/torrents", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	if basic_auth_username != "" && basic_auth_password != "" {
		req.SetBasicAuth(basic_auth_username, basic_auth_password)
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Failed to send request : %d", resp.StatusCode)
	}

	return resp, nil
}
