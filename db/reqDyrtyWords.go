package db

import (
	"database/sql"
	"fmt"
	"log"

	// Register some standard stuff
	_ "github.com/go-sql-driver/mysql"
)

// DirtyWords возвращает случайное слово из таблицы dirtywords
func DirtyWords() string {
	db, err := sql.Open("mysql", "root:A7bje8971@@/tgbotwords")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from tgbotwords.dirtywords")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	dirtyWords := []WordsDB{}

	for rows.Next() {
		d := WordsDB{}
		err := rows.Scan(&d.id, &d.text)
		if err != nil {
			fmt.Println(err)
			continue
		}
		dirtyWords = append(dirtyWords, d)
	}
	return randWords(dirtyWords)
}
