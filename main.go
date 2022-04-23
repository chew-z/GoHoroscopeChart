package main

import (
	"net/http"
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
	// router := gin.Default()
	router := gin.New() // gin.Default() installs gin.Recovery() so use gin.New() instead
	// A zero/default http.Server, like the one used by the package-level helpers
	// http.ListenAndServe and http.ListenAndServeTLS, comes with no timeouts.
	// You don't want that.
	server := &http.Server{
		Addr:              ":8080",
		Handler:           router,
		ReadHeaderTimeout: 3 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      25 * time.Second,
		IdleTimeout:       90 * time.Second,
	}
	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")
	router.StaticFile("/favicon.ico", "./static/favicon.ico") // some clients don't read webmanifest
	router.StaticFile("/astrochart.min.js", "./static/astrochart.min.js")
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/", PrintHoroscope)
	router.GET("/radix", printRadix)
	router.GET("/transit", printTransit)
	// router.Run()
	server.ListenAndServe()
}
