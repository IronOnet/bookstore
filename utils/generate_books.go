package utils 

import(
	"encoding/json" 
	"io/ioutil" 
	"log"
	"math/rand"
	"time"
)





func GenerateBooks(){
	var books []Book 

	// Generate 100 book instances 
	for i:= 0; i < 100; i++{
		book := Book{
			Title: generateRandomString(10), 
			Author: generateRandomString(8),
			YearPublished: generateRandomYear(),
		}

		books = append(books, book)
	}

	// Convert the book slice to JSON 
	data, err := json.MarshalIndent(books, "", " ") 
	if err != nil{
		log.Fatal("failed to marshal JSON:", err)
	}

	// Write JSON data to file 
	err = ioutil.WriteFile("books.json", data, 0644) 
	if err != nil{
		log.Fatal("failed to write json file:", err)
	}

	log.Println("JSON file generated successfully")
}

func generateRandomString(length int) string{
	rand.Seed(time.Now().UnixNano()) 
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" 
	result := make([]byte, length) 
	for i:= 0; i < length; i++{
		result[i] = chars[rand.Intn(len(chars))] 
	}
	return string(result)
}

func generateRandomYear() int{
	rand.Seed(time.Now().UnixNano()) 
	return rand.Intn(2023-1960+1) + 1960
}