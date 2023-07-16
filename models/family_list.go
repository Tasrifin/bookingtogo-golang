package models

type FamilyList struct {
	FamilyID       int    `gorm:"primaryKey;column:fl_id" json:"fl_id"`
	CustomerID     int    `gorm:"column:cst_id" json:"cst_id"`
	FamilyRelation string `gorm:"column:fl_relation" json:"fl_relation"`
	FamilyName     string `gorm:"column:fl_name" json:"fl_name"`
	FamilyDoB      string `gorm:"column:fl_dob" json:"fl_dob"`
}
