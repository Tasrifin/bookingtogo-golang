package repositories

import (
	"github.com/Tasrifin/bookingtogo-golang/models"
	"gorm.io/gorm"
)

type CustomerRepo interface {
	GetAllCustomers() (*[]models.Customer, error)
	GetCustomerByID(id int) (*models.Customer, error)
}

type customerRepo struct {
	db *gorm.DB
}

func NewCustomerRepo(db *gorm.DB) CustomerRepo {
	return &customerRepo{db}
}

func (c *customerRepo) GetAllCustomers() (*[]models.Customer, error) {
	var customer []models.Customer
	err := c.db.Preload("Nationality").Preload("FamilyList").Find(&customer).Error
	return &customer, err
}

func (c *customerRepo) GetCustomerByID(id int) (*models.Customer, error) {
	var customer models.Customer
	err := c.db.Preload("Nationality").Preload("FamilyList").First(&customer, "cst_id= ?", id).Error
	return &customer, err
}
