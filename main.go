package request

import (
	"net/http"
	"fmt"
	"strings"
	"log"
	"encoding/json"
)

func Get(url string, params map[string]interface{}) (interface{}, error) {
	url += makeQueryString(params)
	errors := make([]string, 0)
	res, err := http.Get(url)
	if err != nil {
		errors = append(errors, err.Error())
	}
	var data map[string]interface{}
	err2 := json.NewDecoder(res.Body).Decode(data)
	if err2 != nil {
		errors = append(errors, err2.Error())
	}
	if len(errors) > 0 {
		log.Fatalf("Error GETing url %s:\n\tStatus: %v\n\tError: %v", url, res.Status, strings.Join(errors, ", "))
	}
	return data, err
}

func Post(url string, params map[string]interface{}) {

}

func makeQueryString(params map[string]interface{}) string {
	queryString := "?"
	querySlice := make([]string, 0)
	for key, value := range params {
		querySlice = append(querySlice, fmt.Sprintf("%v=%v", key, value))
	}
	queryString += strings.Join(querySlice, "&")
	return queryString
}
