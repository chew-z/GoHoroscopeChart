package main

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/mshafiee/swephgo"
)

var (
	city        = os.Getenv("CITY")
	houseSystem = os.Getenv("HOUSE_SYSTEM")
	loc         string
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
	router.SetTrustedProxies(nil)
	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")
	router.StaticFile("/favicon.ico", "./static/favicon.ico") // some clients don't read webmanifest
	router.StaticFile("/astrochart.min.js", "./static/astrochart.min.js")
	router.GET("/", PrintHoroscope)
	router.GET("/radix", printRadix)
	router.GET("/transit", printTransit)
	router.Run(":8080")
}
