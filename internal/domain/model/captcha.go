package model

type CaptchaImage struct {
	Code           int    `json:"code"`
	Message        string `json:"msg"`
	Uuid           string `json:"uuid"`
	CaptchaEnabled bool   `json:"captchaEnabled,omitempty"`
	Img            string `json:"img"`
}
