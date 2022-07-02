package main

import (
    "net/http"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
    "github.com/gin-gonic/gin"
)

// album represents data about a record album.
type Estrategia struct {
    FromDate    		string  	`json:"FromDate" binding:"required`
    ToDate  			string  	`json:"ToDate" binding:"required`
    MaxContractAmount 	float64  	`json:"MaxContractAmount" binding:"required"`
}

type Contract struct {
    Id int `json:"id"`
	NewMaxContractAmount float64 `json:"NewMaxContractAmount"`
	ToDate string `json:"ToDate"`
}


func contracts(c *gin.Context) {

	var ap Estrategia

	if err := c.ShouldBindJSON(&ap); err != nil {
        c.JSON(400, gin.H{"type": "Parsing JSON failed",
						  "error": err.Error()})
        return
    }

	var contracts = []Contract{
		{Id: 1, NewMaxContractAmount: ap.MaxContractAmount * 1, ToDate: ap.ToDate},
		{Id: 2, NewMaxContractAmount: ap.MaxContractAmount * 2, ToDate: ap.ToDate},
		{Id: 3, NewMaxContractAmount: ap.MaxContractAmount * 3, ToDate: ap.ToDate},
	}

    c.IndentedJSON(http.StatusOK, contracts)

}


func get_tickets_from_csv_demo() {

	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	var filename = filepath.Join(path, "data", "tickers.csv")

    // open file
    f, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }

    // remember to close the file at the end of the program
    defer f.Close()

    // read csv values using csv.Reader
    csvReader := csv.NewReader(f)
    for {
        rec, err := csvReader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatal(err)
        }


        // do something with read line
        fmt.Printf("%+v\n", rec)
    }
}

func get_tickers() {

}

func get_share_price(string ticker, string FromDate, string ToDate) {

}


// album represents data about a record album.
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {

	get_tickets_from_csv_demo()

    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.GET("/albums/:id", getAlbumByID)
    router.POST("/albums", postAlbums)
    router.POST("/contracts", contracts)

    router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
    var newAlbum album

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    // Add the new album to the slice.
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

    // Loop through the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}