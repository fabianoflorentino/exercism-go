package main

import (
	annalyn "exercism/go/annalyns-infiltration"
	blackjack "exercism/go/blackjack"
	cars "exercism/go/cars-assemble"
	greeting "exercism/go/hello-world"
	lasagna "exercism/go/lasagna"
	weather "exercism/go/weather-forecast"

	"flag"
	"fmt"
)

func main() {
	helloWorldFlag := flag.Bool("hello-world", false, "Run the hello-world exercism")
	annalynsInfiltrationFlag := flag.Bool("annalyns-infiltration", false, "Run the annalyns-infiltration exercism")
	weatherForecastFlag := flag.Bool("weather-forecast", false, "Run the weather-forecast exercism")
	carsAssembleFlag := flag.Bool("cars-assemble", false, "Run the cars-assemble exercism")
	blackJackFlag := flag.Bool("blackjack", false, "Run the blackjack exercism")
	lasagnaFlag := flag.Bool("lasagna", false, "Run the lasagna exercism")

	flag.Parse()

	if *helloWorldFlag {
		exercismHelloWorld()
	}
	if *annalynsInfiltrationFlag {
		exercismAnnalynsInfiltration()
	}
	if *weatherForecastFlag {
		exercismWeatherForecast()
	}
	if *carsAssembleFlag {
		exercismCarsAssemble()
	}
	if *blackJackFlag {
		exercismBlackJack()
	}
	if *lasagnaFlag {
		exercismLasagna()
	}
}

func exercismLasagna() {
	fmt.Println(lasagna.OvenTime)

	fmt.Println(lasagna.RemainingOvenTime(30))
	fmt.Println(lasagna.PreparationTime(2))
	fmt.Println(lasagna.ElapsedTime(30, 2))
}

func exercismHelloWorld() {
	fmt.Println(greeting.HelloWorld())
}

func exercismAnnalynsInfiltration() {
	fmt.Println(annalyn.CanFastAttack(true))
	fmt.Println(annalyn.CanSpy(true, true, true))
	fmt.Println(annalyn.CanSignalPrisoner(true, true))
	fmt.Println(annalyn.CanFreePrisoner(true, true, true, true))
}

func exercismWeatherForecast() {
	fmt.Println(weather.Forecast("Hortolândia", "Sunny"))
}

func exercismCarsAssemble() {
	fmt.Println(cars.CalculateWorkingCarsPerHour(221, 80))
	fmt.Println(cars.CalculateWorkingCarsPerMinute(221, 80))
	fmt.Println(cars.CalculateCost(221))
}

func exercismBlackJack() {
	println(blackjack.ParseCard("ace"))
	println(blackjack.FirstTurn("two", "two", "ace"))
}
