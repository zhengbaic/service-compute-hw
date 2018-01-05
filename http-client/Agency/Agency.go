package Agency

import "time"

type Quoting struct{}

type Weather struct{}

type Destinations struct{}

type Customers struct{}

type Info struct {
	Destinations Destinations
	Quote        Quoting
	Weather      Weather
}

func GetCustomer() Customers {
	time.Sleep(150 * time.Millisecond)
	return Customers{}
}

func GetDestinations(customers Customers) [10]Destinations {
	time.Sleep(250 * time.Millisecond)
	return [10]Destinations{}
}

func GetWeather(destinations Destinations) Weather {
	time.Sleep(330 * time.Millisecond)
	return Weather{}
}


func GetQuote(destination Destinations) Quoting {
	time.Sleep(170 * time.Millisecond)
	return Quoting{}
}
