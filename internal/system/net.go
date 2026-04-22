package system

import (
	"io"
	"net/http"
	"strings"
)

func GetPublicIP() (string, error) {
	resp, err := http.Get("https://api.ipify.org")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	
	ip, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	
	return strings.TrimSpace(string(ip)), nil
}