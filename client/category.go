package client

import (
	"a21hc3NpZ25tZW50/config"
	"a21hc3NpZ25tZW50/model"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CategoryClient interface {
	CategoryList(token string) ([]*model.Category, error)
	AddCategory(token, name string) (respCode int, err error)
	UpdateCategory(token, id, name string) (respCode int, err error)
	DeleteCategory(token, id string) (respCode int, err error)
}

type categoryClient struct {
}

func NewCategoryClient() *categoryClient {
	return &categoryClient{}
}

func (c *categoryClient) CategoryList(token string) ([]*model.Category, error) {
	client, err := GetClientWithCookie(token)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", config.SetUrl("/api/v1/category/list"), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("status code not 200")
	}

	var Categories []*model.Category
	err = json.Unmarshal(b, &Categories)
	if err != nil {
		return nil, err
	}

	return Categories, nil
}

func (c *categoryClient) AddCategory(token, name string) (respCode int, err error) {
	client, err := GetClientWithCookie(token)
	if err != nil {
		return -1, err
	}

	datajson := map[string]string{
		"name": name,
	}

	data, err := json.Marshal(datajson)
	if err != nil {
		return -1, err
	}

	req, err := http.NewRequest("POST", config.SetUrl("/api/v1/category/add"), bytes.NewBuffer(data))
	if err != nil {
		return -1, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return -1, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		return -1, errors.New("status code not 201")
	}

	return resp.StatusCode, nil
}

func (c *categoryClient) UpdateCategory(token, id, name string) (respCode int, err error) {
	client, err := GetClientWithCookie(token)
	if err != nil {
		return -1, err
	}

	datajson := map[string]string{
		"name": name,
	}

	data, err := json.Marshal(datajson)
	if err != nil {
		return -1, err
	}

	req, err := http.NewRequest("PUT", config.SetUrl(fmt.Sprintf("/api/v1/category/update/%s", id)), bytes.NewBuffer(data))
	if err != nil {
		return -1, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return -1, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return -1, errors.New("status code not 200")
	}

	return resp.StatusCode, nil
}

func (c *categoryClient) DeleteCategory(token, id string) (respCode int, err error) {
	client, err := GetClientWithCookie(token)
	if err != nil {
		return -1, err
	}

	req, err := http.NewRequest("DELETE", config.SetUrl(fmt.Sprintf("/api/v1/category/delete/%s", id)), nil)
	if err != nil {
		return -1, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return -1, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return -1, errors.New("status code not 200")
	}

	return resp.StatusCode, nil
}
