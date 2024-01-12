package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float64 `json:"rates"`
}

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func getRates(c *gin.Context) {
	t := time.Now()

	// Generate a random sleep duration between 0 and 6 seconds
	minSleep := 0
	maxSleep := 6
	randomSleep := minSleep + rng.Intn(maxSleep-minSleep+1)

	data := Response{
		Base: "EUR",
		Date: t.Format("2006-01-02"),
		Rates: map[string]float64{
			"USD": 1.0741,
			"JPY": 160.87,
			"BGN": 1.9558,
			"CZK": 24.443,
			"DKK": 7.4587,
			"GBP": 0.86640,
			"HUF": 379.98,
			"PLN": 4.4670,
			"RON": 4.9692,
			"SEK": 11.6575,
			"CHF": 0.9646,
			"ISK": 150.50,
			"NOK": 11.8235,
			"TRY": 30.5458,
			"AUD": 1.6524,
			"BRL": 5.2595,
			"CAD": 1.4664,
			"CNY": 7.8094,
			"HKD": 8.4020,
			"IDR": 16684.96,
			"ILS": 4.1723,
			"INR": 89.3960,
			"KRW": 1394.34,
			"MXN": 18.7638,
			"MYR": 4.9801,
			"NZD": 1.7987,
			"PHP": 60.049,
			"SGD": 1.4501,
			"THB": 38.200,
			"ZAR": 19.6149,
		},
	}

	time.Sleep(time.Duration(randomSleep) * time.Second)

	if randomSleep > 3 {
		c.String(http.StatusServiceUnavailable, "")
	} else {
		c.IndentedJSON(http.StatusOK, data)
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.GET("/rates", getRates)

	router.Run("localhost:8082")
}
