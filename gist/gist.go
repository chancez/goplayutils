package gist

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"code.google.com/p/goauth2/oauth"

	"github.com/ecnahc515/goplayutils/playground"
	"github.com/google/go-github/github"
	"github.com/gregjones/httpcache"
	"github.com/sourcegraph/apiproxy"
	"github.com/sourcegraph/apiproxy/service/github"
)

var ErrNoPackageMain = errors.New("Could not find a \"package main\" in any gist files.")

func FindMain(gist *github.Gist) (string, error) {
	for _, file := range gist.Files {
		if playground.HasMain(file.Content) {
			return *file.Content, nil
		}
	}
	return "", ErrNoPackageMain
}

func NewCachingHttpClient(token string, cache httpcache.Cache,
	validator apiproxy.Validator) *http.Client {

	if cache == nil {
		cache = httpcache.NewMemoryCache()
	}

	trans := httpcache.NewTransport(cache)

	if token != "" {
		t := &oauth.Transport{
			Token: &oauth.Token{AccessToken: token},
		}
		trans.Transport = t
	}

	if validator == nil {
		age := &githubproxy.MaxAge{
			User:         time.Hour * 24,
			Repository:   time.Hour * 24,
			Repositories: time.Hour * 24,
			Activity:     time.Hour * 12,
		}
		validator = age.Validator()
	}

	transport := &apiproxy.RevalidationTransport{
		Transport: trans,
		Check:     validator,
	}
	return &http.Client{Transport: transport}
}

func GetGist(client *github.Client, gistId string) (string, error) {
	// passing in a url for a gist
	gst, _, err := client.Gists.Get(gistId)
	if err != nil {
		return "", fmt.Errorf("Error retrieving gist: %s", err)
	}
	var content string
	content, err = FindMain(gst)
	if err != nil {
		return "", err
	}
	return content, nil
}
