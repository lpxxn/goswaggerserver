// go-swagger examples.
//
// The purpose of this application is to provide some
// use cases describing how to generate docs for your API
//
//     Schemes: http, https
//     Host: localhost
//     BasePath: /
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package goswaggerserver

import (
	"net/http"

	"github.com/lpxxn/goswaggerserver/serverapi"
)

func main() {
	http.HandleFunc("querydata", serverapi.QueryData)
	http.HandleFunc("saveuser", serverapi.SaveUserInfo)

	http.ListenAndServe(":9999", nil)
}




