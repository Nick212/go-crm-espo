package app

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// CRMHandlerGetService make request for CRM.
func (a *App) CRMHandlerGetService(resource string, params string) ([]byte, error) {
	url := a.Config.HOST_CRM + resource + "/" + params
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println("Error: ", err)
	}
	req.Header.Add("Authorization", "Basic "+a.Config.TOKEN_CRM)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	return body, err
}

func (a *App) CRMHandlerPutService(resource string, payload []byte) ([]byte, error) {
	url := a.Config.HOST_CRM + resource
	method := "PUT"

	fmt.Println(url)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, strings.NewReader(string(payload)))

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Authorization", "Basic "+a.Config.TOKEN_CRM)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	body = bytes.TrimPrefix(body, []byte("239, 187, 191, 26"))

	// Enable only for test
	fmt.Println("Status Code: " + strconv.Itoa(res.StatusCode))
	// fmt.Println(string(body))

	return body, err
}

func (a *App) CRMHandlerPostService(resource string, payload []byte) ([]byte, error) {
	url := a.Config.HOST_CRM + resource
	method := "POST"

	// fmt.Println(string(payload))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, strings.NewReader(string(payload)))

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Authorization", "Basic "+a.Config.TOKEN_CRM)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println("Status Code: " + strconv.Itoa(res.StatusCode))
	// fmt.Println(string(body))

	return body, err
}

// CRMHandlerDeleteService delete request in the CRM API
func (a *App) CRMHandlerDeleteService(ctx *Context, resource, params string) ([]byte, error) {

	url := a.Config.HOST_CRM + resource + params
	method := "DELETE"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		ctx.Logger.Println(err)
		return nil, err
	}
	req.Header.Add("Authorization", "Basic "+a.Config.TOKEN_CRM)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		ctx.Logger.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	return body, err
}
