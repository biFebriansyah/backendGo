package libs

import "golang.org/x/crypto/bcrypt"

func HashPassword(pass string) (string, error) {
	hassPass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hassPass), nil
}

func CheckPassword(bodyPass, dbPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(bodyPass), []byte(dbPass))

	if err != nil {
		return false
	}

	return true
}
