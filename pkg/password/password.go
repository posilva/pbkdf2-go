package password

import (
	"crypto/rand"
	"crypto/sha1"
	b64 "encoding/base64"
	"strings"

	uuid "github.com/google/uuid"
	"golang.org/x/crypto/pbkdf2"
)

// CheckPassword verifies if the password and a given hash match
func CheckPassword(password string, hash string) (bool, error) {

	decodedHash, err := b64.StdEncoding.DecodeString(hash)
	if err != nil {
		return false, err
	}
	salt := decodedHash[1:17]
	return HashPasswordWithSalt(password, salt) == hash, nil
}

// HashPasswordWithSalt generates an hashed password using a salt
func HashPasswordWithSalt(password string, salt []byte) string {
	dk := pbkdf2.Key([]byte(password), salt, 1000, 32, sha1.New)
	var hp []byte
	hp = append(hp, 0)
	hp = append(hp, salt...)
	hp = append(hp, dk...)
	return b64.StdEncoding.EncodeToString(hp)

}

// HashPassword generates an hashed password using a generated salt
func HashPassword(password string) string {
	salt := make([]byte, 16)
	rand.Read(salt)
	return HashPasswordWithSalt(password, salt)
}

func generateAccessCode() (string, error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return strings.ToUpper(uid.String()[:8]), nil
}
