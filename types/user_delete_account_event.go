package types

type UserDeleteAccountEvent struct {
	UserId    string `json:"userId"`
	UserType  string `json:"userType"`
	UserEmail string `json:"userEmail"`
}
