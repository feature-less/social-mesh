# Authz

Authz is an authorizationd middleware for [chi](https://github.com/go-chi/chi), and is based on [casbin](https://github.com/casbin/casbin)

## Installaton

all you have to do is import it

## Example

```Go
package main

import(
	"net/http"

	"authz"
	"github.com/casbin/casbin/v2"
	"github.com/go-chi/chi/v5"
)

func main(){
	router := chi.newRouter()

	// create your enforcer

	enforcer := casbin.NewEnforcer("./model.conf", "./policy.csv")
	router.use(authz.Authorizor(enforcer))

	router.HandleFunc("/*", func(writer http.ResponseWriter, req *http.Request){
		writer.WriteHeader(200)
	})
}
```

## Getting Help

- [casbin](https://github.com/casbin/casbin)
- [chi](https://github.com/go-chi/chi)
- [Oussama M. Bouchareb](mailto://commensalism@proton.me)
