package middleware

import (
	"github.com/gin-gonic/gin"
)

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		// token := c.GetHeader("token")
		// if token == "" {
		// 	c.Abort()
		// 	c.JSON(http.StatusUnauthorized, response.NewBaseResponse(response.ResponseCodeUnauthorized))
		// 	return
		// }
		// userID := c.GetHeader("user_id")
		// if userID == "" {
		// 	c.Abort()
		// 	c.JSON(http.StatusBadRequest, response.NewBaseResponse(response.ResponseCodeBadRequest))
		// 	return
		// }

		// localUser, err := redis.Store.GetString("token:" + token)
		// if err != nil || localUser != userID {
		// 	c.Abort()
		// 	c.JSON(http.StatusForbidden, response.NewBaseResponse(response.ResponseCodeForbidden))
		// 	return
		// }
		c.Next()
	}
}
