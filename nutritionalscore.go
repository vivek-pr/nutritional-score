package main

type ScoreType int

const (
	Food ScoreType = iota
	Beverage 
	Water 
	Cheese

)
type NutritionalScore struct {
	Value int
	Positive int
	Negative int
	ScoreType ScoreType
}

var scoreToLetter =	[]string{"A", "B", "C", "D", "E"}

type EnergyKJ float64

type SugerGram float64

type SaturatedFattyAcids float64

type SodiumMilligram float64

type FiberGram float64

type proteinGram float64

type FruitsPercentage float64

type NutritionalData struct {
	Energy EnergyKJ
	Suger SugerGram
	SaturatedFattyAcids SaturatedFattyAcids
	Sodium SodiumMilligram
	Fiber FiberGram
	Fruits FruitsPercentage
	protein proteinGram
	IsWater bool
}

var energyLevels = [] float64{
	335,
	670,
	1005,
	1340,
	1675,
	2010,
	2345,
	2680,
	3015,
	3350,
	3685,
	4020,
	4355,
	4690}
var sugerLevels = [] float64{
	4.5,
	9,
	13.5,
	18,
	22.5,
	27,
	31.5,
	36,
	40.5,
	45,
	49.5,
	54,
	58.5,
	63,
}
var saturatedFattyAcidsLevels = [] float64{
	1,
	2,
	3,
	4,
	5,
	6,
	7,
	8,
	9,
	10,
}
var sodiumLevels = [] float64{
	90,
	180,
	270,
	360,
	450,
	540,
	630,
	720,
	810,
	900,
}
var fiberLevels = [] float64{
	0.9,
	1.8,
	2.7,
	3.6,
	4.5,
	5.4,
	6.3,
	7.2,
	8.1,
	9,
	9.9,
	10.8,
	11.7,
	12.6,
	13.5,
}
var proteinLevels = [] float64{
	1.6,
	3.2,
	4.8,
	6.4,
	8,
	9.6,
	11.2,
	12.8,
}
var fruitsLevels = [] float64{
	40,
	60,
	80,
	100,
	120,
	140,
	160,

}

var energyLevelsBeverage = [] float64{
	0,
	30,
	60,
	90,
	120,
	150,
	180,
	210,
	240,
	270,
	300,
	330,
	360,
}
var sugerLevelsBeverage = [] float64{
	0.5,
	1,
	1.5,
	2,
	2.5,
	3,
	3.5,
	4,
	4.5,
	5,
	5.5,
	6,
	6.5,
}

func (e EnergyKJ) GetPoints(st ScoreType) int {
	switch st {
	case Food:
		return getPointsFromRange(float64(e), energyLevels)
	case Beverage:
		return getPointsFromRange(float64(e), energyLevelsBeverage)
	}
	return 0

}
func (s SugerGram) GetPoints(st ScoreType) int {
	switch st {
	case Food:
		return getPointsFromRange(float64(s), sugerLevels)
	case Beverage:
		return getPointsFromRange(float64(s), sugerLevelsBeverage)
	}
	return 0
}

func (s SaturatedFattyAcids) GetPoints(st ScoreType) int {
	return getPointsFromRange(float64(s), saturatedFattyAcidsLevels)
}

func (s SodiumMilligram) GetPoints(st ScoreType) int {
	return getPointsFromRange(float64(s), sodiumLevels)
}

func (s FiberGram) GetPoints(st ScoreType) int {
	return getPointsFromRange(float64(s), fiberLevels)
}

func (s proteinGram) GetPoints(st ScoreType) int {
	return getPointsFromRange(float64(s), proteinLevels)
}

func (f FruitsPercentage) GetPoints(st ScoreType) int {
	if st == Beverage{
		if f > 80 {
			return 10
		} else if f > 60{
			return 4

		}else if f >40{
			return 2
		}
		return 0
	}else{
		if f > 80 {
			return 5
		} else if f > 60{
			return 2

		}else if f >40{
			return 1
		}
		return 0

	}
}

func EnergyFromKcal(kcal float64) EnergyKJ {
	return EnergyKJ(kcal * 4.184)
}

func SodiumFromSalt(salt float64) SodiumMilligram {
	return SodiumMilligram(salt / 2.54)
}

func ProteinGram() proteinGram {
	return proteinGram(0)
}

func FatGram() SaturatedFattyAcids {
	return SaturatedFattyAcids(0)
}



func GetNutrionalScore(data NutritionalData, scoreType ScoreType) NutritionalScore {
	value := 0
	positive := 0
	negative := 0
	if data.IsWater {
		return NutritionalScore{
			Value: 0,
			Positive: 0,
			Negative: 0,
			ScoreType: scoreType,
		}
	}

	fruitsPointsdata := data.Fruits.GetPoints(scoreType)
	fiberPoints := data.Fiber.GetPoints(scoreType)
	negative += data.Sodium.GetPoints(scoreType) + data.SaturatedFattyAcids.GetPoints(scoreType) + data.Suger.GetPoints(scoreType) + data.Energy.GetPoints(scoreType) 
	positive += data.protein.GetPoints(scoreType) + fruitsPointsdata + fiberPoints
	if scoreType == Cheese{
		value = negative - positive
	} else{
		if negative >= 11 && fruitsPointsdata<5{
			value = negative - positive - fiberPoints
		}else{
			value = negative - positive
		}
	}
	
	return NutritionalScore{
		Value: value,
		Positive: positive,
		Negative: negative,
		ScoreType: scoreType,
	}
}

func (n NutritionalScore) GetNutriScore() string {
	if n.ScoreType == Food{
		return scoreToLetter[getPointsFromRange(float64(n.Value), []float64{-1, 2, 10, 18, 26, 34, 42, 50, 58, 66, 74, 82, 90})]
	}
	if n.ScoreType == Water{
		return scoreToLetter[0]

	}
	return scoreToLetter[getPointsFromRange(float64(n.Value), []float64{-1, 2, 10, 18, 26, 34, 42, 50, 58, 66, 74, 82, 90})]
}

func getPointsFromRange(value float64, levels []float64) int {
	for i, level := range levels {
		if value < level {
			return i
		}
	}
	return len(levels)
}