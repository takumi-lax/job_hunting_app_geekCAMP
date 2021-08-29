package main // import "server"

import (
    "github.com/gin-gonic/gin"
    "server/mypkg/textserch"
    "server/mypkg/nearbyserch"
    "github.com/gin-contrib/cors"
    "fmt"
    // "encoding/json"
    // "net/http"
	"log"
    "bufio"
	"os"
    "time"
	// "io/ioutil"
)

func main() {
    
    r := gin.Default()

      // ここからCorsの設定
    
    r.Use(cors.New(cors.Config{
    // アクセスを許可したいアクセス元
    AllowOrigins: []string{
        "https://humansyukatu.herokuapp.com/",
        "http://localhost:8080",
       
    },
    // アクセスを許可したいHTTPメソッド(以下の例だとPUTやDELETEはアクセスできません)
    AllowMethods: []string{
        "POST",
        "GET",
        "OPTIONS",
    },
    // 許可したいHTTPリクエストヘッダ
    AllowHeaders: []string{
        "Access-Control-Allow-Credentials",
        "Access-Control-Allow-Headers",
        "Content-Type",
        "Content-Length",
        "Accept-Encoding",
        "Authorization",
    },
    // cookieなどの情報を必要とするかどうか
    AllowCredentials: true,
    // preflightリクエストの結果をキャッシュする時間
    MaxAge: 24 * time.Hour,
}))

    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "ping",
        })
    })
    r.POST("/",func(c *gin.Context) {
        // var NAG NameAndGeo
        var N InputCompany
        fmt.Println(c)
        c.BindJSON(&N)
        fmt.Println(N.Name)
        data := textserch.Textserch(N.Name,tokenread())
        data2 := nearbyserch.Nearbyserch(data.Lat,data.Lng,"1500","restaurant",tokenread())
        // textserch.Textserch(N)
        c.JSON(200, gin.H{
            "Name": data2.Name,
            "Lat": data2.Lat,
            "Lng": data2.Lng,

        })
    })
    r.Run(":8888")
}

// func postfunc(c *gin.Context) {
// 	var NAG NameAndGeo
// 	c.BindJSON(&NAG)
// 	fmt.Println(NAG.Name)
//     textserch.Textserch(NAG.Name)
// 	c.JSON(200, gin.H{
// 		"message": "textserch.Textserch(NAG.Name)",
// 	})
// }
// type NameAndGeo struct{
//     Name string

// }

type InputCompany struct {
	Name string `json:"Name"`
    // Locate string `json:"Locate"`
}

func tokenread()(string) {
	// open the file
	file, err := os.Open("token.txt")

	//handle errors while opening
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	fileScanner := bufio.NewScanner(file)


	var token string
	// read line by line
	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
		token = fileScanner.Text()
	}
	// handle first encountered error while reading
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	file.Close()
	return token
}
