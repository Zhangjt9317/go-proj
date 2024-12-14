package models

type Poll struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Question string `json:"question"`
	Options []Option `json:"options" gorm:"foreignKey:PollID"`
}

type Option struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Text string `json:"text"`
	Votes int `json:"votes"`
	PollID uint `json:"poll_id"`
}