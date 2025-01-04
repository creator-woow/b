package pass

import "golang.org/x/crypto/bcrypt"

func GenerateHash(p string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
	return string(bytes), err
}

func ComparePasswordAndHash(p, h string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(h), []byte(p))
	return err == nil
}
