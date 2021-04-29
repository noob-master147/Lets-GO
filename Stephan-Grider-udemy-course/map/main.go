package main

import (
	"fmt"
)

func main() {
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#00ff00",
	// 	"blue":  "#0000ff",
	// }

	colors := make(map[string]string)
	colors["white"] = "#ffffff"
	colors["red"] = "#ff0000"
	colors["blue"] = "#0000ff"

	// delete(colors, "red")

	for color, hex := range colors {
		fmt.Print(color)
		fmt.Println(hex)
	}

}
