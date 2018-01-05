package Agency

func AsyncRun() {
	customer := GetCustomer()
	destinations := GetDestinations(customer)
	var infos [10]Info

	quotes := [10]chan Quoting{}
	weathers := [10]chan Weather{}

	for i := range weathers {
		weathers[i] = make(chan Weather)
	}

	for i := range quotes {
		quotes[i] = make(chan Quoting)
	}

	for index, dest := range destinations {
		i := index
		d := dest
		
		go func() {
			quotes[i] <- GetQuote(d)
		}()

		go func() {
			weathers[i] <- GetWeather(d)
		}()
	}

	for index, d := range destinations {
		infos[index] = Info{Destinations: d, Quote: <-quotes[index], Weather: <-weathers[index]}
	}
}
