package login

type RegisterRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
    Role string `json:"role"`
    Id string `json:"id"`
}