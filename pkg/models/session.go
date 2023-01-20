package models

type Session struct {
	ID       int    `json:"id" example:"11234"`
	ClientID int    `json:"phone" example:"+79167003020"`
	PsyID    int    `json:"login" example:"rubella19"`
	Date     string `json:"password" example:"1Qwerty!"`
	Type     string `json:"name" example:"Анастасия"`
	Note     string `json:"birth" example:"2000-07-28"`
	Created  string `json:"tg" example:"@Rubella19"`
	Updated  string `json:"email" example:"anastasia.a.krasnova@gmail.com"`
}
