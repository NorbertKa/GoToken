package jwt

import (
	"github.com/NorbertKa/GoToken/models"
)


type Token struct {
	Claims *Claims `json:"claims"`
	RefreshToken string `json:"refreshToken"`
	refreshTokenInitialized bool `json:"-"`
	User *models.Userprofile `json:"user"`
}