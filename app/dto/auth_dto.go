package dto

type AuthRequestLogin struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

func AuthToResponse(token string) *AuthResponse {
	return &AuthResponse{
		Token: token,
	}
}
