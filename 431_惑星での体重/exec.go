package main

import (
	"fmt"
)

func main() {
	getWeightOnEarth := weightFormulaByPlanet("Earth")
	fmt.Println(getWeightOnEarth(50))
	getWeightOnMars := weightFormulaByPlanet("Mars")
	fmt.Println(getWeightOnMars(70))
	getWeightOnJupiter := weightFormulaByPlanet("Jupiter")
	fmt.Println(getWeightOnJupiter(90))

}

type Func func(weight int) string

// 受け取った文字列をもとに関数を返す 関数
func weightFormulaByPlanet(planet string) Func {
	// 惑星名から重力加速度（定数）を選択
	var g float64
	switch planet {
	case "Earth":
		g = 9.8
	case "Mars":
		g = 3.7
	case "Jupiter":
		g = 24.8
	}

	return func(weight int) string {
		w := int(g * float64(weight))
		message := fmt.Sprintf("The weight on %s is %d", planet, w)
		return message
	}
}
