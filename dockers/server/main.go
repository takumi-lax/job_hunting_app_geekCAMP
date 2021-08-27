package main // import "server"

import (
    "github.com/gin-gonic/gin"
    "server/mypkg/textserch"
    "server/mypkg/nearbyserch"
    "fmt"
    // "encoding/json"
    // "net/http"
	"log"
    "bufio"
	"os"
	// "io/ioutil"
)

func main() {
    
    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "ping",
        })
    })
    r.POST("/",func(c *gin.Context) {
        // var NAG NameAndGeo
        var N InputCompany
        // fmt.Println(c)
        c.BindJSON(&N)
        // fmt.Println(N.Name)
        data := textserch.Textserch(N.Name,tokenread())
        data2 := nearbyserch.Nearbyserch(data.Lat,data.Lng,"1500","restaurant",tokenread())
        // textserch.Textserch(N)
        c.JSON(200, gin.H{
            "Company": data,
            "restaurant": data2,
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
