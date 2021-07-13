package Service

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

// DecodeUserRequest
func DecodeUserRequest(ctx context.Context, r * http.Request)(interface{}, error) {
	uid := r.URL.Query().Get("uid")
	if uid != "" {
		uid, _ := strconv.Atoi(uid)
		return UserRequest{
			Uid: uid,
		}, nil
	}
	return nil, errors.New("参数错误")
}

// EncodeUserResponse
func EncodeUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}