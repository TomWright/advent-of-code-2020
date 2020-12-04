package main

import (
	"fmt"
	"github.com/tomwright/advent-of-code-2020/util"
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	BirthYear      string
	IssueYear      string
	ExpirationYear string
	Height         string
	HairColour     string
	EyeColour      string
	PassportID     string
	CountryID      string
}

func (p *Passport) Set(field string, val string) {
	switch field {
	case "byr":
		p.BirthYear = val
	case "iyr":
		p.IssueYear = val
	case "eyr":
		p.ExpirationYear = val
	case "hgt":
		p.Height = val
	case "hcl":
		p.HairColour = val
	case "ecl":
		p.EyeColour = val
	case "pid":
		p.PassportID = val
	case "cid":
		p.CountryID = val
	}
}

func (p Passport) IsValidA() bool {
	if p.BirthYear == "" {
		return false
	}
	if p.IssueYear == "" {
		return false
	}
	if p.ExpirationYear == "" {
		return false
	}
	if p.Height == "" {
		return false
	}
	if p.HairColour == "" {
		return false
	}
	if p.EyeColour == "" {
		return false
	}
	if p.PassportID == "" {
		return false
	}
	return true
}

func (p Passport) IsValidB() bool {
	if !p.validBirthYear() {
		return false
	}
	if !p.validIssueYear() {
		return false
	}
	if !p.validExpirationYear() {
		return false
	}
	if !p.validHeight() {
		return false
	}
	if !p.validHairColour() {
		return false
	}
	if !p.validEyeColour() {
		return false
	}
	if !p.validPassportID() {
		return false
	}
	return true
}

func (p Passport) validBirthYear() bool {
	if p.BirthYear == "" {
		return false
	}
	i, err := strconv.Atoi(p.BirthYear)
	if err != nil {
		return false
	}
	if i < 1920 || i > 2002 {
		return false
	}
	return true
}

func (p Passport) validIssueYear() bool {
	if p.IssueYear == "" {
		return false
	}
	i, err := strconv.Atoi(p.IssueYear)
	if err != nil {
		return false
	}
	if i < 2010 || i > 2020 {
		return false
	}
	return true
}

func (p Passport) validExpirationYear() bool {
	if p.ExpirationYear == "" {
		return false
	}
	i, err := strconv.Atoi(p.ExpirationYear)
	if err != nil {
		return false
	}
	if i < 2020 || i > 2030 {
		return false
	}
	return true
}

func (p Passport) validHeight() bool {
	if p.Height == "" {
		return false
	}
	var min, max int
	var val string
	switch {
	case strings.HasSuffix(p.Height, "cm"):
		val = strings.TrimSuffix(p.Height, "cm")
		min = 150
		max = 193
	case strings.HasSuffix(p.Height, "in"):
		val = strings.TrimSuffix(p.Height, "in")
		min = 59
		max = 76
	default:
		return false
	}
	i, err := strconv.Atoi(val)
	if err != nil {
		return false
	}
	if i < min || i > max {
		return false
	}
	return true
}

var hairColourRegexp = regexp.MustCompile(`^#[0-9a-f]{6}$`)

func (p Passport) validHairColour() bool {
	if p.HairColour == "" {
		return false
	}

	if !hairColourRegexp.MatchString(p.HairColour) {
		return false
	}

	return true
}

func (p Passport) validEyeColour() bool {
	switch p.EyeColour {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		return true
	default:
		return false
	}
}

var passportIDRegexp = regexp.MustCompile(`^[0-9]{9}$`)

func (p Passport) validPassportID() bool {
	if p.PassportID == "" {
		return false
	}

	if !passportIDRegexp.MatchString(p.PassportID) {
		return false
	}

	return true
}

func main() {
	passports := make([]*Passport, 0)

	passport := &Passport{}
	if err := util.ParseFileLineByLine("day4/input.txt", func(line string) error {
		if line == "" {
			passports = append(passports, passport)
			passport = &Passport{}
			return nil
		}

		parts := strings.Split(line, " ")
		for _, p := range parts {
			split := strings.Split(p, ":")
			passport.Set(split[0], split[1])
		}

		return nil
	}); err != nil {
		panic(err)
	}
	passports = append(passports, passport)

	validA := 0
	validB := 0
	for _, p := range passports {
		if p.IsValidA() {
			validA++
		}
		if p.IsValidB() {
			validB++
		}
	}

	fmt.Println(validA)
	fmt.Println(validB)
}
