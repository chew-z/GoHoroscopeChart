package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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
	houseSystem = "Placidus"
	now := time.Now()
	Now := fmt.Sprintf("\n%s - lat: %.2f, lon: %.2f\n", now.In(location).Format(time.RFC822), lat, lon)
	var Haus []int
	var Planeta []Planet
	var Asc, Mc float64

	if Cusps, Asmc, err := Cusps(now, lat, lon, houseSystem); err != nil {
		log.Println(err.Error())
		return
	} else {
		Asc = Asmc[0]
		Mc = Asmc[1]
		H := Houses(Cusps)
		for _, h := range *H {
			Haus = append(Haus, int(h.DegreeUt))
		}
		B := Bodies(now)
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
			"Now":       Now,
			"title":     "radix",
		},
	)
}

func printTransit(c *gin.Context) {

	houseSystem = "Placidus"
	var Haus1, Haus2 []int
	var Planeta1, Planeta2 []Planet
	var AscNow, McNow float64

	now := time.Now()
	Now := fmt.Sprintf("\n%s - lat: %.2f, lon: %.2f\n", now.In(location).Format(time.RFC822), lat, lon)
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
	when := time.Date(2022, time.Month(4), 30, 22, 27, 55, 0, time.Local)
	Then := fmt.Sprintf("\n%s - lat: %.2f, lon: %.2f\n", when.In(location).Format(time.RFC822), lat, lon)
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
		"transit.html",
		gin.H{
			"Now":         Now,
			"Then":        Then,
			"PlanetsNow":  Planeta1,
			"HousesNow":   Haus1,
			"PlanetsThen": Planeta2,
			"HousesThen":  Haus2,
			"Ascendant":   AscNow,
			"Mc":          McNow,
			"title":       "transit",
		},
	)
}

func PrintHoroscope(c *gin.Context) {
	houseSystem = "Placidus"
	var Haus []int
	var Planeta []Planet
	var Asc, Mc float64
	var T1 []HousePrint
	var T2 []PlanetPrint
	var Ascendant string
	now := time.Now()
	Now := fmt.Sprintf("\n%s - lat: %.2f, lon: %.2f\n", now.In(location).Format(time.RFC822), lat, lon)
	if Cusps, Asmc, err := Cusps(now, lat, lon, houseSystem); err != nil {
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
		B := Bodies(now)
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
			"Now":       Now,
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