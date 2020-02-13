package day4

import "strconv"

func Problem1Answer() int64 {
	return CountValidPasswords(264360, 746325, IsValidPassword)
}

func Problem2Answer() int64 {
	return CountValidPasswords(264360, 746325, IsValidPassword2)
}

func CountValidPasswords(from, to int64, f func(s string) bool) int64 {
	var count int64 = 0

	for i := from; i <= to; i++ {
		password := strconv.FormatInt(i, 10)
		if f(password) {
			count += 1
		}
	}

	return count
}

func IsValidPassword2(password string) bool {
	if len(password) != 6 {
		return false
	}

	var hasDouble = false
	for i := 1; i < len(password); {
		if password[i-1] > password[i] {
			return false
		}

		j := 0
		for i < len(password) && password[i-1] == password[i] {
			i += 1
			j += 1
		}
		if j == 0 {
			i += 1
		} else if j == 1 {
			hasDouble = true
		}
	}

	return hasDouble
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
