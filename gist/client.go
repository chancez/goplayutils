package gist

import (
	"net/http"
	"time"

	"github.com/gregjones/httpcache"
	"github.com/sourcegraph/apiproxy"
	"github.com/sourcegraph/apiproxy/service/github"
)

func NewCachingHttpClient(trans *httpcache.Transport,
	validator *apiproxy.Validator) *http.Client {

	if trans == nil {
		trans = httpcache.NewMemoryCacheTransport()
	}
	var check apiproxy.Validator
	if validator == nil {
		age := &githubproxy.MaxAge{
			User:         time.Hour * 24,
			Repository:   time.Hour * 24,
			Repositories: time.Hour * 24,
			Activity:     time.Hour * 12,
		}
		check = age.Validator()
	} else {
		check = *validator
	}
	transport := &apiproxy.RevalidationTransport{
		Transport: trans,
		Check:     check,
	}
	return &http.Client{Transport: transport}
}
