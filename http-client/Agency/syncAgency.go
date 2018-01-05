package Agency

func SyncRun() {
	customers := GetCustomer()
	destinations := GetDestinations(customers)

	var infos [10]Info
	for index, destination := range destinations {
		weather := GetWeather(destination)
		quote := GetQuote(destination)
		infos[index] = Info{Destinations: destination, Quote: quote, Weather: weather}
	}

}
