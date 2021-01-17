// 用户认证相关服务
package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/fenixsoft/monolithic_arch_golang/domain"
	"github.com/fenixsoft/monolithic_arch_golang/infrasturcture/db"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

const signingPrivateKey = "601304E0-8AD4-40B0-BD51-0B432DC47461"

type JWT struct {
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
	Authorities  []string `json:"authorities"`
	Expires      int64    `json:"expires_in"`
	Scope        string   `json:"scope"`
	TokenType    string   `json:"token_type"`
	Username     string   `json:"username"`
}

// 判断用户名、密码是否对应一个有效的用户
func CheckUserAccount(context context.Context, username, password string) (*domain.Account, error) {
	account, err := domain.NewAccount(db.NewDB()).GetByName(username)
	if err != nil {
		return nil, errors.New("用户不存在")
	} else if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password)); err != nil {
		return nil, errors.New("密码不正确")
	}
	return account, nil
}

// 根据指定用户信息输出JWT令牌
func BuildJWTAccessToken(account *domain.Account) *JWT {
	tokenId := uuid.New()
	refreshTokenId := uuid.New()
	r := &JWT{
		Authorities: []string{"ROLE_USER", "ROLE_ADMIN"},
		Expires:     60 * 60 * 3,
		Scope:       "BROWSER",
		TokenType:   "bearer",
		Username:    account.Username,
	}

	// 生成访问令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"jti":         tokenId,
		"scope":       [...]string{r.Scope},
		"authorities": r.Authorities,
		"exp":         time.Now().Unix() + r.Expires,
		"client_id":   "bookstore_frontend",
		"user_name":   r.Username,
		"username":    r.Username,
	})
	tokenString, _ := token.SignedString([]byte(signingPrivateKey))
	r.AccessToken = tokenString

	// 生成刷新令牌
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"jti":         refreshTokenId,
		"ati":         tokenId,
		"scope":       [...]string{r.Scope},
		"authorities": r.Authorities,
		"exp":         time.Now().Unix() + 60*60*24*15,
		"client_id":   "bookstore_frontend",
		"user_name":   r.Username,
		"username":    r.Username,
	})
	tokenString, _ = token.SignedString([]byte(signingPrivateKey))
	r.RefreshToken = tokenString

	return r
}

// 解析JWT令牌字符串，验证有效性，如令牌有效，提取其中存放的信息以jwt.MapClaims的形式返回
func ValidatingJWTAccessToken(tokenString string) (jwt.MapClaims, error) {
	if strings.HasPrefix(tokenString, "bearer ") {
		tokenString = tokenString[7:]
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v\n", token.Header["alg"])
		}
		return []byte(signingPrivateKey), nil
	})

	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, jwt.NewValidationError("Not a validation JWT token", jwt.ValidationErrorMalformed)
	}
}

// 将明文密码用Bcrypt加密为密文
func PasswordEncoding(rawPassword string) string {
	encode, _ := bcrypt.GenerateFromPassword([]byte(rawPassword), -1)
	return string(encode)
}
