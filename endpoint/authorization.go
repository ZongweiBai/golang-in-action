package endpoint

import (
	"errors"
	"github.com/ZongweiBai/learning-go/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

// GenerateAcessToken 生成JWTtoken
// @Summary 生成JWTtoken
// @Description 通过用户信息生成JWTtoken
// @Tags Token相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Basic 用户令牌"
// @Success 200 {object} JwtTokenMessage
// @Router /v1/oauth/token [post]
func GenerateAcessToken(c *gin.Context) {
	config.LOG.Debugf("进入到GenerateAcessToken方法")

	jwtTokenMessage := JwtTokenMessage{}
	err := c.BindJSON(&jwtTokenMessage)
	if err != nil {
		c.JSON(400, gin.H{"error_code": "invalid_request", "description": "Post Data Err"})
	} else {
		accessToken, err := generateToken(jwtTokenMessage.UserName, jwtTokenMessage.UserId)
		if err != nil {
			config.LOG.Errorf("生成Token失败:%s", err.Error())
			c.JSON(500, gin.H{"error_code": "server_error", "description": "Generate AccessToken Err"})
			return
		}

		jwtTokenMessage.AccessToken = accessToken
		c.JSON(200, &jwtTokenMessage)
	}
}

// ValidateAcessToken 校验JWTtoken
// @Summary 校验JWTtoken
// @Description 校验JWTtoken
// @Tags Token相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Basic 用户令牌"
// @Success 200 {object} JwtTokenMessage
// @Router /v1/oauth/token/validate [get]
func ValidateAcessToken(c *gin.Context) {
	config.LOG.Debugf("进入到ValidateAcessToken方法")

	accessToken := c.Query("accessToken")
	if accessToken == "" {
		c.JSON(400, gin.H{"error_code": "invalid_request", "description": "Query Data accessToken is Empty"})
	} else {
		claims, err := validateToken(accessToken)
		if err != nil {
			config.LOG.Errorf("校验Token失败:%s", err.Error())
			c.JSON(500, gin.H{"error_code": "server_error", "description": "Validate AccessToken Err"})
			return
		}

		c.JSON(200, claims)
	}
}

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	Username string `json:"userName"`
	UserId   string `json:"userId"`
	jwt.StandardClaims
}

type JwtTokenMessage struct {
	UserName    string `json:"userName"`
	UserId      string `json:"userId"`
	AccessToken string `json:"accessToken"`
}

// 生成Token
func generateToken(username string, userId string) (string, error) {
	c := MyClaims{
		username,
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(config.CONFIG.Jwt.Expires) * time.Second).Unix(), // 过期时间
			Issuer:    "learning-go",
		},
	}

	config.LOG.Infof("用户名：%s, ID: %s", username, userId)
	// 创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	config.LOG.Infof("Token: %s", token)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	accessToken, err := token.SignedString([]byte(config.CONFIG.Jwt.JwtSecret))
	return accessToken, err
}

// 解析|校验Token
func validateToken(tokenString string) (*MyClaims, error) {
	// 解析Token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(config.CONFIG.Jwt.JwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
