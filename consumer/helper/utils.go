package helper

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

// JSONResponse is a helper function to send JSON response with status code
func JSONResponse(c *gin.Context, statusCode int, msg interface{}) {
	c.JSON(statusCode, gin.H{"message": msg})
}

// ToJSON converts a struct to JSON string
func ToJSON(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return "{}"
	}
	return string(b)
}
