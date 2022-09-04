// an authorization middleware for chi based on casbin
// https://github.com/casbin/casbin

package authz

import (
	"log"
	"net/http"

	"github.com/casbin/casbin/v2"
)

// this middleware does not authenticate the use it only authorizes them
// user: the authenticated user
// path: the resource in which the use performed an action against
// method: the http method in which the user perform their action
//
// this middleware should be implemented as early as possible
// but only after you confirm the encing is correct, the ratelimits ...etc
func Authorizor(enforcer *casbin.Enforcer) func(next http.Handler) http.Handler {
	log.SetPrefix("[Authz]: ")
	return func(next http.Handler) http.Handler {
		fn := func(writer http.ResponseWriter, req *http.Request) {
			// used to get the role of the user
			role := req.Header.Get("role")

			// make sure role is valid
			if role == "" || (role != "admin" && role != "mod" && role != "member") {
				role = "anonymouse"
			}

			// used to extract the http method in which the user tried to
			// access the resource
			method := req.Method

			// the resource path in which an action was performed against
			path := req.URL.Path

			// store enforcer results
			enforcerResult, err := enforcer.Enforce(role, path, method)

			// check if enforcer enfountered any errors
			if err != nil {
				http.Error(writer, http.StatusText(500), 500)
				log.Println(err)
			}

			// check against our enforcer
			if enforcerResult == true {
				// if successful we let the action continue to go down
				// to the resource
				next.ServeHTTP(writer, req)
			} else {
				// if unsuccessful we return an unauthorized action status
				// 403
				http.Error(writer, http.StatusText(403), 403)
			}

		}
		return http.HandlerFunc(fn)
	}
}
