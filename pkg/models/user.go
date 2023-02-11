package models

// @Description Структура данных пользователя
type User struct {
	ID       int    `json:"id" example:"11234"`
	Role     string `json:"role" example:"client (psychologist, moderator)"`
	Phone    string `json:"phone" example:"+79161234567"`
	Password string `json:"password" example:"1Qwerty!"`
	Email    string `json:"email" example:"some.mail@gmail.com"`
	Name     string `json:"name" example:"Анастасия"`
	Age      int    `json:"age" example:"35"`
	Sex      string `json:"sex" example:"female"`
	Icon     string `json:"icon" example:"url to icon"`
}

type UserIdentity struct {
	Phone    string `json:"phone" example:"+79161234567"`
	SmsCode  int    `json:"sms_code" example:"123456"`
	Role     string `json:"role" example:"client (psychologist, moderator)"`
	Password string `json:"password" example:"1Qwerty!"`
}
