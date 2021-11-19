package utils

import "github.com/sirupsen/logrus"

func LogInfo(c int, n string, p float32) {
	logrus.WithFields(logrus.Fields{
		"count": c,
		"price": p,
	}).Info(n)
}

func DupCounter(list []string) map[string]int {
	duplicates := make(map[string]int)

	for _, item := range list {
		_, exist := duplicates[item]

		if exist {
			duplicates[item] += 1
		} else {
			duplicates[item] = 1
		}
	}
	return duplicates
}
