package main
import(
	"github.com/gin-gonic/gin"
	"net/http"
)
func main() {
	r := gin.Default()
	r.GET("/hello", func (c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
		"method" : "GET",
	    })
    })
	r.PUT("/hello", func (c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"method" : "PUT",
		})
	})
	r.POST("/hello", func (c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"method" : "POST",
		})
	})
	r.DELETE("/hello", func (c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"method" : "DELETE",
		})
	})

	r.Run(":9090")
}