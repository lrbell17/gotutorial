package cityutils

type (
	City struct {
		CityName   string
		Country    string
		Population int
	}
)

func NewCity(cityName, country string, population int) (c City) {
	c.CityName = cityName
	c.Country = country
	c.Population = population
	return
}
