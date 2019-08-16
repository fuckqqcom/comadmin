package admin

/**
user session 用户session时长
*/

type UserSession struct {
	Uid     string `json:"uid"`
	Web     int    `json:"web"`
	Android int    `json:"android"`
	Ios     int    `json:"ios"`
}
