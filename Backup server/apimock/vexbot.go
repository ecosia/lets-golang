package apimock

import (
	"encoding/json"
	"net/http"
)

type Point struct {
	X int
	Y int
}

type Vector struct {
	A     Point
	B     Point
	Speed int
}

type Vexbot struct {
	Vectors []Vector
}

func VexbotHandler(w http.ResponseWriter, r *http.Request) {
	a := Point{X: 444, Y: 859}
	b := Point{X: 866, Y: 956}
	vectors := Vexbot{Vectors: []Vector{Vector{A: a, B: b, Speed: 73}}}
	js, err := json.Marshal(vectors)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
