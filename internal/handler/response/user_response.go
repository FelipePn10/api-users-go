package response

type UserResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Created string `json:"created_at"`
	Updated string `json:"updated_at"`
}
