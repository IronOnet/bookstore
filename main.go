package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/irononet/bookstore/utils"
)


type Bookstore struct{
	Books []*utils.Book 
}

func init(){
	// check if the books.json file has some content in it 
	// if it has not, generate the local database
	utils.GenerateBooks()
}


func main(){

	router := gin.Default() 

	router.GET("/books", getBookHandler) 

	log.Fatal(router.Run(":8080"))

}

func getBookHandler(c *gin.Context){
	// Read the JSON file containing book data 
	data, err := os.ReadFile("books.json") 
	if err != nil{
		log.Println("failed to read json file:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read book data"})
		return 
	}

	// Unmarshall the JSON data into a slice of Book structs 
	var books []utils.Book 
	err = json.Unmarshal(data, &books)
	if err != nil{
		log.Println("Failed to unmarshal JSON:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse book data"})
		return 
	}

	c.JSON(http.StatusOK, books)
}