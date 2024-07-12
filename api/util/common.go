package util

import "time"



func CalculateAge(birthdate time.Time) int {
	today := time.Now()
	age := today.Year() - birthdate.Year()

	// 检查是否已经过了生日
	if today.YearDay() < birthdate.YearDay() {
		age--
	}

	return age
}

