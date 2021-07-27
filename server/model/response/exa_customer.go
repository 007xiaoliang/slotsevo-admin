package response

import "slotsevo-admin/model"

type ExaCustomerResponse struct {
	Customer model.ExaCustomer `json:"customer"`
}
