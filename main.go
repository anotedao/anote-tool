package main

import (
	"log"
)

// var db *gorm.DB

var conf *Config

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	conf = initConfig()

	// file, err := os.Open("aint.json")

	// if err != nil {
	// 	log.Printf("Got error while opening config file: %v", err)
	// }

	// body, err := io.ReadAll(file)
	// if err != nil {
	// 	log.Println(err)
	// }

	// var items map[string]uint64
	// err = json.Unmarshal(body, &items)
	// if err != nil {
	// 	fmt.Println("Error parsing JSON: ", err)
	// }

	// for i, _ := range items {
	// 	// log.Printf("%s: %d", i, items[i])
	// 	if i != "3ANmnLHt8mR9c36mdfQVpBtxUs8z1mMAHQW" && items[i] > Multi8 {
	// 		fmt.Printf("%s: %.8f\n", i, float64(items[i])/float64(Multi8))
	// 		// err := sendAsset(items[i], AnoteId, i)
	// 		// if err != nil {
	// 		// 	log.Println()
	// 		// }
	// 	}
	// }

	val := int64(Multi8)
	dataTransaction("%s__price", nil, &val, nil)
}
