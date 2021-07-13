package Service

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

// UserRequest
type UserRequest struct {
	Uid int `json:"uid"`
}

// UserResponse
type UserResponse struct {
	Result string `json:"result"`
}

// GenUserEndpoint
func GenUserEndpoint(userService IUserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r := request.(UserRequest)
		result := userService.GetName(r.Uid)
		return UserResponse{Result: result}, nil
	}
}