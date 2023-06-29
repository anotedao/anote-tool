package main

import (
	"log"

	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	db = initDb()

	var miners []*Miner
	db.Find(&miners)

	counter := 0

	for _, m := range miners {
		if m.TelegramId == 0 {
			m.TelegramId = int64(m.ID)
			db.Save(m)
			counter++
		}
	}

	log.Println(counter)
}
