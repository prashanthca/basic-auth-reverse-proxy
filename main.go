package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"github.com/jessevdk/go-flags"
)

var opts struct {
	Address string `short:"a" long:"addr" description:"Proxy server address" required:"true"`
	Deta bool `long:"deta" description:"Enable Deta (deta.sh) auth strategy"`
	DetaProjectKey string `long:"deta-project-key" description:"Deta project key"`
	DetaBaseId string `long:"deta-base-id" description:"Deta base"`
	Scheme string `short:"p" long:"scheme" description:"Proxy target scheme" required:"true"`
	Host string `short:"h" long:"host" description:"Proxy target host" required:"true"`
}

func main(){
	_, err := flags.Parse(&opts)
	if err != nil {
		return
	}
	rp := httputil.ReverseProxy {
		Director: NewDirector(),
	}
	fmt.Printf("Address: %v\nTarget: %v\n", opts.Address, opts.Scheme+"://"+opts.Host)
	r, err := NewRouter(&rp)
	if err != nil {	
		fmt.Println("ERROR:", err)
		return
	}
	srv := http.Server {
		Handler: r,
		Addr: opts.Address,
	}
	srv.ListenAndServe()
}