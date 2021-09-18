package main

import (
	"fmt"
	"net/http/httputil"
	"github.com/gorilla/mux"
	"github.com/prashanthca/basic-auth-go/strategies"
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
		r.PathPrefix("/").HandlerFunc(strategies.FwdOptionsReq(rp)).Methods("OPTIONS")
		r.PathPrefix("/").HandlerFunc(strategies.DefaultUnauthorized(rp))
		return r, errors.New("Invalid strategy")
	}
	fmt.Println("Routes Registered")
	return r,nil
}