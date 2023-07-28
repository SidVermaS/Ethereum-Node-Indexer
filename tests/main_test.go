package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/modules"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

type APIRequest struct {
	Body   interface{}
	Method string
	Query  *url.Values
	Param  string
	Url    string
}
type TestAPIRequest struct {
	T                    *testing.T
	ExpectedResponseCode int
	APIRequest           *APIRequest
}

// Declate the Fiber app.
var App *fiber.App

func Setup() {
	// Load the .env file
	godotenv.Load("../.env")
	// Connect to the DB, cache, start node scheduler and start event listenter
	modules.ActivateAll()
	// Create an instance of the fiber app
	App = fiber.New()
	// Enable cors
	App.Use(cors.New())
	// Setting up the API routes
	routes.SetupRoutes(App)
}

// Makes a network call to an API whose URL is passed.
func makeRequest(method, url string, body interface{}) (*http.Response, error) {
	// Parsing the payload to bytes
	requestBody, _ := json.Marshal(body)

	request, _ := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	// Perform the request plain with the app, the second argument is a request latency. (set to -1 for no latency)
	response, err := App.Test(request, -1)
	return response, err
}

func checkResponseCode(t *testing.T, expectedResponseCode, actualResponseCode int, url string) {
	// It Verifies, whether the status code is as expected or not
	assert.Equalf(t, expectedResponseCode, actualResponseCode, url)
}

// Common function to create URL and check the expected response code
func testAPI(testAPIRequest *TestAPIRequest) {
	// Converting map[string]string to a query parameter format
	query := testAPIRequest.APIRequest.Query.Encode()
	if len(testAPIRequest.APIRequest.Param) > 0 {
		// Fetch an individual parameter (for eg. - id) and add a slash
		testAPIRequest.APIRequest.Param = "/" + testAPIRequest.APIRequest.Param
	}
	if len(query) > 0 {
		// Add a ? before the query parameters
		query = "?" + query
	}
	// Creating the URL with parameters and a query
	u := fmt.Sprintf("%s%s%s", testAPIRequest.APIRequest.Url, testAPIRequest.APIRequest.Param, query)
	// Call the API on the server
	response, err := makeRequest(testAPIRequest.APIRequest.Method, u, testAPIRequest.APIRequest.Body)
	if err != nil {
		testAPIRequest.T.Errorf("~~~ Request Error %v", err)
	}
	checkResponseCode(testAPIRequest.T, testAPIRequest.ExpectedResponseCode, response.StatusCode, u)
}

func TestParticipationRate(t *testing.T) {
	// Connect to the DB, cache, start node scheduler and start event listenter
	Setup()

	// Multiple set of parameters to test the validator APIs
	var tests = []*TestAPIRequest{
		{
			APIRequest: &APIRequest{
				Url:    "/api/v1/participation/network",
				Method: "GET",
			},
			ExpectedResponseCode: http.StatusOK,
		},
		{
			APIRequest: &APIRequest{
				Url:    "/api/v1/participation/validators",
				Method: "GET",
				Param:  "1",
			},
			ExpectedResponseCode: http.StatusOK,
		},
	}

	for _, testItem := range tests {

		testAPI(&TestAPIRequest{
			T:                    t,
			APIRequest:           testItem.APIRequest,
			ExpectedResponseCode: testItem.ExpectedResponseCode,
		})
	}

}

func TestIndexers(t *testing.T) {
	// Connect to the DB, cache, start node scheduler and start event listenter
	Setup()
	// Define a structure for specifying input and output data of a single test case

	// Multiple set of parameters to test the validator APIs
	var tests = []*TestAPIRequest{
		{
			APIRequest: &APIRequest{
				Url:    "/api/v1/blocks",
				Method: "GET",
				Query:  &url.Values{},
			},
			ExpectedResponseCode: http.StatusOK,
		},
		{
			APIRequest: &APIRequest{
				Url:    "/api/v1/epochs",
				Method: "GET",
				Query:  &url.Values{},
			},
			ExpectedResponseCode: http.StatusOK,
		},
		{
			APIRequest: &APIRequest{
				Url:    "/api/v1/validators",
				Method: "GET",
				Query:  &url.Values{},
			},
			ExpectedResponseCode: http.StatusOK,
		},
		{
			APIRequest: &APIRequest{
				Url:    "/api/v1/slots",
				Method: "GET",
				Query:  &url.Values{},
			},
			ExpectedResponseCode: http.StatusOK,
		},
		{
			APIRequest: &APIRequest{
				Url:    "/api/v1/states",
				Method: "GET",
				Query:  &url.Values{},
			},
			ExpectedResponseCode: http.StatusOK,
		},
		{
			APIRequest: &APIRequest{
				Url:    "/api/v1/committees",
				Method: "GET",
				Query:  &url.Values{},
			},
			ExpectedResponseCode: http.StatusOK,
		},
		{
			APIRequest: &APIRequest{
				Url:    "/api/v1/validator-statuses",
				Method: "GET",
				Query:  &url.Values{},
			},
			ExpectedResponseCode: http.StatusOK,
		},
	}

	for _, testItem := range tests {
		testItem.APIRequest.Query.Add("page", "1")
		testItem.APIRequest.Query.Add("limit", "1")
		// testItem.APIRequest.Query = url.Values{
		// 	"page":  {"1"},
		// 	"limit": {"10"},
		// }

		testAPI(&TestAPIRequest{
			T:                    t,
			APIRequest:           testItem.APIRequest,
			ExpectedResponseCode: testItem.ExpectedResponseCode,
		})
	}
}

// func TestValidatorParticipationRate(t *testing.T) {
// 	// Connect to the DB, cache, start node scheduler and start event listenter
// 	Setup()
// 	// Define a structure for specifying input and output data of a single test case
// 	type FetchingValidatorParticipationRate struct {
// 		id                   int
// 		expectedResponseCode int
// 	}
// 	// Multiple set of parameters to test the validator APIs
// 	var tests = []FetchingValidatorParticipationRate{
// 		{
// 			id:                   1,
// 			expectedResponseCode: http.StatusOK,
// 		},
// 		{
// 			id:                   2,
// 			expectedResponseCode: http.StatusOK,
// 		},
// 		{
// 			id:                   -1,
// 			expectedResponseCode: http.StatusBadRequest,
// 		},
// 	}
// 	var apiRequest = &APIRequest{
// 		Url:    "/api/v1/indexers/validators",
// 		Method: "GET",
// 	}
// 	for _, testItem := range tests {
// 		apiRequest.Param = fmt.Sprint(testItem.id)

// 		testAPI(&TestAPIRequest{
// 			T:                    t,
// 			APIRequest:           apiRequest,
// 			ExpectedResponseCode: testItem.expectedResponseCode,
// 		})
// 	}

// }
