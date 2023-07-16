package models

type Customer struct {
	CustomerId          int          `gorm:"primaryKey;column:cst_id" json:"cst_id"`
	NationalityID       int          `gorm:"column:nationality_id" json:"nationality_id"`
	CustomerName        string       `gorm:"column:cst_name" json:"cst_name"`
	CustomerDoB         string       `gorm:"column:cst_dob" json:"cst_dob"`
	CustomerPhoneNumber string       `gorm:"column:cst_phoneNum" json:"cst_phoneNum"`
	CustomerEmail       string       `gorm:"column:cst_email" json:"cst_email"`
	Nationality         *Nationality `json:"nationality"`
	FamilyList          []FamilyList `json:"families"`
}
