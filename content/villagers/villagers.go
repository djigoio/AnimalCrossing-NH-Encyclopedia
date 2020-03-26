package villagers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type Villager struct {
	ID           int
	Name         string
	JapaneseName string
	Gender       string
	Species      string
	Personality  string
	Games        []int
	Birthday     *time.Time
}

func MainPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "This is villagers page")

}

func ByTypePage(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, "This is %s villagers page", params.ByName("specie"))

}
