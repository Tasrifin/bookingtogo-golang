package services

import (
	"fmt"
	"net/http"

	"github.com/Tasrifin/bookingtogo-golang/params"
	"github.com/Tasrifin/bookingtogo-golang/repositories"
)

type CustomerService struct {
	customerRepo repositories.CustomerRepo
}

func NewBattleService(customerRepo repositories.CustomerRepo) *CustomerService {
	return &CustomerService{
		customerRepo: customerRepo,
	}
}

func (c *CustomerService) GetAllCustomers() *params.Response {
	data, err := c.customerRepo.GetAllCustomers()

	if err != nil {
		return &params.Response{
			Status: http.StatusBadRequest,
			Payload: map[string]interface{}{
				"error": err.Error(),
			},
		}
	}

	result := []params.Customer{}

	for _, val := range *data {
		customer := params.Customer{
			CustomerID:          val.CustomerId,
			CustomerName:        val.CustomerName,
			CustomerDoB:         val.CustomerDoB[0:10],
			CustomerPhoneNumber: val.CustomerPhoneNumber,
			CustomerEmail:       val.CustomerEmail,
			Nationality:         fmt.Sprintf("%v (%v)", val.Nationality.NationalityName, val.Nationality.NationalityCode),
		}

		families := []params.FamilyList{}
		for _, fam := range val.FamilyList {
			customerFamilyList := params.FamilyList{
				FamilyName:     fam.FamilyName,
				FamilyRelation: fam.FamilyRelation,
				FamilyDoB:      fam.FamilyDoB,
			}
			families = append(families, customerFamilyList)
		}

		customer.FamilyList = families
		result = append(result, customer)
	}

	return &params.Response{
		Status: http.StatusOK,
		Payload: map[string]interface{}{
			"data": result,
		},
	}
}

func (c *CustomerService) GetCustomerByID(id int) *params.Response {
	data, err := c.customerRepo.GetCustomerByID(id)

	if err != nil {
		return &params.Response{
			Status: http.StatusBadRequest,
			Payload: map[string]interface{}{
				"error": err.Error(),
			},
		}
	}

	result := params.Customer{
		CustomerName:        data.CustomerName,
		CustomerDoB:         data.CustomerDoB[0:10],
		CustomerPhoneNumber: data.CustomerPhoneNumber,
		CustomerEmail:       data.CustomerEmail,
		Nationality:         fmt.Sprintf("%v (%v)", data.Nationality.NationalityName, data.Nationality.NationalityCode),
	}

	families := []params.FamilyList{}
	for _, fam := range data.FamilyList {
		customerFamilyList := params.FamilyList{
			FamilyName:     fam.FamilyName,
			FamilyRelation: fam.FamilyRelation,
			FamilyDoB:      fam.FamilyDoB,
		}
		families = append(families, customerFamilyList)
	}

	result.FamilyList = families

	return &params.Response{
		Status: http.StatusOK,
		Payload: map[string]interface{}{
			"data": result,
		},
	}
}
