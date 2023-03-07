//my first GUI app written in go using fyne.io
//bascially I used chucknorris api and this app diplays a chucknorris joke for every click of a button :DDD

package main

import (
	"encoding/json"
	"image/color"
	"net/http"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var client *http.Client

type chuck struct {
	Text string `json:"value"`
}

func getChuck() (chuck, error) {
	var ch chuck
	resp, err := client.Get("https://api.chucknorris.io/jokes/random")
	if err != nil {
		return chuck{}, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&ch)
	if err != nil {
		return chuck{}, err
	}
	return ch, nil
}

func main() {
	client = &http.Client{Timeout: 5 * time.Second}
	a := app.New()
	w := a.NewWindow("chuck chuck")
	w.Resize(fyne.NewSize(400, 300))

	title := canvas.NewText("get chuck norris joke", color.White)
	title.TextStyle = fyne.TextStyle{
		Bold: true,
	}
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 20

	jokeText := widget.NewLabel("")
	jokeText.Wrapping = fyne.TextWrapWord

	button := widget.NewButton("get joke", func() {
		ch, err := getChuck()
		if err != nil {
			dialog.ShowError(err, w)
		} else {
			jokeText.SetText(ch.Text)
		}
	})

	heightBox := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), button, layout.NewSpacer())
	verticalBox := container.New(layout.NewVBoxLayout(), title, heightBox, widget.NewSeparator(), jokeText)

	w.SetContent(verticalBox)
	w.ShowAndRun()

}
