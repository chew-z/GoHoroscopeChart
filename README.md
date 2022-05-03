# GoHoroscopeCharts

http server in Golang for visualizing horoscope data (radix, transit, houses, planetary cycles, etc.) with an astrology charts using splendid [AstroChart](https://github.com/AstroDraw/AstroChart).

It is using [Swiss Ephemeris](https://www.astro.com/swisseph/swephprg.htm) underneath wraped with [swephgo](https://github.com/mshafiee/swephgo). You will need compiled Swiss Ephemem library and ephemeris files on your server. [Read here](https://github.com/chew-z/GoHoroscope) for instructions.

For CLI version check [GoHoroscope](https://github.com/chew-z/GoHoroscope).

## Endpoints

-   /radix
-   /transit
-   /horoscope
-   /cycles

You can pass some parameters in a call like latitude and longitude, time as Unix milliseconds or time as string interpreted by [go-httpdate](https://github.com/songmu/go-httpdate) ("1994-02-03T14:15" does not require encoding). Read [handlers.go](https://github.com/chew-z/GoHoroscopeChart/blob/main/handlers.go) to figure out what you need.

Some other configuration is stored in .env file. Check .env.example

## Screenshot

![Horoscope](images/horoscope.png)
