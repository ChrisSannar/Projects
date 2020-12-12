package main

import (
	"fmt"
	"strings"
	"strconv"
	"regexp"
	"../util"
)

func main() {
	text := util.ReadFileToString("./input.txt")
	maps := parseToMaps(text)
	count := 0
	for _, mappers := range maps {
		if (validMap(mappers)) {
			count++
		}
	}
	fmt.Println("Count:", count)
}

func parseToMaps(text string) []map[string]string {
	entries := strings.Split(text, "\n\n")
	maps := make([]map[string]string, len(entries))
	for i, val := range entries {
		mappers := parseSingleEntryToMap(val)
		maps[i] = mappers
	}
	return maps
}

func parseSingleEntryToMap(text string) map[string]string {
	whitespaceRegex := regexp.MustCompile("[^\\s]+")
	var mappers map[string]string = make(map[string]string)
	split := whitespaceRegex.FindAllString(text, -1)
	for _, val := range split {
		keyValueSplit := strings.Split(val, ":");
		key, value := keyValueSplit[0], keyValueSplit[1]
		mappers[key] = value
	}
	return mappers
}

func validMap(mappers map[string]string) bool {
	validKeys := []string{"ecl", "pid", "eyr", "hcl", "byr", "iyr", "hgt"}
	for _, validKey := range validKeys {
		if _, ok := mappers[validKey]; !ok {
			return false
		}
	}
	var valid bool = true
	valid = valid && validYear(mappers["byr"], 1920, 2002)
	valid = valid && validYear(mappers["iyr"], 2010, 2020)
	valid = valid && validYear(mappers["eyr"], 2020, 2030)
	valid = valid && validHeight(mappers["hgt"])
	validHairColor, _ := regexp.MatchString("^#[0-9a-f]{6}$", mappers["hcl"])
	valid = valid && validHairColor
	valid = valid && validEyeColor(mappers["ecl"])
	validPassportId, _ := regexp.MatchString("^[0-9]{9}$", mappers["pid"])
	valid = valid && validPassportId
	return valid
}

func validYear(year string, low int, high int) bool {
	if len(year) != 4 {
		return false
	} else {
		intYear, _ := strconv.Atoi(year)
		if intYear >= low && intYear <= high {
			return true
		} else {
			return false
		}
	}
}

func validHeight(height string) bool {
	divideNumsAndLetters := regexp.MustCompile("\\d+|\\D+")
	split := divideNumsAndLetters.FindAllString(height, -1)
	if len(split) != 2 {
		return false
	}
	heightNum, _ := strconv.Atoi(split[0])
	if split[1] == "in" {
		if heightNum >= 59 && heightNum <= 76 {
			return true
		}
	} else {
		if heightNum >= 150 && heightNum <= 193 {
			return true
		}
	}
	return false
}

func validEyeColor(color string) bool {
	return color == "amb" ||
		color == "blu" ||
		color == "brn" ||
		color == "gry" ||
		color == "grn" ||
		color == "hzl" ||
		color == "oth"
}