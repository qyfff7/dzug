package jwt

import (
	userservice "dzug/app/services/user/cmd"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

const CtxUserIDKey = "userID"

// CustomSecret 用于加盐的字符串
var CustomSecret = []byte("抖音")

// CustomClaims 自定义声明类型 并内嵌jwt.RegisteredClaims
// jwt包自带的jwt.RegisteredClaims只包含了官方字段
// 假设我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type CustomClaims struct {
	// 可根据需要自行添加字段
	UserID             int64 `json:"user_id"`
	jwt.StandardClaims       // 内嵌标准的声明
}

// GenToken 生成JWT
func GenToken(userID int64) (aToken string, err error) {
	// 创建一个我们自己的声明
	claims := CustomClaims{
		userID, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(userservice.UserConf.JwtExpire) * time.Hour).Unix(), //过期时间
			Issuer:    "douyin",                                                                         // 签发人
		},
	}
	//Access Token 使用指定的签名方法创建签名对象
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(CustomSecret)
	return

}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*CustomClaims, error) {
	// 解析token
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return CustomSecret, nil
	})
	if err != nil {
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// RefreshToken 刷新AccessToken，这里返回值可以只有AccessToken，也可以两个都有，根据实际业务情况进行调整
func RefreshToken(aToken, rToken string) (newAToken string, err error) {
	// 1.refresh token无效直接返回
	if _, err = jwt.Parse(rToken, func(token *jwt.Token) (i interface{}, err error) {
		return CustomSecret, nil
	}); err != nil {
		return
	}
	// 2.从旧access token中解析出claims数据
	var claims CustomClaims
	_, err = jwt.ParseWithClaims(aToken, &claims, func(token *jwt.Token) (i interface{}, err error) {

		return CustomSecret, nil
	})
	v, _ := err.(*jwt.ValidationError)
	// 3.当access token是过期错误 并且 refresh token没有过期时就创建⼀一个新的access token
	if v.Errors == jwt.ValidationErrorExpired {
		return GenToken(claims.UserID)
	}
	return
}

// GetUserID 获取当前登录的用户ID
func GetUserID(ctx *gin.Context) (userID int64, err error) {
	uid, ok := ctx.Get(CtxUserIDKey)
	if !ok {
		err = errors.New("获取用户 ID 出错")
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = errors.New("获取用户 ID 出错")
		return
	}
	return
}
