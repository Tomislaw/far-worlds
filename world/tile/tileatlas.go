package tile

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var Atlas = TileAtlas{}.Load()

type TileAtlas struct {
	Tiles []Tile `json:"tiles"`
}

func (atlas TileAtlas) Load() TileAtlas {
	path := "tiles.json"
	fmt.Println("Loading tile atlas: " + path)
	jsonFile, err := os.Open(path)

	if err != nil {
		fmt.Println("Failed to load tile list: " + err.Error())
		panic(err.Error())
	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Failed to read tile list: ", err.Error())
		panic(err.Error())
	}

	err = json.Unmarshal(byteValue, &atlas)
	if err != nil {
		fmt.Println("Failed to parse tile list: ", err.Error())
		panic(err.Error())
	}
	return atlas
}

func (atlas *TileAtlas) String() (s string) {
	s += ""
	for key, val := range atlas.Tiles {
		s += fmt.Sprintf("%v=\"{%v,%v,%v}\"\n", key, val.Name, val.MaterialID, val.Block)
	}
	s = strings.TrimSuffix(s, "\n")
	return
}
