package responses

import "github.com/gin-gonic/gin"

// SendResourceNotFound sends an error message indicating that the specified
// resource was not found.
func SendResourceNotFound(c *gin.Context) {
	c.JSON(404, gin.H{
		"status":   "error",
		"messages": []string{"The specified resource was not found."},
	})
}

// SendSingleResource sends a single resource to the client
func SendSingleResource(c *gin.Context, name string, instance interface{}) {
	c.JSON(200, gin.H{
		"status": "success",
		name:     instance,
	})
}
