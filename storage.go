package main

var userCounts = make(map[string]int)

func IncreaseCount(userID string) {
	if _, exists := userCounts[userID]; exists {
		userCounts[userID]++
	} else {
		userCounts[userID] = 1
	}
}

func GetTopUser() (string, int) {
	var topUser string
	var maxCount int
	for user, count := range userCounts {
		if count > maxCount {
			maxCount = count
			topUser = user
		}
	}
	return topUser, maxCount
}
