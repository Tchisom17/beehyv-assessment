package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sync"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func ExtractHandler(c *gin.Context) {
	data := c.Query("data")
	if data == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "data query parameter cannot be empty"})
		return
	}
	log.Printf("Extracting data: %s\n", data)

	var wg sync.WaitGroup
	resultChan := make(chan User)

	//for i := 0; i < 2; i++ {
	//	wg.Add(1)
	//	go ExtractData(data, i, &wg, resultChan)
	//}
	wg.Add(1)
	go ExtractData(data, &wg, resultChan)

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	go func() {
		for result := range resultChan {
			log.Printf("Extracted result: %v\n", result)
		}
	}()
	c.JSON(http.StatusOK, gin.H{"message": "Processing started"})
}
func ExtractData(data string, wg *sync.WaitGroup, resultChan chan<- User) {
	defer wg.Done()
	url := "https://jsonplaceholder.typicode.com/users/" + data
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching user data from external API: %v\n", err)
		return
	}

	defer resp.Body.Close()

	var apiResponse User
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		log.Printf("Error decoding JSON response from external API: %v\n", err)
		return
	}
	result := apiResponse
	resultChan <- result
}

func main() {
	r := gin.Default()
	r.GET("/process", ExtractHandler)
	r.Run(":8080")
}
