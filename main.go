package main
import ("github.com/gin-gonic/gin")

func main(){
    r := gin.New()
    r.GET("/hello", func(c *gin.Context) {
        c.JSON(200, map[string]string{"message": "Hello World"})
    })

    r.Run(":9000") // listen and serve on 0.0.0.0:9000
}
