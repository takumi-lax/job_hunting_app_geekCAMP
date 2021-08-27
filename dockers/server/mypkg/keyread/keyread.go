package keyread

import (
	"bufio"
	"fmt"

	// "encoding/json"

	// nearbyserch "getmapinfo/mypkg/nearbyserch"
	"log"

	"os"
	// "github.com/kr/pretty"
)

func keyread() string {
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
