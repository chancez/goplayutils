package playground

import (
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	shareUrl     = "http://play.golang.org/share"
	contentType  = "application/x-www-form-urlencoded; charset=UTF-8"
	PlayUrl      = "http://play.golang.org/p/"
	package_main = "package main"
)

func HasMain(content *string) bool {
	return strings.Contains(*content, package_main)
}

func SendToPlayground(content *string) (string, error) {
	buf := strings.NewReader(*content)
	resp, err := http.Post(shareUrl, contentType, buf)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func GetPlayUrl(content *string) (string, error) {
	shareId, err := SendToPlayground(content)
	if err != nil {
		return "", err
	}
	return PlayUrl + shareId, nil
}
