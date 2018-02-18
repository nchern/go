package wrappers

import "github.com/gin-gonic/gin"

// JSONHandler represents a JSON-response web handler
type JSONHandler func(c *gin.Context) (int, interface{})

// JSONHandler represents a String-response web handler
type StringHandler func(c *gin.Context) (int, string)

// JSON makes gin.HandlerFunc out of JSONHandler
func JSON(handler JSONHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		status, resp := handler(c)
		c.JSON(status, resp)
	}
}

// Text makes gin.HandlerFunc out of StringHandler
func Text(handler StringHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		status, resp := handler(c)
		c.String(status, resp)
	}
}
