package strategies

import (
	"net/http"
	"net/http/httputil"
)

type Strategy interface {
	RequestHandler()
}

func FwdOptionsReq(rp *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func (w http.ResponseWriter, r *http.Request){
		rp.ServeHTTP(w, r)
	}
}

func httpError(w http.ResponseWriter, status int, err string) {
	w.Header().Set("Content-Type", "text/plain")
	if status == http.StatusUnauthorized {
		w.Header().Set("WWW-Authenticate", "Basic")
	}
	http.Error(w, err, status)
	return
}

func DefaultUnauthorized(rp *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func (w http.ResponseWriter, r *http.Request){
		httpError(w, 403, "Unauthorized\n")
	}
}