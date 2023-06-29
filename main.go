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

	for _, m := range miners {
		if m.TelegramId != 0 {
			var duplicate []*Miner
			db.Where("telegram_id = ?", m.TelegramId).Find(&duplicate)
			if len(duplicate) > 1 {
				log.Printf("%s %d %d", m.Address, m.TelegramId, m.MinedTelegram)
			}
		}
	}
}
