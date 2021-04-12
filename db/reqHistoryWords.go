package db

import (
	"database/sql"
	"fmt"
	"log"

	// Register some standard stuff
	_ "github.com/go-sql-driver/mysql"
)

// HistoryWords возвращает случайное слово из таблицы historywords
func HistoryWords() string {
	db, err := sql.Open("mysql", "root:A7bje8971@@/tgbotwords")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from tgbotwords.historywords")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	historyWords := []WordsDB{}

	for rows.Next() {
		h := WordsDB{}
		err := rows.Scan(&h.id, &h.text)
		if err != nil {
			fmt.Println(err)
			continue
		}
		historyWords = append(historyWords, h)
	}
	return randWords(historyWords)
}
