package models

type Nationality struct {
	NationalityID   int    `gorm:"primaryKey" json:"nationality_id"`
	NationalityName string `json:"nationality_name"`
	NationalityCode string `json:"nationality_code"`
}
