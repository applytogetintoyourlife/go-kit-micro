package Service

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	goKitMux "github.com/gorilla/mux"
)

// DecodeUserRequest
func DecodeUserRequest(ctx context.Context, r * http.Request)(interface{}, error) {

	vars := goKitMux.Vars(r)
	if uid, ok := vars["uid"]; ok {
		uidToInt, _ := strconv.Atoi(uid)
		return UserRequest{
			Uid: uidToInt,
			Method: r.Method,
		}, nil
	}
	return nil, errors.New("参数错误")
}

// EncodeUserResponse
func EncodeUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-type", "application/json")
	return json.NewEncoder(w).Encode(response)
}