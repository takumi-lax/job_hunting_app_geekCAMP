package textserch

import (
	"context"
	// "fmt"

	// "encoding/json"
	// nearbyserch "getmapinfo/mypkg/nearbyserch"
	"log"
	// "strconv"


	// "github.com/kr/pretty"
	"googlemaps.github.io/maps"
)

func Textserch(name string,KEY string) (data) {
	c, err := maps.NewClient(maps.WithAPIKey(KEY))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	r := &maps.TextSearchRequest{
		Query: name,
		// Type: "university",
	}

	res, err := c.TextSearch(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	// pretty.Println(res)
	var gmapinfo data
	// fmt.Println(res.Results[0].Name)
	// for _, result := range res.Results {
	// 	// fmt.Println(result.Geometry.Location)
	// 	// fmt.Println(resu)
	// }
	gmapinfo.Name = res.Results[0].Name
	gmapinfo.Lat = res.Results[0].Geometry.Location.Lat
	gmapinfo.Lng = res.Results[0].Geometry.Location.Lng

	return gmapinfo
}

type data struct {
    Name string `json: "lat"`
	Lat   float64    `json: "lat"`
    Lng float64 `json: "lng"`
}
