package permission

/**
user session 用户session时长
*/

type UserSession struct {
	ID      int64
	Uid     string `json:"uid"`
	Web     int    `json:"web"`
	Android int    `json:"android"`
	Ios     int    `json:"ios"`
}
