package app

import (
	"encoding/json"

	"github.com/Nick212/go-crm-espo/models/crm"
)

const (
	RESOURCE string = RESOURCE
)

func (a *App) CRMGetUserAll() (*crm.ResponseUser, error) {
	var user crm.ResponseUser
	response, err := a.CRMHandlerGetService(RESOURCE, "")

	if err == nil {
		err = json.Unmarshal(response, &user)
	}

	return &user, err
}

func (a *App) CRMGetUserById(id string) (*crm.User, error) {
	var user crm.User
	response, err := a.CRMHandlerGetService(RESOURCE, id)

	if err == nil {
		err = json.Unmarshal(response, &user)
	}

	return &user, err
}
