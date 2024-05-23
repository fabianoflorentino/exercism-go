package main

import (
	annalyn "exercism/go/annalyns-infiltration"
	cars "exercism/go/cars-assemble"
	greeting "exercism/go/hello-world"
	weather "exercism/go/weather-forecast"

	"exercism/go/lasagna"
	"fmt"
	"time"
)

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

func main() {
	// exercismHelloWorld()
	// exercismLasagna()
	// exercismAnnalynsInfiltration()
	// exercismCarsAssemble()
	// exercismWeatherForecast()
	fmt.Println(time.Now().Local())
}
