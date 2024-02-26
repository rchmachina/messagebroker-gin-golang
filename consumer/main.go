// package main

// import (
// 	//"github.com/gin-gonic/gin"
// 	handler "github.com/rchmachina/playingBrokerMessage/consumer/handler"
// 	//"net/http"
// )

// func main() {
// 	// Kafka broker address
// 	go handler.ConsumeMessage()
// 	select {}
// 	// router := gin.Default()

//     // // Define a route handler
//     // router.GET("/hello", func(c *gin.Context) {
//     //     c.JSON(http.StatusOK, gin.H{
//     //         "message": "Hello, Gin!",
//     //     })
//     // })
//     // router.GET("/getCsv",handler.DownloadCSV)
//     // router.GET("/getData",handler.GetDataCsv)

//     // // Run the server
//     // router.Run(":8080")

// }
package main
import (
	"github.com/gin-gonic/gin"
	handler "github.com/rchmachina/playingBrokerMessage/consumer/handler"
	"net/http"
	"log"
)

func main() {
	// Start ConsumeMessage function in a goroutine\
	go func(){
		handler.ConsumeMessage()
	}()


	// Start HTTP server in a goroutine

		router := gin.Default()

		// Define a route handler
		router.GET("/hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Hello, Gin!",
			})
		})
		router.GET("/getCsv", handler.DownloadCSV)
		router.GET("/getData", handler.GetDataCsv)
		if err := router.Run(":8080"); err != nil {
			log.Fatalf("Failed to start HTTP server: %v", err)
		}
	


}