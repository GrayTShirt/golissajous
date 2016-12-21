package main

import (
	"log"
	"net/http"
	"strconv"
	"github.com/graytshirt/server2/lissajous"
)


func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	settings := lissajous.New()
	for k, v := range r.URL.Query() {
		switch {
			case k == "size":
				temp, err := strconv.Atoi(v[0])
				if err != nil {
				}
				settings.Size = temp
			case k == "cycles":
				temp, err := strconv.ParseFloat(v[0], 64)
				if err != nil {
				}
				settings.Cycles = temp
			case k == "res":
				temp, err := strconv.ParseFloat(v[0], 64)
				if err != nil {
				}
				settings.Res = temp
			case k == "nframes":
				temp, err := strconv.Atoi(v[0])
				if err != nil {
				}
				settings.Nframes = temp
			case k == "delay":
				temp, err := strconv.Atoi(v[0])
				if err != nil {
				}
				settings.Delay = temp
		}
	}
	lissajous.RenderGif(w, settings)
}
