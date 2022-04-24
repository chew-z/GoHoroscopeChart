package main

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/mshafiee/swephgo"
)

var (
	loc         string
	city        = os.Getenv("CITY")
	houseSystem = os.Getenv("HOUSE_SYSTEM")
	location    *time.Location
	swisspath   = os.Getenv("SWISSPATH")
)

func init() {
	location, _ = time.LoadLocation(city)
	swephgo.SetEphePath([]byte(swisspath))
}

func main() {
	defer swephgo.Close()
	//
	router := gin.Default()
	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")
	router.StaticFile("/favicon.ico", "./static/favicon.ico") // some clients don't read webmanifest
	router.StaticFile("/astrochart.min.js", "./static/astrochart.min.js")
	router.GET("/", PrintHoroscope)
	router.GET("/animate", printTransit)
	router.GET("/radix", printRadix)
	router.GET("/transit", printTransit)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run()
}
