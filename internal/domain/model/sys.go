package model

type LoginSuccess struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Token   string `json:"token"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Code     string `json:"code"`
	Uuid     string `json:"uuid"`
}
