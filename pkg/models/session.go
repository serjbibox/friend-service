package models

//@Description Структура данных сессии с психологом
type Session struct {
	ID       int    `json:"id" example:"11234"`
	ClientID int    `json:"client_id" example:"123"`
	PsyID    int    `json:"psy_id" example:"2345"`
	Date     string `json:"date" example:"2000-07-28"`
	Type     string `json:"type" example:"video(audio, chat, etc..)"`
	Note     string `json:"session_notes" example:"video url"`
	Created  string `json:"created_at" example:"01.01.2023 22:00"`
	Updated  string `json:"updated_at" example:"02.01.2023 22:00"`
}
