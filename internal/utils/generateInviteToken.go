package utils

import (
	"github.com/example/internal/entity"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateInviteToken(invite entity.Invite) (string, error) {
	// generate a token for the invite using
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"invite_id":  invite.ID,
		"email":      invite.Email,
		"role_id":    invite.RoleId,
		"expires_at": invite.ExpiresAt.Unix(),
	}).SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", err
	}
	return token, nil
}
