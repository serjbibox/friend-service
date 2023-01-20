package models

// @Description Структура карточки пользователя в теле запроса POST /v1/auth/register
type User struct {
	ID       int    `json:"id" example:"11234"`
	Role     string `json:"role" example:"client"`
	Phone    string `json:"phone" example:"+79167003020"`
	Login    string `json:"login" example:"rubella19"`
	Password string `json:"password" example:"1Qwerty!"`
	Name     string `json:"name" example:"Анастасия"`
	Birth    string `json:"birth" example:"2000-07-28"`
	Tag      string `json:"tg" example:"@Rubella19"`
	Video    string `json:"video_url" example:"videourl"`
	Email    string `json:"email" example:"some.mail@gmail.com"`
}
