package models

type Anketa struct {
	ID       int    `json:"id" example:"11234"`
	ClientID int    `json:"phone" example:"+79167003020"`
	Age      int    `json:"login" example:"rubella19"`
	Sex      string `json:"password" example:"1Qwerty!"`
	Case     string `json:"name" example:"Анастасия"`
	Created  string `json:"tg" example:"@Rubella19"`
	Updated  string `json:"email" example:"anastasia.a.krasnova@gmail.com"`
}
