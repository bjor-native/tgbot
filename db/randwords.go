package db

import (
	"math/rand"
	"time"
)

// Принимает на вход слайс заполненый из БД и возвращает рандомное слово из этого слайса
func randWords(words []WordsDB) string {
	rand.Seed(time.Now().UnixNano())
	var result = words[rand.Intn(len(words))]
	return result.text
}
