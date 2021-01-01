package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/elazarl/goproxy"
	proxy2 "github.com/vay3t/proxy-rotate/pkg/proxy"
)

func main() {
	host := flag.String("host", "127.0.0.1", "host address")
	port := flag.Int("port", 9999, "host port")

	flag.Parse()

	portString := strconv.Itoa(*port)

	httpServer := &http.Server{
		Addr:         *host + ":" + portString,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	httpServer.SetKeepAlivesEnabled(false)
	proxy2.ProxyList = new(proxy2.ProxyBucket)
	go proxy2.ProxyList.Start()
	fmt.Println("Starting proxy server ", "http://"+*host+":"+portString)
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	proxy.Tr = proxy2.NewTransport(proxy)

	httpServer.Handler = proxy
	log.Fatal(httpServer.ListenAndServe())
}
