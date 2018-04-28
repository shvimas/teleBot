package teleBot

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/proxy"
)

var timeout = time.Second * 30
var httpTransport = &http.Transport{}
var defaultClient = http.Client{Transport: httpTransport, Timeout: timeout}

func SendSOCKS5Request(req *http.Request, proxyAddr string, auth *proxy.Auth, client *http.Client) (*http.Response, error) {
	// create a socks5 dialer
	dialer, err := proxy.SOCKS5("tcp", proxyAddr, auth, proxy.Direct)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't connect to proxy:", err)
		return nil, err
	}
	// set our socks5 as the dialer
	// httpTransport.Dial = dialer.Dial -- deprecated
	httpTransport.DialContext = func(_ context.Context, network, addr string) (net.Conn, error) {
		return dialer.Dial(network, addr)
	}
	if client != nil {
		return client.Do(req)
	}
	return defaultClient.Do(req)
}
