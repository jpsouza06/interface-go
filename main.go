package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

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

func Create() {
	posturl := "http://localhost:9000/albuns"

	body := []byte(`{
		"title": "Titulo",
		"artist": "Artista",
		"price": 20.50
	}`)

	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	album := &Album{}
	derr := json.NewDecoder(res.Body).Decode(album)
	if derr != nil {
		panic(derr)
	}

	if res.StatusCode != http.StatusCreated {
		panic(res.Status)
	}

	fmt.Println("Ttitle:", album.Title)
	fmt.Println("Artist:", album.Artist)
	fmt.Println("Price:", album.Price)
}

func main() {
	myapp := app.New()
	myWindow := myapp.NewWindow("Web Service")

	myWindow.Resize(fyne.NewSize(500, 500))

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
			Create()
		}))

	myWindow.SetContent(content)
	myWindow.ShowAndRun()

	tidyUp()
}

func tidyUp() {
	fmt.Println("Exited")
}
