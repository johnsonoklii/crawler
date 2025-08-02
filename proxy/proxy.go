package proxy

import (
	"fmt"
	"net/http"
	"net/url"
	"sync/atomic"
)

type ProxyFunc func(*http.Request) (*url.URL, error)

func RoundRobinProxySwitcher(proxyUrls ...string) (ProxyFunc, error) {
	if len(proxyUrls) < 1 {
		return nil, fmt.Errorf("no proxy url provided")
	}

	urls := make([]*url.URL, len(proxyUrls))
	for i, proxyUrl := range proxyUrls {
		u, err := url.Parse(proxyUrl)
		if err != nil {
			return nil, fmt.Errorf("parse proxy url failed: %v", err)
		}
		urls[i] = u
	}

	return (&roundRobinSwitcher{urls: urls, index: 0}).GetProxy, nil
}

type roundRobinSwitcher struct {
	urls  []*url.URL
	index uint32
}

func (r *roundRobinSwitcher) GetProxy(pr *http.Request) (*url.URL, error) {
	index := atomic.AddUint32(&(r.index), 1) - 1
	u := r.urls[index%uint32(len(r.urls))]
	return u, nil
}
