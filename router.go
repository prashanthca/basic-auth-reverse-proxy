package main

import (
	"fmt"
	"net/http/httputil"
	"github.com/gorilla/mux"
	"github.com/prashanthca/basic-auth-reverse-proxy/strategies"
	"errors"
)

func NewRouter(rp *httputil.ReverseProxy) (*mux.Router, error){
	r := mux.NewRouter()
	if opts.Deta {
		if opts.DetaProjectKey == "" || opts.DetaBaseId == "" {
			return nil, errors.New("Invalid Deta project key or Base ID")
		}
		deta := strategies.Deta{opts.DetaProjectKey, opts.DetaBaseId}
		r.PathPrefix("/").HandlerFunc(strategies.FwdOptionsReq(rp)).Methods("OPTIONS")
		r.PathPrefix("/").HandlerFunc(deta.RequestHandler(rp))
	} else {
		return nil, errors.New("Invalid strategy")
	}
	return r,nil
}