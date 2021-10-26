package httpExample

import (
	"fmt"
	"github.com/pkg/errors"
)

type users interface {
	UpsertUserData(userUUID string, data UserData) (UserData, error)
	GetUserData(userUUID string) (UserData, error)

	LikeProduct(userUUID string, dishId int64) (UserData, error)
	DislikeProduct(userUUID string, dishId int64) (UserData, error)
}

type UserData struct {
	Favorites Favorites `json:"favorites"`
}

type Favorites struct {
	ProductIds []int64 `json:"productIds"`
}

func (api *exampleApi) UpsertUserData(userUUID string, userData UserData) (UserData, error) {
	var response UserData
	if err := api.Post(api.userDataPath(userUUID), userData, &response); err != nil {
		return UserData{}, errors.Wrap(err, "upsert_user_data")
	}
	return response, nil
}

func (api *exampleApi) GetUserData(userUUID string) (UserData, error) {
	var response UserData
	if err := api.Get(api.userDataPath(userUUID), &response); err != nil {
		return UserData{}, errors.Wrap(err, "get_user_data")
	}
	return response, nil
}

func (api *exampleApi) LikeProduct(userUUID string, productId int64) (UserData, error) {
	var response UserData
	if err := api.Post(api.userProductPath(userUUID, productId), nil, &response); err != nil {
		return UserData{}, errors.Wrap(err, "like_product")
	}
	return response, nil
}

func (api *exampleApi) DislikeProduct(userUUID string, productId int64) (UserData, error) {
	var response UserData
	if err := api.Delete(api.userProductPath(userUUID, productId), &response); err != nil {
		return UserData{}, errors.Wrap(err, "dislike_product")
	}
	return response, nil
}

func (api *exampleApi) userDataPath(userUUID string) string {
	return fmt.Sprintf("/user/%s/settings", userUUID)
}

func (api *exampleApi) userProductPath(userUUID string, productID int64) string {
	return fmt.Sprintf("/user/%s/settings/product/%v", userUUID, productID)
}
