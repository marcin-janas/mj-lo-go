package main

import (
	"log"
	"net/http"

	svg "github.com/ajstarks/svgo"
)

func main() {
	log.Println("0:8080/mj")
	http.Handle("/mj", http.HandlerFunc(mj))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func mj(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	s.Start(60, 110)
	s.Polyline([]int{30, 10, 20, 20, 30, 40, 40, 50, 30}, []int{40, 20, 10, 20, 30, 20, 10, 20, 40}, "stroke:black", `id="j1"`)
	s.Polyline([]int{30, 10, 10, 30, 50, 50, 30}, []int{60, 40, 30, 50, 30, 40, 60}, "stroke:black", `id="m1"`)
	s.Polyline([]int{20, 10, 10, 20}, []int{60, 70, 50, 60}, "stroke:black", `id="m2"`)
	s.Polyline([]int{40, 50, 50, 40}, []int{60, 50, 70, 60}, "stroke:black", `id="m3"`)
	s.Polyline([]int{30, 10, 20, 20, 30, 40, 40, 50, 30}, []int{100, 80, 70, 80, 90, 80, 70, 80, 100}, "stroke:black", `id="j2"`)
	s.End()
}
