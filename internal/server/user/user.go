package user

import (
	"mesto-goback/internal/common"
)

type User struct {
	Name   string           `json:"name"`
	About  string           `json:"about"`
	Avatar common.ImageLink `json:"avatar"`
	ID     UserId           `json:"_id"`
}

type UserMe struct {
	Name  string `json:"name" binding:"required"`
	About string `json:"about" binding:"required"`
}

type UserAuth struct {
	Authorization string `json:"authorization" binding:"required"`
}

type UserId string
