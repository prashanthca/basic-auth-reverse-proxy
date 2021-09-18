package strategies

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"github.com/deta/deta-go/deta"
	"github.com/deta/deta-go/service/base"
)

type User struct {
	Key string `json:"key"`
	User string `json:"user"`
	Password string `json:"pass"`
}

type Deta struct {
	DetaProjectKey string
	DetaBaseId string
}

func (x Deta) RequestHandler(rp *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		username, pass, ok := r.BasicAuth()
		if !ok {
			httpError(w, http.StatusUnauthorized, "Not authorized\n")
			return 
		}
		d, err := deta.New(deta.WithProjectKey(x.DetaProjectKey))
		if err != nil {
			httpError(w, http.StatusForbidden, fmt.Sprintf("failed to init new Deta instance: %s\n", err))
			return
		}
		db, err := base.New(d, x.DetaBaseId)
		if err != nil {
			httpError(w, http.StatusForbidden, fmt.Sprintf("failed to init new Base instance: %s\n", err))
			return
		}
		var results []*User
		_, err = db.Fetch(&base.FetchInput{
			Q: base.Query{
				{"user": username},
			},
			Limit: 1,
			Dest: &results,
		})
		if err != nil {
			httpError(w, http.StatusForbidden, fmt.Sprintf("Failed to get item: %s\n", err))
			return
		}
		if pass != results[0].Password {
			httpError(w, http.StatusForbidden, "Unauthorized\n")
			return
		}
		rp.ServeHTTP(w, r)
	}
}