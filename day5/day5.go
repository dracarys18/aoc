package main

import (
	"fmt"
	"strings"
	"utils"
)

type Solve interface {
	Parse(filename string) Game
	Solve(game Game) int
}

type Game struct {
	seeds                   []int
	seeds_to_soil           []Mapping
	soil_to_fertilizer      []Mapping
	fertilizer_to_water     []Mapping
	water_to_light          []Mapping
	light_to_temperature    []Mapping
	temperature_to_humidity []Mapping
	humidity_to_location    []Mapping
}

type Mapping struct {
	source int
	dest   int
	count  int
}

type Solution1 struct{}
type Solution2 struct{}

func Map(mapping []Mapping, number int) int {
	for _, mapp := range mapping {
		if mapp.source <= number && number >= mapp.source+mapp.count {
			return mapp.dest + number - mapp.source
		}
	}
	return number
}

func Min(values []int) (min int) {
	min = values[0]
	for _, v := range values {
		if v < min {
			min = v
		}
	}

	return min
}

func HandleMap(game Game) int {
	locations := utils.MapParallel(game.seeds, func(seed int) int {
		soil := Map(game.seeds_to_soil, seed)
		fertilizer := Map(game.soil_to_fertilizer, soil)
		water := Map(game.fertilizer_to_water, fertilizer)
		light := Map(game.water_to_light, water)
		temperature := Map(game.light_to_temperature, light)
		humidity := Map(game.temperature_to_humidity, temperature)
		return Map(game.humidity_to_location, humidity)
	})

	fmt.Println(locations)
	return Min(locations)
}

func parseMapping(lines string, mapping_name string) []Mapping {
	numbers := strings.Split(fmt.Sprintf("%s\n", lines), mapping_name)[1]
	split := strings.Split(numbers, "\n\n")

	mapping := utils.MapParallel(strings.Split(split[0], "\n"), func(line string) Mapping {
		if len(line) != 0 {
			num := utils.AtoiArr(strings.Split(line, " "))
			return Mapping{dest: num[0], source: num[1], count: num[2]}
		}
		return Mapping{}
	})

	return mapping
}

func (_ Solution1) Parse(filename string) Game {
	contents := utils.ReadString(filename)
	splitcontent := strings.Split(contents, "\n\n")
	seeds := utils.AtoiArr(strings.Split(strings.Split(splitcontent[0], "seeds: ")[1], " "))

	seed_to_soil := parseMapping(contents, "seed-to-soil map:")
	soil_to_fertilizer := parseMapping(contents, "soil-to-fertilizer map:")
	fertilizer_to_water := parseMapping(contents, "fertilizer-to-water map:")
	water_to_light := parseMapping(contents, "water-to-light map:")
	light_to_temperature := parseMapping(contents, "light-to-temperature map:")
	temperature_to_humidity := parseMapping(contents, "temperature-to-humidity map:")
	humidity_to_location := parseMapping(contents, "humidity-to-location map:")

	return Game{
		seeds,
		seed_to_soil,
		soil_to_fertilizer,
		fertilizer_to_water,
		water_to_light,
		light_to_temperature,
		temperature_to_humidity,
		humidity_to_location,
	}
}

func (_ Solution1) Solve(game Game) int {
	return HandleMap(game)
}

func ParseAndSolve[T Solve](filename string, problem T) int {
	parsed := problem.Parse(filename)
	return problem.Solve(parsed)
}

func main() {
	fmt.Println(ParseAndSolve("./day5/input.txt", Solution1{}))
}
