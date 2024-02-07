package user

type RequestUserParam struct {
	ID               string `json:"id,omitempty"`
	Email            string `json:"email,omitempty"`
	Password         string `json:"password,omitempty"`
	User_ID          string `json:"user_id,omitempty"`
	FirstName        string `json:"first_name,omitempty"`
	LastName         string `json:"last_name,omitempty"`
	Gender           string `json:"gender,omitempty"`
	BirthDay         string `json:"birth_day,omitempty"`
	PhoneNumber      string `json:"phone_number,omitempty"`
	PostOfficeNumber string `json:"post_office_number,omitempty"`
	Pref             string `json:"pref,omitempty"`
	City             string `json:"city,omitempty"`
	Extra            string `json:"extra,omitempty"`
}
