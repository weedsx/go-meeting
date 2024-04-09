package helper

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/google/uuid"
	"go-meeting/internal/define"
)

type UserClaims struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	jwt.StandardClaims
}

// GenerateToken 生成 token
func GenerateToken(id uint, name string) (string, error) {
	userClaims := &UserClaims{
		Id:             id,
		Name:           name,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, userClaims)
	tokenStr, err := token.SignedString(define.MyKey)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

// ParseToken 解析token
func ParseToken(tokenStr string) (*UserClaims, error) {
	userClaims := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenStr, userClaims, func(token *jwt.Token) (interface{}, error) {
		return define.MyKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, fmt.Errorf("analyse Token Error:%v", err)
	}
	return userClaims, nil
}

func GetMd5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func GetUUID() string {
	return uuid.New().String()
}

func Encode(obj any) string {
	b, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(b)
}

func Decode(in string, objPointer any) {
	b, err := base64.StdEncoding.DecodeString(in)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, objPointer)
	if err != nil {
		panic(err)
	}
}
