package controllers

import (
	"devops-api/common"
)

var (
	queryPhoneEntryType = "Query Phone Location"
)

// Get Get方法
func (p *PhoneController) Get() {
	phone := p.GetString("phone")
	m, err := common.queryPhone(phone)
	if err!=nil{
		p.jsonerr
	}
}
