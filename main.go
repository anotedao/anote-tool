package main

import (
	"log"
	"strings"

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
		if m.TelegramId != 0 {
			var duplicate []*Miner
			db.Where("telegram_id = ?", m.TelegramId).Find(&duplicate)
			basic := &Miner{}
			if len(duplicate) > 1 {
				// log.Printf("%s %d %d", m.Address, m.TelegramId, m.MinedTelegram)
				for _, d := range duplicate {
					if d.MinedTelegram > basic.MinedTelegram {
						basic = d
					}

					if !strings.HasPrefix(basic.Address, "3A") && strings.HasPrefix(d.Address, "3A") {
						basic.Address = d.Address
					}

					if err := db.Save(basic).Error; err != nil {
						log.Println(err)
					}
				}
			}
		}
	}

	log.Println(counter)
}
