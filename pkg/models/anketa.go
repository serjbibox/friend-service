package models

//@Description Структура данных анкеты пользователя
type Anketa struct {
	ID       int    `json:"id" example:"11234"`
	ClientID int    `json:"client_id" example:"123"`
	Age      int    `json:"age" example:"35"`
	Sex      string `json:"sex" example:"female"`
	Case     string `json:"problem" example:"описание проблемы"`
	Created  string `json:"created_at" example:"01.01.2023 22:00"`
	Updated  string `json:"updated_at" example:"02.01.2023 22:00"`
}
