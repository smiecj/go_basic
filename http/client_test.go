package http

import (
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/smiecj/go_common/util/log"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/proxy"
)

// client with proxy
// extend: server with proxy: https://github.com/ananclub/ss5.git
func TestRequestWithSockProxy(t *testing.T) {
	dialSocksProxy, err := proxy.SOCKS5("tcp", "host:port", nil, proxy.Direct)
	if err != nil {
		fmt.Println("Error connecting to proxy:", err)
	}
	tr := &http.Transport{Dial: dialSocksProxy.Dial}
	log.Info("[test] transport: %v", tr)

	myClient := &http.Client{
		Transport: tr,
	}

	rsp, err := myClient.Get("http://www.inner_hostname")
	require.Nil(t, err)

	retBytes, readErr := io.ReadAll(rsp.Body)
	require.Empty(t, readErr)
	log.Info("[test] read ret: %s", string(retBytes))
}
