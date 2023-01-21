package models

// @Description Структура данных пользователя
type User struct {
	ID       int    `json:"id" example:"11234"`
	Role     string `json:"role" example:"client (psychologist, moderator)"`
	Phone    string `json:"phone" example:"+79161234567"`
	Login    string `json:"login" example:"somelogin"`
	Password string `json:"password" example:"1Qwerty!"`
	Name     string `json:"name" example:"Анастасия"`
	Age      int    `json:"age" example:"35"`
	Tag      string `json:"tag" example:"@user"`
	Video    string `json:"video_url" example:"video url"`
	Email    string `json:"email" example:"some.mail@gmail.com"`
	Anketa   Anketa `json:"anketa"`
}
