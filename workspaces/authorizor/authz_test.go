package authz

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/casbin/casbin/v2"
	"github.com/go-chi/chi/v5"
)

func testAuthzRequest(tester *testing.T, router *chi.Mux, user string, path string, method string, code int) {
	req, _ := http.NewRequest(method, path, nil)
	req.SetBasicAuth(user, "test")
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, req)

	if writer.Code != code {
		tester.Errorf("%s, %s, %s:, %d supposed to be %d", user, path, method, writer.Code, code)
	}
}

func TestBasic(tester *testing.T) {
	log.SetPrefix("[Authz Testing]: ")

	router := chi.NewRouter()

	enforcer, err := casbin.NewEnforcer("model.conf", "policy.csv")

	if err != nil {
		log.Fatalf("Encounter an Error while trying to create an new enforcer:\n %s", err)
	}

	router.Use(Authorizor(enforcer))

	router.HandleFunc("/*", func(writer http.ResponseWriter, req *http.Request) {
		writer.WriteHeader(200)
	})

	testAuthzRequest(tester, router, "alice", "/resource1", "GET", 200)
	testAuthzRequest(tester, router, "alice", "/resource2", "GET", 200)
	testAuthzRequest(tester, router, "alice", "/resource3", "GET", 200)
	testAuthzRequest(tester, router, "alice", "/resource4", "GET", 403)
}

func TestPathWildCard(tester *testing.T) {
	log.SetPrefix("[Authz Testing]: ")

	router := chi.NewRouter()

	enforcer, err := casbin.NewEnforcer("./model.conf", "./policy.csv")

	if err != nil {
		log.Fatalf("Encounter an Error while trying to create an new enforcer:\n %s", err)
	}

	router.Use(Authorizor(enforcer))

	router.HandleFunc("/*", func(writer http.ResponseWriter, req *http.Request) {
		writer.WriteHeader(200)
	})

	testAuthzRequest(tester, router, "bob", "/resource1", "GET", 200)
	testAuthzRequest(tester, router, "bob", "/resource2", "GET", 200)
	testAuthzRequest(tester, router, "bob", "/resource3", "GET", 200)
	testAuthzRequest(tester, router, "bob", "/resource4", "GET", 200)
}
