package security

import "golang.org/x/crypto/bcrypt"

// HashPassword はパスワードをハッシュ化する
func HashPassword(password string) (string, error) {
    hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashed), nil
}
