package user

type User struct {
	ID     int    `json:"_id"`
	Name   string `json:"name"`
	About  string `json:"about"`
	Avatar string `json:"avatar"`
}

type UserMe struct {
	Name  string `json:"name" binding:"required"`
	About string `json:"about" binding:"required"`
}

type UserAuth struct {
	Authorization string `json:"authorization" binding:"required"`
}

type UserId string

// func UserFromString(id string, name string, about string, avatar string) (*User, err) {
// 	if name != "" {
// 		if common.IsImage(avatar) {
// 			return &User{}
// 		} else {
// 			log.Println(avatar, "is not image!")
// 		}
// 	}

// }
