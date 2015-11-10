package main

import (
	"net/http"

	"github.com/drhayt/arith"
	"github.com/drhayt/utility"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc/v2/json2"
)

func main() {
	// s := rpcv2.NewServer()
	s := rpc
	s.RegisterCodec(json2.NewCodec(), "application/json")
	s.RegisterCodec(json2.NewCodec(), "application/json;charset=UTF-8")
	arith := new(arith.Arith)
	utility := new(utility.Utility)
	s.RegisterService(arith, "")
	s.RegisterService(utility, "")
	r := mux.NewRouter()
	r.Handle("/rpc", s)
	http.ListenAndServe(":1234", r)
}
