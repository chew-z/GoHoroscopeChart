package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Songmu/go-httpdate"
	"github.com/gin-gonic/gin"
	"github.com/go-echarts/go-echarts/v2/components"
	_ "github.com/joho/godotenv/autoload"
)

type Planet struct {
	Name     string
	Position float64
}

type HousePrint struct {
	Name     string
	Position string
	Cusp     string
	Sign     string
}

type PlanetPrint struct {
	Planet   string
	Position string
	House    string
	Sign     string
	Aspects  []string
}

func printRadix(c *gin.Context) {
	var Haus []int
	var Planeta []Planet
	var Asc, Mc float64
	when := time.Now()
	latitude := c.DefaultQuery("lat", os.Getenv("LATITUDE"))
	longitude := c.DefaultQuery("lon", os.Getenv("LONGITUDE"))
	lat, _ = strconv.ParseFloat(latitude, 64)
	lon, _ = strconv.ParseFloat(longitude, 64)
	if t := c.Query("t"); t != "" {
		if t1, err := httpdate.Str2Time(t, location); err == nil {
			when = t1
		}
	}
	if u := c.Query("u"); u != "" {
		if i, err := strconv.ParseInt(u, 10, 64); err == nil {
			when = time.Unix(i, 0)
		}
	}
	When := fmt.Sprintf("\n%s - lat: %.2f, lon: %.2f\n", when.In(location).Format(time.RFC822), lat, lon)
	if Cusps, Asmc, err := Cusps(when, lat, lon, houseSystem); err != nil {
		log.Println(err.Error())
		return
	} else {
		Asc = Asmc[0]
		Mc = Asmc[1]
		H := Houses(Cusps)
		for _, h := range *H {
			Haus = append(Haus, int(h.DegreeUt))
		}
		B := Bodies(when)
		for i, b1 := range B {
			var pl Planet
			pl.Name = getPlanetName(bodies[i])
			pl.Position = rad2deg(b1)
			Planeta = append(Planeta, pl)
		}
	}
	c.HTML(
		http.StatusOK,
		"radix.html",
		gin.H{
			"Planets":   Planeta,
			"Houses":    Haus,
			"Ascendant": Asc,
			"Mc":        Mc,
			"Now":       When,
			"title":     "radix",
		},
	)
}

func printTransit(c *gin.Context) {
	var Haus1, Haus2 []int
	var Planeta1, Planeta2 []Planet
	var AscNow, McNow float64
	latitude := c.DefaultQuery("lat", os.Getenv("LATITUDE"))
	longitude := c.DefaultQuery("lon", os.Getenv("LONGITUDE"))
	lat, _ = strconv.ParseFloat(latitude, 64)
	lon, _ = strconv.ParseFloat(longitude, 64)

	now := time.Now()
	Now := fmt.Sprintf("\n%s - lat: %.2f, lon: %.2f\n", now.In(location).Format(time.RFC822), lat, lon)
	when := time.Now().AddDate(0, 0, 7)
	if t := c.Query("t"); t != "" {
		if t1, err := httpdate.Str2Time(t, location); err == nil {
			when = t1
		}
	}
	if u := c.Query("u"); u != "" {
		if i, err := strconv.ParseInt(u, 10, 64); err == nil {
			when = time.Unix(i, 0)
		}
	}
	Then := fmt.Sprintf("\n%s - lat: %.2f, lon: %.2f\n", when.In(location).Format(time.RFC822), lat, lon)
	if Cusps, Asmc, err := Cusps(now, lat, lon, houseSystem); err != nil {
		log.Println(err.Error())
		return
	} else {
		AscNow = Asmc[0]
		McNow = Asmc[1]
		H := Houses(Cusps)
		for _, h := range *H {
			Haus1 = append(Haus1, int(h.DegreeUt))
		}
		B := Bodies(now)
		for i, b1 := range B {
			var pl Planet
			pl.Name = getPlanetName(bodies[i])
			pl.Position = rad2deg(b1)
			Planeta1 = append(Planeta1, pl)
		}
	}
	if Cusps, _, err := Cusps(when, lat, lon, houseSystem); err != nil {
		log.Println(err.Error())
		return
	} else {
		H := Houses(Cusps)
		for _, h := range *H {
			Haus2 = append(Haus2, int(h.DegreeUt))
		}
		B := Bodies(when)
		for i, b1 := range B {
			var pl Planet
			pl.Name = getPlanetName(bodies[i])
			pl.Position = rad2deg(b1)
			Planeta2 = append(Planeta2, pl)
		}
	}

	c.HTML(
		http.StatusOK,
		"animate.html",
		gin.H{
			"Now":         Now,
			"Then":        Then,
			"PlanetsNow":  Planeta1,
			"HousesNow":   Haus1,
			"PlanetsThen": Planeta2,
			"HousesThen":  Haus2,
			"Ascendant":   AscNow,
			"Mc":          McNow,
			"title":       "animate",
		},
	)
}

func printHoroscope(c *gin.Context) {
	var Haus []int
	var Planeta []Planet
	var Asc, Mc float64
	var T1 []HousePrint
	var T2 []PlanetPrint
	var Ascendant string
	when := time.Now()
	latitude := c.DefaultQuery("lat", os.Getenv("LATITUDE"))
	longitude := c.DefaultQuery("lon", os.Getenv("LONGITUDE"))
	lat, _ = strconv.ParseFloat(latitude, 64)
	lon, _ = strconv.ParseFloat(longitude, 64)
	if t := c.Query("t"); t != "" {
		if t1, err := httpdate.Str2Time(t, location); err == nil {
			when = t1
		}
	}
	if u := c.Query("u"); u != "" {
		if i, err := strconv.ParseInt(u, 10, 64); err == nil {
			when = time.Unix(i, 0)
		}
	}
	When := fmt.Sprintf("\n%s - lat: %.2f, lon: %.2f\n", when.In(location).Format(time.RFC822), lat, lon)

	if Cusps, Asmc, err := Cusps(when, lat, lon, houseSystem); err != nil {
		log.Println(err.Error())
		return
	} else {
		Asc = Asmc[0]
		Mc = Asmc[1]
		Ascendant = fmt.Sprintf("Ascendant: %.2f MC: %.2f, House system: %s\n", Asmc[0], Asmc[1], houseSystem)
		H := Houses(Cusps)
		// table1.AddHeaders("House", "Position", "Cusp", "Sign")
		for _, h := range *H {
			Haus = append(Haus, int(h.DegreeUt))
			var r1 HousePrint
			r1.Name = h.Number
			r1.Position = fmt.Sprintf("%.0f", h.DegreeUt)
			r1.Cusp = fmt.Sprintf("%.0f", h.Degree)
			r1.Sign = h.SignName
			T1 = append(T1, r1)
			// log.Printf("%s\t%.2f\t%.2f\t%s\n", h.Number, h.DegreeUt, h.Degree, h.SignName)
		}
		B := Bodies(when)
		// table2.AddHeaders("Planet", "Position", "House", "Sign", "Aspects")
		for i, b1 := range B {
			var r2 PlanetPrint
			var pl Planet

			pl.Name = getPlanetName(bodies[i])
			pl.Position = rad2deg(b1)
			Planeta = append(Planeta, pl)

			r2.Planet = getPlanetName(bodies[i])
			r2.House = getHouse(b1, H)
			r2.Position = fmt.Sprintf("%.0f", rad2deg(b1))
			r2.Sign = getSign(b1)
			// log.Printf("House %s: %s - %.2f in %s\n", getHouse(b1, H), getPlanetName(bodies[i]), rad2deg(b1), getSign(b1))
			for j, b2 := range B[i+1:] {
				if asp := Aspect(b1, b2); asp != "" {
					r3 := fmt.Sprintf("%s %s in %s", asp, getPlanetName(bodies[i+j+1]), getSign(b2))
					r2.Aspects = append(r2.Aspects, r3)
				}
			}
			T2 = append(T2, r2)
		}
	}
	c.HTML(
		http.StatusOK,
		"main.html",
		gin.H{
			"Now":       When,
			"Ascendant": Ascendant,
			"Planets":   Planeta,
			"Houses":    Haus,
			"Asc":       Asc,
			"Mc":        Mc,
			"T1":        T1,
			"T2":        T2,
			"title":     "main",
		},
	)

}

func printCycles(c *gin.Context) {
	m := c.DefaultQuery("m", "1")
	y := c.DefaultQuery("y", "1")
	months, _ := strconv.Atoi(m)
	years, _ := strconv.Atoi(y)
	when := time.Now().AddDate(-years, -months, 0)
	if t := c.Query("t"); t != "" {
		if t1, err := httpdate.Str2Time(t, location); err == nil {
			when = t1
		}
	}
	page := components.NewPage()
	page.AddCharts(
		innerPositionChart(when, years, months),
		outerPositionChart(when, years, months),
	)
	page.Render(c.Writer)
}
