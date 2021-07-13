package Service

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
)

// UserRequest
type UserRequest struct {
	Uid int `json:"uid"`
	Method string
}

// UserResponse
type UserResponse struct {
	Result string `json:"result"`
}

// GenUserEndpoint
func GenUserEndpoint(userService IUserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r := request.(UserRequest)
		result := "nothing"
		if r.Method == "GET" {
			result = userService.GetName(r.Uid)
		}
		if r.Method == "DELETE" {
			delErr := userService.DelUser(r.Uid)
			if delErr != nil {
				result = delErr.Error()
			} else {
				result = fmt.Sprintf("删除成功,uid: %d", r.Uid)
			}
		}
		return UserResponse{Result: result}, nil
	}
}