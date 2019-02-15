package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type IP struct {
	IP string `json:"ip"`
}

type Update struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Content string `json:"content"`
	Proxied bool   `json:"proxied"`
	TTL     int    `json:"ttl"`
}

const (
	RecordName = "RECORD_NAME"
	RecordID   = "RECORD_ID"
	ZoneID     = "ZONE_ID"
	XAuthEmail = "X_AUTH_EMAIL"
	XAuthKey   = "X_AUTH_KEY"
)

func main() {
	res, err := http.Get("https://api.ipify.org?format=json")
	if err != nil {
		log.Fatal("get ip: ", err)
	}
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		log.Fatal("res status ", res.StatusCode, " ", res.Status)
	}
	defer res.Body.Close()

	var ip IP
	err = json.NewDecoder(res.Body).Decode(&ip)
	if err != nil {
		log.Fatal("decode: ", err)
	}

	up := Update{
		Type:    "A",
		Name:    os.Getenv(RecordName),
		Content: ip.IP,
		Proxied: true,
		TTL:     1,
	}
	var buf *bytes.Buffer
	err = json.NewEncoder(buf).Encode(up)
	if err != nil {
		log.Fatal("encode: ", err)
	}

	url := "https://api.cloudflare.com/client/v4/zones/" + os.Getenv(ZoneID) + "/dns_records/" + os.Getenv(RecordID)
	req, err := http.NewRequest(http.MethodPut, url, buf)
	if err != nil {
		log.Fatal("new request: ", err)
	}
	req.Header.Set("X-Auth-Email", os.Getenv(XAuthEmail))
	req.Header.Set("X-Auth-Key", os.Getenv(XAuthKey))
	req.Header.Set("Content-Type", "application/json")

	res, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("update: ", err)
	}
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		log.Fatal("update status: ", res.StatusCode, " ", res.Status)
	}
}
