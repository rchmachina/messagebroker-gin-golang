package handler

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"sync"
	helper "github.com/rchmachina/playingBrokerMessage/consumer/helper"
	"github.com/gin-gonic/gin"
)



type PersonDataNow struct {
	Person []Person
}

func getPersonNow() PersonDataNow {
	return PersonDataNow{Person: ReceivedMessages}
}

func GetDataCsv(c *gin.Context) {
	data := getPersonNow().Person
	helper.JSONResponse(c,200,data)
}


func DownloadCSV(c *gin.Context) {
	var mtx sync.Mutex
	mtx.Lock()
	defer mtx.Unlock()

	// Create CSV file
	file, err := os.Create("output.csv")
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to create CSV file")
		return
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	header := getStructFieldNames(Person{})
	if err := writer.Write(header); err != nil {
		c.String(http.StatusInternalServerError, "Failed to write CSV header")
		return
	}

	// Write CSV data
	data := getPersonNow().Person
	log.Println(data)
	for _, person := range data {
		fields := getStructFieldValues(person)
		if err := writer.Write(fields); err != nil {
			c.String(http.StatusInternalServerError, "Failed to write CSV data")
			return
		}
	}

	// Set response headers
	c.Header("Content-Type", "text/csv")
	c.Header("Content-Disposition", "attachment; filename=data.csv")

	// Serve the file
	c.File("output.csv")
}


func getStructFieldNames(s interface{}) []string {
	val := reflect.ValueOf(s)
	typ := val.Type()

	var fields []string
	for i := 0; i < typ.NumField(); i++ {
		fields = append(fields, typ.Field(i).Name)
	}
	return fields
}

func getStructFieldValues(s interface{}) []string {
	val := reflect.ValueOf(s)

	var fields []string
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		switch field.Kind() {
		case reflect.String:
			fields = append(fields, field.String())
		case reflect.Int:
			fields = append(fields, strconv.FormatInt(field.Int(), 10))
		// Add cases for other types if needed
		}
	}
	return fields
}
