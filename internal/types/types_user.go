package mestoTypes

type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	About  string `json:"about"`
	Avatar string `json:"avatar"`
	Cohort string `json:"cohort"`
}

type UserMe struct {
	Name  string `json:"name" binding:"required"`
	About string `json:"about" binding:"required"`
}

type UserAuth struct {
	Authorization string `json:"authorization" binding:"required"`
}

type UserAvatar struct {
	Avatar string `json:"avatar" binding:"required"`
}
