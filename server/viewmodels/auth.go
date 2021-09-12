package viewmodels

type AuthResponse struct {
	Token          string `json:"token"`
	ExpirationTime int64  `json:"expiration_time"`
}
