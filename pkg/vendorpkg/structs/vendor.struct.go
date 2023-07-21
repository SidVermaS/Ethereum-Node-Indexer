package vstructs

type Vendor struct {
	BaseURL      string
	Username     string
	Password     string
	Token        string
	CustomConfig map[string]interface{}
}
