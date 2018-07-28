package controllers

import "github.com/gin-gonic/gin"

// ContextData data
func ContextData(c *gin.Context) (map[string]interface{}, bool) {
	ctx, exists := c.Get("context")
	context := ctx.(map[string]interface{})
	return context, exists
}
