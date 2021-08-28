package nearbyserch

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
)

func Nearbyserch(Lat float64,Lng float64,radius string,typ string,key string) (datas) {

	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=%f,%f&radius=%s&type=%s&key=%s",Lat,Lng,radius,typ,key)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		// return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		// return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		// return
	}
	var datas  datas
	s := string(body)
    var jsondata map[string]interface{}
    json.Unmarshal([]byte(s), &jsondata)
    results := jsondata["results"].([]interface{})
	for _,result := range results{
		Name := result.(map[string]interface{})["name"]
		geometry := result.(map[string]interface{})["geometry"].(map[string]interface{})
		location := geometry["location"].(map[string]interface{})
		lat := location["lat"]
		lng := location["lng"]
		// fmt.Println(Name,lat, lng)
		var inf data
		inf.Name = Name.(string)
		inf.Lat = lat.(float64)
		inf.Lng = lng.(float64)
		datas = append(datas, inf)
	}

	return datas
}

type data struct {
    Name string `json: "lat"`
	Lat   float64    `json: "lat"`
    Lng float64 `json: "lng"`
}

type datas []data