package main

import (
	goKitHttp "github.com/go-kit/kit/transport/http"
	. "go-kit/Service"
	"net/http"
)


func main()  {
	serviceHandler := goKitHttp.NewServer(GenUserEndpoint(UserService{}), DecodeUserRequest, EncodeUserResponse)
	_ = http.ListenAndServe(":8080", serviceHandler)
}
