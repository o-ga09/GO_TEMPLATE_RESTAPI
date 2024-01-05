package user

type ResponseUser struct {
	User ResponseUserModel `json:"user,omitempty"`
}

type ResponseUserModel struct {
	ID               string `json:"id,omitempty"`
	Email            string `json:"email,omitempty"`
	Password         string `json:"password,omitempty"`
	User_ID          string `json:"user___id,omitempty"`
	FirstName        string `json:"first_name,omitempty"`
	LastName         string `json:"last_name,omitempty"`
	Gender           string `json:"gender,omitempty"`
	BirthDay         string `json:"birth_day,omitempty"`
	PhoneNumber      string `json:"phone_number,omitempty"`
	PostOfficeNumber string `json:"post_office_number,omitempty"`
	Address          string `json:"address,omitempty"`
}
type Response struct {
	Status string `json:"status,omitempty"`
}
