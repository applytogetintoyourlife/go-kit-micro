package main

import (
	goKitHttp "github.com/go-kit/kit/transport/http"
	goKitMux "github.com/gorilla/mux"
	. "go-kit/Service"
	"net/http"
)


func main()  {
	serviceHandler := goKitHttp.NewServer(GenUserEndpoint(UserService{}), DecodeUserRequest, EncodeUserResponse)
	r := goKitMux.NewRouter()
	// r.Handle(`/user/(uid:\$d+)`, serviceHandler)
	r.Methods("GET", "DELETE").Path(`/user/(uid:\$d+)`).Handler(serviceHandler)
	_ = http.ListenAndServe(":8080", r)
}
