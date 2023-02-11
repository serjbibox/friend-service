package models

//@Description Структура данных анкеты клиента
type ClientData struct {
	ID      int    `json:"id" example:"11234"`
	UserID  int    `json:"user_id" example:"123"`
	Case    string `json:"problem" example:"описание проблемы"`
	Created string `json:"created_at" example:"01.01.2023 22:00"`
	Updated string `json:"updated_at" example:"02.01.2023 22:00"`
}

//@Description Структура данных анкеты психолога
type PsyData struct {
	ID      int    `json:"id" example:"11234"`
	UserID  int    `json:"user_id" example:"123"`
	Intro   string `json:"intro" example:"резюме психолога"`
	Video   string `json:"video_url" example:"video url"`
	Created string `json:"created_at" example:"01.01.2023 22:00"`
	Updated string `json:"updated_at" example:"02.01.2023 22:00"`
}
