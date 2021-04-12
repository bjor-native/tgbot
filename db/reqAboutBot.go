package db

import (
	"database/sql"
	"fmt"
	"log"

	// Register some standard stuff
	_ "github.com/go-sql-driver/mysql"
)

// AboutBot возвращает случайное слово из таблицы aboutbot
func AboutBot() string {
	db, err := sql.Open("mysql", "root:A7bje8971@@/tgbotwords")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from tgbotwords.aboutbot")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	aboutbot := []WordsDB{}

	for rows.Next() {
		a := WordsDB{}
		err := rows.Scan(&a.id, &a.text)
		if err != nil {
			fmt.Println(err)
			continue
		}
		aboutbot = append(aboutbot, a)
	}
	return randWords(aboutbot)
}
