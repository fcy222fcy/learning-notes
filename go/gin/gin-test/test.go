package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type User struct {
	Username string `binding:"required" json:"Username"`
	Password string `binding:"required" json:"Password"`
}

func main() {
	r := gin.Default()

	// 注册路由
	r.POST("/register", register)
	r.POST("/login", login)
	r.GET("/profile", profile)

	userGroup := r.Group("/users")
	userGroup.Use(authMiddleWare())
	{
		// 路由分组
		userGroup.GET("", userList)
		userGroup.GET("/:id", getUserList)
	}
	r.Run(":8080")
}

func register(c *gin.Context) {

	var u User
	err := c.BindJSON(&u)
	if err != nil {
		// 代表服务器内部错误
		c.JSON(500, gin.H{"err": err.Error(), "status": false, "cause": fmt.Sprint("服务器内部错误")})
		return
	}
	// 已经用了required 保证数据不会为空
	// 业务层判断用户名是否重复
	if !IsRepetition(u.Username) {
		c.JSON(400, gin.H{"status": false, "cause": fmt.Sprint("用户名重复,请更换一个")})
		return
	}
	c.JSON(200, gin.H{"status": true})
}
func login(c *gin.Context) {

	var u User
	err := c.BindJSON(&u)
	if err != nil {
		c.JSON(500, gin.H{"status": false, "err": err.Error(), "cause": fmt.Sprint("服务器内部错误")})
		return
	}

	// 检验用户密码
	// 登录成功返回token
}
func profile(c *gin.Context) {

}
func userList(c *gin.Context) {

}
func getUserList(c *gin.Context) {

}

// IsRepetition 模拟
func IsRepetition(s string) bool {
	return true
}

func authMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			c.JSON(401, gin.H{"error": "未携带token"})
			c.Abort()
			return
		}
		claims, err := ParseToken(tokenStr)
		if err != nil {
			c.JSON(401, gin.H{"error": "token无效"})
			c.Abort()
			return
		}
		c.Set("user_id", claims.UserID)

		c.Next()
	}
}

type Claims struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

func TokenGenerate(userID int64) (string, error) {
	jwtSecret := ""

	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			// 过期时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseToken(tokenStr string) (*Claims, error) {
	jwtSecret := ""
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("方法不对：%s", token.Method.Alg())
		}
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("无效token")
	}
	return claims, nil
}
