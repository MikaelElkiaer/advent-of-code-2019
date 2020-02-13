package day4

import "strconv"

func Problem1Answer() int64 {
	return CountValidPasswords(264360, 746325)
}

func CountValidPasswords(from, to int64) int64 {
	var count int64 = 0

	for i := from; i <= to; i++ {
		password := strconv.FormatInt(i, 10)
		if IsValidPassword(password) {
			count += 1
		}
	}

	return count
}

func IsValidPassword(password string) bool {
	if len(password) != 6 {
		return false
	}

	var hasDouble = false
	for i := 1; i < len(password); i++ {
		if password[i-1] > password[i] {
			return false
		}

		if password[i-1] == password[i] {
			hasDouble = true
		}
	}

	return hasDouble
}