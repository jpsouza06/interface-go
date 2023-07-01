package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Album struct {
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func Create(album *Album) {
	posturl := "http://localhost:9000/albuns"

	json_data, err := json.Marshal(album)

	if err != nil {
		panic(err)
	}

	r, err := http.Post(posturl, "application/json", bytes.NewBuffer(json_data))

	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	var res map[string]interface{}

	json.NewDecoder(r.Body).Decode(&res)

	if r.StatusCode != http.StatusCreated {
		panic(r.Status)
	}

	fmt.Println("Title:", album.Title)
	fmt.Println("Artist:", album.Artist)
	fmt.Println("Price:", album.Price)
}

func main() {
	var album Album
	myapp := app.New()
	myWindow := myapp.NewWindow("Web Service")

	myWindow.Resize(fyne.NewSize(500, 300))
	myWindow.CenterOnScreen()

	input := widget.NewEntry()
	input.SetPlaceHolder("Album")

	input1 := widget.NewEntry()
	input1.SetPlaceHolder("Artist")

	input2 := widget.NewEntry()
	input2.SetPlaceHolder("Price")

	content := container.NewVBox(
		input,
		input1,
		input2,
		widget.NewButton("Create", func() {

			album.Title = input.Text
			album.Artist = input1.Text
			if s, err := strconv.ParseFloat(input2.Text, 64); err == nil {
				album.Price = s
			}
			Create(&album)
		}))

	myWindow.SetContent(content)
	myWindow.ShowAndRun()

	tidyUp()
}

func tidyUp() {
	fmt.Println("Exited")
}
