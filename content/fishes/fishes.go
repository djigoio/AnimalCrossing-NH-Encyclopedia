package fishes

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func FishPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "This is the fish page")

}

func FishByTypePage(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, "This is %s fish page", params.ByName("water_type"))

}
