package gist

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"code.google.com/p/goauth2/oauth"

	"github.com/google/go-github/github"
	"github.com/gregjones/httpcache"
	"github.com/sourcegraph/apiproxy"
	"github.com/sourcegraph/apiproxy/service/github"
)

const package_main = "package main"

var ErrNoPackageMain = errors.New("Could not find a \"package main\" in any gist files.")

func FindMain(gist *github.Gist) (string, error) {
	for _, file := range gist.Files {
		if strings.Contains(*file.Content, package_main) {
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
