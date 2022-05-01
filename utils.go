package main

import (
	"math"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/mshafiee/swephgo"
)

/* charts */

func innerPositionChart(startTime time.Time, years int, months int) *charts.Line {
	lineChart := charts.NewLine()
	var x []string
	var z, m, v, u []opts.LineData
	start := nod(startTime)
	end := start.AddDate(years, months, 0)
	ipl := swephgo.SeMoon
	for d := start; d.After(end) == false; d = d.AddDate(0, 0, 1) {
		x = append(x, d.In(location).Format("Jan _2, 06"))
		waldo, _ := Waldo(d, ipl, swephgo.SeflgSwieph+swephgo.SeflgRadians)
		z = append(z, opts.LineData{Value: math.Cos(waldo[0])})
		waldo, _ = Waldo(d, ipl+1, swephgo.SeflgSwieph+swephgo.SeflgRadians)
		m = append(m, opts.LineData{Value: math.Cos(waldo[0])})
		waldo, _ = Waldo(d, ipl+2, swephgo.SeflgSwieph+swephgo.SeflgRadians)
		v = append(v, opts.LineData{Value: math.Cos(waldo[0])})
		waldo, _ = Waldo(d, ipl+3, swephgo.SeflgSwieph+swephgo.SeflgRadians)
		u = append(u, opts.LineData{Value: math.Cos(waldo[0])})
	}

	lineChart.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Moon - Mercury - Venus - Mars Position",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Scale: true,
			Min:   -1.0,
			Max:   1.0,
		}),
		charts.WithDataZoomOpts(opts.DataZoom{
			Type:       "slider",
			XAxisIndex: []int{0},
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Show: true,
		}),
		charts.WithToolboxOpts(opts.Toolbox{
			Show:  true,
			Right: "20%",
			Feature: &opts.ToolBoxFeature{
				SaveAsImage: &opts.ToolBoxFeatureSaveAsImage{
					Show:  true,
					Type:  "png",
					Title: "save as image",
				},
			}},
		),
	)
	lineChart.SetXAxis(x).AddSeries("Moon", z)
	lineChart.SetXAxis(x).AddSeries("Mercury", m)
	lineChart.SetXAxis(x).AddSeries("Venus", v)
	lineChart.SetXAxis(x).AddSeries("Mars", u)
	return lineChart
}

func outerPositionChart(startTime time.Time, years int, months int) *charts.Line {
	lineChart := charts.NewLine()
	var x []string
	var j, s, u, n, p []opts.LineData
	start := nod(startTime)
	end := start.AddDate(years, months, 0)
	ipl := swephgo.SeJupiter
	for d := start; d.After(end) == false; d = d.AddDate(0, 0, 1) {
		x = append(x, d.In(location).Format("Jan _2, 06"))
		waldo, _ := Waldo(d, ipl, swephgo.SeflgSwieph+swephgo.SeflgRadians)
		j = append(j, opts.LineData{Value: math.Cos(waldo[0])})
		waldo, _ = Waldo(d, ipl+1, swephgo.SeflgSwieph+swephgo.SeflgRadians)
		s = append(s, opts.LineData{Value: math.Cos(waldo[0])})
		waldo, _ = Waldo(d, ipl+2, swephgo.SeflgSwieph+swephgo.SeflgRadians)
		u = append(u, opts.LineData{Value: math.Cos(waldo[0])})
		waldo, _ = Waldo(d, ipl+3, swephgo.SeflgSwieph+swephgo.SeflgRadians)
		n = append(n, opts.LineData{Value: math.Cos(waldo[0])})
		waldo, _ = Waldo(d, ipl+4, swephgo.SeflgSwieph+swephgo.SeflgRadians)
		p = append(p, opts.LineData{Value: math.Cos(waldo[0])})
	}

	lineChart.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Jupiter - Saturn - Uranus - Neptune - Pluto",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Scale: true,
			Min:   -1.0,
			Max:   1.0,
		}),
		charts.WithDataZoomOpts(opts.DataZoom{
			Type:       "slider",
			XAxisIndex: []int{0},
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Show: true,
		}),
		charts.WithToolboxOpts(opts.Toolbox{
			Show:  true,
			Right: "20%",
			Feature: &opts.ToolBoxFeature{
				SaveAsImage: &opts.ToolBoxFeatureSaveAsImage{
					Show:  true,
					Type:  "png",
					Title: "save as image",
				},
			}},
		),
	)
	lineChart.SetXAxis(x).AddSeries("Jupiter", j)
	lineChart.SetXAxis(x).AddSeries("Saturn", s)
	lineChart.SetXAxis(x).AddSeries("Uranus", u)
	lineChart.SetXAxis(x).AddSeries("Neptune", n)
	lineChart.SetXAxis(x).AddSeries("Pluto", p)
	return lineChart
}

/* general helpers */

func jdToUTC(jd *float64) time.Time {
	year := make([]int, 1)
	month := make([]int, 1)
	day := make([]int, 1)
	hour := make([]float64, 1)
	swephgo.Revjul(*jd, swephgo.SeGregCal, year, month, day, hour)
	h := int(hour[0])
	m := int(60 * (hour[0] - float64(h)))
	utc := time.Date(year[0], time.Month(month[0]), day[0], h, m, 0, 0, time.UTC)
	return utc
}

func jdToLocal(jd *float64) time.Time {
	utc := jdToUTC(jd)
	return utc.In(location)
}

func julian(d time.Time) *float64 {
	h := float64(d.Hour()) + float64(d.Minute())/60 + float64(d.Second())/3600
	jd := swephgo.Julday(d.Year(), int(d.Month()), d.Day(), h, swephgo.SeGregCal)
	return &jd
}

/* Begining of the Day */
func bod(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}

/* Noon of the Day */
func nod(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 12, 0, 0, 0, time.Local)
}

func smallestSignedAngleBetween(x float64, y float64) float64 {
	return math.Min(2.0*math.Pi-math.Abs(x-y), math.Abs(x-y))
}

func fixangle(a float64) float64 {
	return (a - 360*math.Floor(a/360))
}

func rad2deg(r float64) float64 {
	return (r * 180) / math.Pi
}

func deg2rad(d float64) float64 {
	return (d * math.Pi) / 180
}

func bool2int(b bool) int {
	if b {
		return 1
	}
	return 0
}
