package app

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/Nick212/go-crm-espo/models/crm"
)

const (
	RESOURCE string = "Account"
)

func (a *App) Get(id string) (*crm.Contact, error) {
	var model crm.Contact
	response, err := a.CRMHandlerGetService(RESOURCE, id)

	if err == nil {
		err = json.Unmarshal(response, &model)
	}

	return &model, err
}

func (a *App) CRMGetAccountAll() (*crm.ResponseAccount, error) {
	var account crm.ResponseAccount

	response, err := a.CRMHandlerGetService(RESOURCE, "")

	if err == nil {
		err = json.Unmarshal(response, &account)
	}

	return &account, err
}

func (a *App) CRMGetAccountByFilter(filter string) (*crm.ResponseAccount, error) {
	var account crm.ResponseAccount
	params := filter

	response, err := a.CRMHandlerGetService(RESOURCE, params)

	if err == nil {
		err = json.Unmarshal(response, &account)
	}

	return &account, err
}

func (a *App) CRMGetAccountByWithoutExternalIdAndCnpj() (*crm.ResponseAccount, error) {
	var responseAccount crm.ResponseAccount
	params := "?select=id&id&where%5B0%5D%5Btype%5D=or&where%5B0%5D%5Bvalue%5D%5B0%5D%5Btype%5D=isNull&where%5B0%5D%5Bvalue%5D%5B0%5D%5Battribute%5D=externalID&where%5B0%5D%5Bvalue%5D%5B1%5D%5Btype%5D=equals&where%5B0%5D%5Bvalue%5D%5B1%5D%5Battribute%5D=externalID&where%5B0%5D%5Bvalue%5D%5B1%5D%5Bvalue%5D=&where%5B1%5D%5Btype%5D=and&where%5B1%5D%5Bvalue%5D%5B0%5D%5Btype%5D=notEquals&where%5B1%5D%5Bvalue%5D%5B0%5D%5Battribute%5D=sicCode&where%5B1%5D%5Bvalue%5D%5B0%5D%5Bvalue%5D=&where%5B1%5D%5Bvalue%5D%5B1%5D%5Btype%5D=isNotNull&where%5B1%5D%5Bvalue%5D%5B1%5D%5Battribute%5D=sicCode&where%5B1%5D%5Bvalue%5D%5B1%5D%5Bvalue%5D="

	response, err := a.CRMHandlerGetService(RESOURCE, params)

	if err == nil {
		err = json.Unmarshal(response, &responseAccount)
	}

	return &responseAccount, err
}

func (a *App) CRMGetAccountByWithoutExternalId() (*crm.ResponseAccount, error) {
	var responseAccount crm.ResponseAccount
	params := "?select=id&where%5B0%5D%5Btype%5D=or&where%5B0%5D%5Bvalue%5D%5B0%5D%5Btype%5D=isNull&where%5B0%5D%5Bvalue%5D%5B0%5D%5Battribute%5D=externalID&where%5B0%5D%5Bvalue%5D%5B1%5D%5Btype%5D=equals&where%5B0%5D%5Bvalue%5D%5B1%5D%5Battribute%5D=externalID&where%5B0%5D%5Bvalue%5D%5B1%5D%5Bvalue%5D=&where%5B1%5D%5Btype%5D=and&where%5B1%5D%5Bvalue%5D%5B0%5D%5Btype%5D=notEquals&where%5B1%5D%5Bvalue%5D%5B0%5D%5Battribute%5D=sicCode&where%5B1%5D%5Bvalue%5D%5B0%5D%5Bvalue%5D=&where%5B1%5D%5Bvalue%5D%5B1%5D%5Btype%5D=isNotNull&where%5B1%5D%5Bvalue%5D%5B1%5D%5Battribute%5D=sicCode&where%5B1%5D%5Bvalue%5D%5B1%5D%5Bvalue%5D="

	response, err := a.CRMHandlerGetService(RESOURCE, params)

	if err == nil {
		err = json.Unmarshal(response, &responseAccount)
	}

	return &responseAccount, err
}

func (a *App) CRMGetAccountByExternalIdEqualZero() (*crm.ResponseAccount, error) {
	var responseAccount crm.ResponseAccount
	params := "?select=id&where%5B0%5D%5Btype%5D=equals&where%5B0%5D%5Battribute%5D=externalID&where%5B0%5D%5Bvalue%5D=0&where%5B1%5D%5Btype%5D=and&where%5B1%5D%5Bvalue%5D%5B0%5D%5Btype%5D=notEquals&where%5B1%5D%5Bvalue%5D%5B0%5D%5Battribute%5D=sicCode&where%5B1%5D%5Bvalue%5D%5B0%5D%5Bvalue%5D=&where%5B1%5D%5Bvalue%5D%5B1%5D%5Btype%5D=isNotNull&where%5B1%5D%5Bvalue%5D%5B1%5D%5Battribute%5D=sicCode&where%5B1%5D%5Bvalue%5D%5B1%5D%5Bvalue%5D="
	response, err := a.CRMHandlerGetService(RESOURCE, params)

	if err == nil {
		err = json.Unmarshal(response, &responseAccount)
	}

	return &responseAccount, err
}

func (a *App) CRMGetAccountByStatus(status string) (*crm.ResponseAccount, error) {
	offset := 0
	maxSize := 100
	total := 100
	var responseAccount crm.ResponseAccount

	var res crm.ResponseAccount

	for offset < total {
		pagination := ("maxSize=" + strconv.Itoa(maxSize) + "&offset=" + strconv.Itoa(offset))
		params := "?select=id,sicCode,erpCode&" + pagination + "&where%5B0%5D%5Btype%5D=in&where%5B0%5D%5Battribute%5D=statusExternal&where%5B0%5D%5Bvalue%5D%5B%5D=" + status
		response, _ := a.CRMHandlerGetService(RESOURCE, params)

		_ = json.Unmarshal(response, &res)

		total = res.Total
		offset = offset + maxSize
		responseAccount.Total = res.Total
		responseAccount.Account = append(responseAccount.Account, res.Account...)
	}

	return &responseAccount, nil
}

func (a *App) CRMGetAccountParentByTeam(teamId string) (*crm.ResponseAccount, error) {
	var responseAccount crm.ResponseAccount
	params := "?select=sicCode,name,phoneNumberIsOptedOut,phoneNumber,phoneNumberData,emailAddressIsOptedOut,emailAddress,emailAddressData,assignedUserId,assignedUserName&maxSize=20&offset=0&orderBy=createdAt&order=desc&where[0][type]=linkedWith&where[0][attribute]=teams&where[0][value][]=" + teamId + "&where[1][type]=isTrue&where[1][attribute]=isPartner"

	response, err := a.CRMHandlerGetService(RESOURCE, params)

	if err == nil {
		err = json.Unmarshal(response, &responseAccount)
	}

	return &responseAccount, err
}

func (a *App) CRMGetAccountResellerBySicCode(codeId string) (*crm.Account, error) {
	var responseAccount *crm.ResponseAccount
	params := "?where%5B0%5D%5Btype%5D=equals&where%5B0%5D%5Battribute%5D=sicCode&where%5B0%5D%5Bvalue%5D=" + codeId + "&where%5B1%5D%5Btype%5D=in&where%5B1%5D%5Battribute%5D=type&where%5B1%5D%5Bvalue%5D%5B%5D=Parceiro"
	response, err := a.CRMHandlerGetService(RESOURCE, params)

	if err == nil {
		err = json.Unmarshal(response, &responseAccount)
	}

	if responseAccount.Total > 0 {
		return &responseAccount.Account[0], err
	}

	return nil, err
}

func (a *App) CRMGetAccountBySicCode(codeId, typeAccount string) (*crm.Account, error) {
	var responseAccount crm.ResponseAccount

	if codeId == "" || strings.ToLower(codeId) == "null" {
		return &crm.Account{}, nil
	}

	params := "?where%5B0%5D%5Btype%5D=equals&where%5B0%5D%5Battribute%5D=sicCode&where%5B0%5D%5Bvalue%5D=" + codeId

	response, err := a.CRMHandlerGetService(RESOURCE, params)

	if err == nil {
		err = json.Unmarshal(response, &responseAccount)
	}

	if len(responseAccount.Account) > 0 {
		for _, account := range responseAccount.Account {
			if strings.ToLower(account.Type) == strings.ToLower(typeAccount) {
				return &account, err
			}
		}
	}

	return nil, err
}

func (a *App) CRMGetAccountByExternalId(externalId string) (*crm.Account, error) {
	var responseAccount crm.ResponseAccount
	params := "?where%5B0%5D%5Btype%5D=equals&where%5B0%5D%5Battribute%5D=externalID&where%5B0%5D%5Bvalue%5D=" + externalId

	response, err := a.CRMHandlerGetService(RESOURCE, params)

	if err == nil {
		err = json.Unmarshal(response, &responseAccount)
	}

	if len(responseAccount.Account) > 0 {
		return &responseAccount.Account[0], err
	}

	return nil, err
}

func (a *App) CRMGetAccountById(id string) (*crm.Account, error) {
	var account crm.Account
	params := "/" + id

	response, err := a.CRMHandlerGetService(RESOURCE, params)

	if err == nil {
		err = json.Unmarshal(response, &account)
	}

	return &account, err
}

func (a *App) CRMUpdateAccount(model crm.Account) (*crm.Account, error) {
	var account crm.Account
	payload, _ := json.Marshal(model)
	response, err := a.CRMHandlerPutService(RESOURCE+"/"+model.ID, payload)

	if err == nil {
		err = json.Unmarshal(response, &account)
	}

	return &account, err
}

func (a *App) CRMUpdateAccountRelationshipAccount(model crm.Account, parent crm.AccountParentId) (*crm.Account, error) {
	var account crm.Account
	payload, _ := json.Marshal(parent)
	response, err := a.CRMHandlerPutService(RESOURCE+"/"+model.ID, payload)

	if err == nil {
		err = json.Unmarshal(response, &account)
	}

	return &account, err
}

func (a *App) CRMUpdateStatusAccount(id string, status string) (res *crm.Account, err error) {
	if id == "" || status == "" {
		return nil, err
	}

	var account crm.Account
	var model = crm.StatusExternal{}
	model.Status = status

	payload, err := json.Marshal(model)
	response, err := a.CRMHandlerPutService(RESOURCE+"/"+id, payload)

	if err == nil {
		err = json.Unmarshal(response, &account)
	}

	return &account, err
}

func (a *App) CRMSaveAccount(model crm.Account) (*crm.Account, error) {
	var account crm.Account
	resource := RESOURCE
	payload, _ := json.Marshal(model)

	response, err := a.CRMHandlerPostService(resource, payload)

	if err == nil {
		err = json.Unmarshal(response, &account)
	}

	return &account, err
}
