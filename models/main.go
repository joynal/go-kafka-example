package models

// Subscription property
type Property struct {
	FirstName string `json:"$first_name" url:"$first_name"`
	LastName  string `json:"$last_name" url:"$last_name"`
}

// Subscription model
type Subscription struct {
	APIKey       string   `json:"api_key" url:"api_key"`
	Email        string   `json:"email" url:"email"`
	Properties   Property `json:"properties" url:"properties"`
	ConfirmOptin bool     `json:"confirm_optin" url:"confirm_optin"`
}

// User data model
type User struct {
	Email     string `json:"email" url:"email"`
	FirstName string `json:"firstName" url:"firstName"`
	LastName  string `json:"lastName" url:"lastName"`
}
