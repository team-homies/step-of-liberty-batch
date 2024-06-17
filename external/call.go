package external

import (
	"fmt"
	"net/http"
	"strings"
)

type CallApi interface {
	Get() (res *http.Response, err error)
}

type callApi struct {
	Host  string
	Path  string
	Query map[string]string
	Param map[string]string
	Body  any
}

func NewCall(host string, path string, query map[string]string, param map[string]string, body any) CallApi {
	return &callApi{
		Host:  host,
		Path:  path,
		Query: query,
		Param: param,
		Body:  body,
	}
}

func (c *callApi) Get() (res *http.Response, err error) {
	client := http.Client{}

	req, err := http.NewRequest("GET", c.Host+c.Path, nil)
	if err != nil {
		return
	}

	if c.Query != nil {
		query := ""
		for k, v := range c.Query {
			query += fmt.Sprintf("%s=%s", k, v)
			query += "&"
		}
		req.URL.RawQuery = query
	}

	if c.Param != nil {
		for k, v := range c.Param {
			req.URL.Path = strings.Replace(req.URL.Path, "{"+k+"}", v, -1)
		}
	}

	return client.Do(req)

	// http.Get(c.host + c.path)
}

func (c *callApi) Post() {

}

func (c *callApi) Put() {

}

func (c *callApi) Delete() {

}
