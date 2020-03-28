package main

// Licensing
// Nookipedia's text content is released under a CC BY-SA 3.0 license.
// When using Nookipedia's text in your application, we request that you follow the terms of the license by citing Nookipedia as the source
// of the information, and linking back to the article's history when possible.

// In the spirit of open source and knowledge, we also request that you open source your application (on a site such as GitHub) when possible.

import (
	"fmt"
	"log"
	"net/http"

	"ac-pedia/content/bugs"
	"ac-pedia/content/clothes"
	"ac-pedia/content/decoration"
	"ac-pedia/content/fishes"
	"ac-pedia/content/fossils"
	"ac-pedia/content/fruits"
	"ac-pedia/content/minerals"
	"ac-pedia/content/villagers"

	"github.com/julienschmidt/httprouter"
)

func main() {

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/bugs", bugs.MainPage)
	router.GET("/bugs/:bug_type", bugs.ByTypePage)
	router.GET("/clothes", clothes.MainPage)
	router.GET("/decoration", decoration.MainPage)
	router.GET("/decoration/:decoration_type", decoration.ByTypePage)
	router.GET("/fishes", fishes.MainPage)
	router.GET("/fossils", fossils.MainPage)
	router.GET("/fishes/:water_type", fishes.ByTypePage)
	router.GET("/fruits", fruits.MainPage)
	router.GET("/minerals", minerals.MainPage)
	router.GET("/villagers", villagers.MainPage)
	router.GET("/villagers/:specie", villagers.ByTypePage)

	fmt.Println("Running server on :8080, by Djigo")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Hello")
}
