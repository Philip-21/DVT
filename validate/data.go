package validate


type RequestData struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	DateTime string `json:"date"`
}