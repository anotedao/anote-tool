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
		if m.TelegramId != 0 {
			var duplicate []*Miner
			db.Where("telegram_id = ?", m.TelegramId).Find(&duplicate)
			// basic := &Miner{}
			biggest := uint64(0)
			if len(duplicate) > 1 {
				for _, d := range duplicate {
					if d.MinedTelegram > biggest {
						biggest = d.MinedTelegram
					}

					log.Printf("%s %d %d", d.Address, d.TelegramId, d.MinedTelegram)
				}

				for _, d := range duplicate {
					if d.MinedTelegram < biggest {
						d.MinedTelegram = biggest
						if err := db.Save(d).Error; err != nil {
							log.Println(err)
						}
					}

					log.Printf("%s %d %d", d.Address, d.TelegramId, d.MinedTelegram)
				}
				// 	if !strings.HasPrefix(basic.Address, "3A") && strings.HasPrefix(d.Address, "3A") {
				// 		basic.Address = d.Address
				// 	}

				// 	if err := db.Save(basic).Error; err != nil {
				// 		log.Println(err)
				// 	}
				// }
				counter++
			}
		}
	}

	log.Println(counter)
}
