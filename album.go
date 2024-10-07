package main

import "github.com/google/uuid"

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// type albumWithDynamicValue struct {
// 	album
// 	Mass string
// }


func newAlbum(title string, artist string, price float64) album {
	return album{
		ID: uuid.New().String(), Title: title, Artist: artist, Price: price,
	}
}

func (album album) format() album {
	album.ID = "success"
	
	return album
}

func (album *album) update() {
	album.ID = "success"
}

// func (album album) addItemToAlbum(key string, value string) {
// 	album := albumWithDynamicValue{
// 		album: album,
//     Mass: value,
// 	}

// 	// return newAlbum
// // 
// 	// album = newAlbum

// 	// newAlbum := make(map[string]interface{})

// 	// newAlbum[key] = value
// }