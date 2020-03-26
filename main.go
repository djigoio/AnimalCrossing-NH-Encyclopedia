package main

import (
    "fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"log"
	"./content/bugs"
	"./content/clothes"
	"./content/decoration"
	"./content/fishes"
	"./content/fossils"
	"./content/fruits"
	"./content/minerals"
	"./content/villagers"

)

func main() {


	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/bugs", bugs.BugsPage)
	router.GET("/bugs/:bug_type", bugs.BugsByTypePage)
	router.GET("/clothes", clothes.ClothesPage)
	router.GET("/decoration", decoration.DecorationPage)
	router.GET("/decoration/:decoration_type", decoration.DecorationByTypePage)
	router.GET("/fishes", fishes.FishPage)
	router.GET("/fossils", fossils.FossilPage)
	router.GET("/fishes/:water_type", fishes.FishByTypePage)
	router.GET("/fruits", fruits.FruitsPage)
	router.GET("/minerals", minerals.MineralsPage)
	router.GET("/villagers", villagers.VillagersPage)
	router.GET("/villagers/:specie", villagers.VillagersByTypePage)


	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprintf(w, "Hello")
}