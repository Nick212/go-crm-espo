package app

import (
	"encoding/json"

	"github.com/Nick212/go-crm-espo/models/crm"
)

const (
	RESOURCE string = "Contact"
)

func (a *App) Get(id string) (*crm.Contact, error) {
	var model crm.Contact
	response, err := a.CRMHandlerGetService(RESOURCE, id)

	if err == nil {
		err = json.Unmarshal(response, &model)
	}

	return &model, err
}

func (a *App) CRMGetOrSaveContactByEmail(email string) (*crm.Contact, error) {
	var model crm.ResposeContact
	resource := "Contact"
	params := "?where%5B0%5D%5Btype%5D=equals&where%5B0%5D%5Battribute%5D=emailAddress&where%5B0%5D%5Bvalue%5D=" + email

	response, err := a.CRMHandlerGetService(RESOURCE, params)

	if err == nil {
		err = json.Unmarshal(response, &model)
	}

	if model.Total == 0 {
		return &crm.Contact{}, err
	}

	//todo refactor
	return &model.Contact[0], err
}

func (a *App) CRMSaveContact(model crm.Contact) (*crm.Contact, error) {
	var contact crm.Contact

	payload, _ := json.Marshal(model)
	response, err := a.CRMHandlerPostService(RESOURCE, payload)

	if err == nil {
		err = json.Unmarshal(response, &contact)
	}

	return &contact, err
}

func (a *App) CRMUpdateContact(model crm.Contact) (*crm.Contact, error) {
	var contact crm.Contact

	payload, _ := json.Marshal(model)
	response, err := a.CRMHandlerPutService(RESOURCE+"/"+model.ID, payload)

	if err == nil {
		err = json.Unmarshal(response, &contact)
	}

	return &contact, err
}
