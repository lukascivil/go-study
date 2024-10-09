package types

import "github.com/google/uuid"

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// type albumWithDynamicValue struct {
// 	album
// 	Mass string
// }


func NewAlbum(title string, artist string, price float64) Album {
	return Album{
		ID: uuid.New().String(), Title: title, Artist: artist, Price: price,
	}
}

func (album Album) Format() Album {
	album.ID = "success"
	
	return album
}

func (album *Album) Update() {
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