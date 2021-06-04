package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type User struct {
	Type            string            `json:"type"`
	Name            string            `json:"name"`
	Email           string            `json:"email"`
	Id              string            `json:"id"`
	Role            string            `json:"role"`
	Contact_methods []Contact_methods `json:"contact_methods"`
}
type Whole_body struct {
	User User `json:"user"`
}
type Contact_methods struct {
	Type    string `json:"type"`
	Summary string `json:"summary"`
}

type Client struct {
	HTTPClient *http.Client
	Token      string
}

func NewClient(token string) *Client {
	return &Client{
		Token:      token,
		HTTPClient: &http.Client{Timeout: 20 * time.Second},
	}
}

func (c *Client) GetUser(id string) (*Whole_body, error) {
	url := "https://api.pagerduty.com/users/" + id
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/vnd.pagerduty+json;version=2")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", c.Token)
	res, _ := c.HTTPClient.Do(req)
	if res.StatusCode != 200 {
		log.Println("[ERROR]: ", res.Status, res.StatusCode)
		return nil, fmt.Errorf("unable to fetch the  user")
	}
	User_response := Whole_body{}
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &User_response)
	return &User_response, nil
}
func (c *Client) CreateUser(payload_body Whole_body) (*Whole_body, error) {
	url := "https://api.pagerduty.com/users/"
	payload, _ := json.Marshal(payload_body)
	req, _ := http.NewRequest("POST", url, strings.NewReader(string(payload)))
	req.Header.Add("accept", "application/vnd.pagerduty+json;version=2")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("from", "")
	req.Header.Add("authorization", c.Token)
	res, _ := c.HTTPClient.Do(req)
	fmt.Println(res.StatusCode)
	if res.StatusCode != 201 {
		log.Println("[ERROR]: ", res.Status, res.StatusCode)
		return nil, fmt.Errorf("unable to create the  user")
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	User_response := Whole_body{}
	json.Unmarshal(body, &User_response)
	return &User_response, nil
}
func (c *Client) UpdateUser(payload_body Whole_body, Id string) (*Whole_body, error) {
	url := "https://api.pagerduty.com/users/" + Id
	payload, _ := json.Marshal(payload_body)
	req, _ := http.NewRequest("PUT", url, strings.NewReader(string(payload)))
	req.Header.Add("accept", "application/vnd.pagerduty+json;version=2")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("from", "")
	req.Header.Add("authorization", c.Token)
	res, _ := c.HTTPClient.Do(req)
	if res.StatusCode != 200 {
		log.Println("[ERROR]: ", res.Status, res.StatusCode)
		return nil, fmt.Errorf("unable to update the user")
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	User_response := Whole_body{}
	json.Unmarshal(body, &User_response)
	return &User_response, nil
}
func (c *Client) DeleteUser(Id string) error {
	url := "https://api.pagerduty.com/users/" + Id
	req, _ := http.NewRequest("DELETE", url, nil)
	req.Header.Add("accept", "application/vnd.pagerduty+json;version=2")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("from", "")
	req.Header.Add("authorization", c.Token)
	res, _ := c.HTTPClient.Do(req)
	if res.StatusCode == 204 {
		return nil
	}
	defer res.Body.Close()
	log.Println("[ERROR]: ", res.Status, res.StatusCode)
	return fmt.Errorf("delete user failed")
}
