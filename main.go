package main

import (
	"fmt"
)

func main() {
	ns := GetNutrionalScore(NutritionalData{
		Energy:  EnergyFromKcal(10),
		SaturatedFattyAcids: SaturatedFattyAcids(500),
		Fiber: FiberGram(4),
		Fruits: FruitsPercentage(60),
		Suger: SugerGram(4),
		Sodium: SodiumFromSalt(0.1),
		IsWater: false,
	}, Food)
	fmt.Println(ns)
	fmt.Println(ns.GetNutriScore())
}