package internal

import (
	"BaiCloud/internal/svc"
	"context"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func JwtVerify(token string, svcCtx *svc.ServiceContext, ctx context.Context) string {
	name := ParseToken(token)
	if name == "" {
		return ""
	}
	userInfo, _ := svcCtx.UserModel.FindOne(ctx, name)
	if name == userInfo.UserName {
		return name
	} else {
		return ""
	}
}

type MyClaims struct {
	Id                   string `json:"Id"`
	jwt.RegisteredClaims        // 注意!这是jwt-go的v4版本新增的，原先是jwt.StandardClaims
}

var MySecret = []byte("dsvkb4d5fv46sv4sd86vsd4v") // 定义secret，后面会用到

func MakeToken(Id string) (tokenString string, err error) {
	claim := MyClaims{
		Id: Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * time.Hour * time.Duration(1))), // 过期时间3小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                       // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                                       // 生效时间
		}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim) // 使用HS256算法
	tokenString, err = token.SignedString(MySecret)
	return tokenString, err
}

func Secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte("dsvkb4d5fv46sv4sd86vsd4v"), nil // 这是我的secret
	}
}
func ParseToken(tokenss string) string {
	token, err := jwt.ParseWithClaims(tokenss, &MyClaims{}, Secret())
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors != 0 {
				return ""
			}
		}
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims.Id
	}
	return ""
}
