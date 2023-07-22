package vstructs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	vconsts "github.com/SidVermaS/Ethereum-Consensus/pkg/vendorpkg/consts"
	"golang.org/x/exp/slices"
)

type Vendor struct {
	BaseURL      string
	Username     string
	Password     string
	Token        string
	CustomConfig map[string]interface{}
}
type APIRequest struct {
	Body                     interface{}
	Method                   vconsts.HttpMethodsE
	Query                    url.Values
	Param                    string
	Url                      string
}

func (vendor *Vendor) CallAPI(apiRequest *APIRequest) (int, []byte, error) {
	// Converting map[string]string to a query parameter format
	query := apiRequest.Query.Encode()
	if len(apiRequest.Param) > 0 {
		// Fetch an individual parameter (for eg. - id) and add a slash
		apiRequest.Param = "/" + apiRequest.Param
	}
	if len(query) > 0 {
		// Add a ? before the query parameters
		query = "?" + query
	}
	url := fmt.Sprintf("%s%s%s%s", vendor.BaseURL, apiRequest.Url, apiRequest.Param, query)
	client := &http.Client{}
	var request *http.Request
	var err error
	jsonPayload, err := json.Marshal(apiRequest.Body)

	if slices.Contains([]vconsts.HttpMethodsE{vconsts.GET, vconsts.DELETE}, apiRequest.Method) {
		request, err = http.NewRequest(string(apiRequest.Method), url, nil)
	} else {
		request, err = http.NewRequest(string(apiRequest.Method), url, bytes.NewBuffer(jsonPayload))
	}
	if err != nil {
		panic(err.Error())
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)
	bodyBytes, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	// if err != nil {
	// 	return response.StatusCode, bodyBytes
	// }

	return response.StatusCode, bodyBytes, err
}
