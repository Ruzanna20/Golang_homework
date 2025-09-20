package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Hotel struct {
	Data []struct {
		SearchType string `json:"search_type"`
		Name       string `json:"name"`
		CityName   string `json:"city_name"`
		DestID     string  `json:"dest_id"`
		ImageURL   string `json:"image_url"`
	} `json:"data"`
}

func main() {
	url := "https://booking-com15.p.rapidapi.com/api/v1/hotels/searchDestination?query=Yerevan"

	req, err := http.NewRequest("GET", url, nil)
	CheckErr(err)

	req.Header.Add("X-RapidAPI-Key", "e57abde3a2msh9c348a96d8a2b21p162704jsnf3640ed052b4")
	req.Header.Add("X-RapidAPI-Host", "booking-com15.p.rapidapi.com")

	client := http.Client{}
	resp, err := client.Do(req)
	CheckErr(err)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	CheckErr(err)

	var hotels Hotel
	err = json.Unmarshal(body, &hotels)
	CheckErr(err)


	connStr := "host=localhost port=5432 user=postgres password=12345 dbname=Travel_db sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	CheckErr(err)

	err = db.Ping()
	if err != nil {
		fmt.Println("Ping error:", err)
		return
	}


	for _, hot := range hotels.Data {
		if hot.SearchType == "hotel" {
			_, err = db.Exec(
				`INSERT INTO hotels (hotelName,hotelAddress,city, price, rating, description, hotelCreatedAt)
				VALUES ($1,NULL, $2, NULL, NULL, '', CURRENT_TIMESTAMP)`,
				hot.Name, hot.CityName, )
			if err != nil {
				log.Println("Error inserting hotel:", hot.Name, err)
			} else {
				fmt.Println("Inserted hotel:", hot.Name)
			}
		}
	}

	fmt.Println("All hotels are in Travel_db")
}
