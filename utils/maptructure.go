package utils

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"go-ginapp/data/response"
	"go-ginapp/models"
	"reflect"
	"time"
)

// ToUserResponse map user to user-response using mapstructure
func ToUserResponse(user models.User) (response.UserResponse, error) {
	var userResponse response.UserResponse
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result:  &userResponse,
		TagName: "json",
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			func(from, to reflect.Type, data interface{}) (interface{}, error) {

				if from == reflect.TypeOf(&time.Time{}) && to == reflect.TypeOf("") {
					if t, ok := data.(*time.Time); ok {
						if t == nil {
							return "", nil
						}
						return t.Format("2006-01-02 15:04:05"), nil
					}
				}
				if from == reflect.TypeOf(time.Time{}) && to == reflect.TypeOf("") {
					if t, ok := data.(time.Time); ok {
						return t.Format("2006-01-02 15:04:05"), nil
					}
				}
				if from == reflect.TypeOf(map[string]interface{}{}) && to == reflect.TypeOf("") {
					return "", nil
				}
				return data, nil
			},
		),
	})
	if err != nil {
		return response.UserResponse{}, err
	}
	fmt.Printf("oke 1")
	if err := decoder.Decode(user); err != nil {
		return response.UserResponse{}, err
	}
	userResponse.RoleName = user.Role.Name
	return userResponse, nil
}

func ToPostResponse(post models.Post) (response.PostResponse, error) {
	var postResponse response.PostResponse
	// tao new decoder de decode cac truong kieu time sang string
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result:  &postResponse,
		TagName: "json",
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			func(from, to reflect.Type, data interface{}) (interface{}, error) {
				if from == reflect.TypeOf(&time.Time{}) && to == reflect.TypeOf("") {
					if t, ok := data.(*time.Time); ok {
						return t.Format(time.DateTime), nil
					}
				}
				return data, nil
			},
		),
	})

	if err != nil {
		return response.PostResponse{}, err
	}
	if err := decoder.Decode(post); err != nil {
		return response.PostResponse{}, err
	}

	userResponse, err := ToUserResponse(post.User)
	if err != nil {
		return response.PostResponse{}, err
	}
	postResponse.User = userResponse

	return postResponse, nil
}
