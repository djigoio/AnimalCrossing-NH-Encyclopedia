package villagers

import (
	acnhapi "ac-pedia/api"
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
	fmt.Fprintf(w, "%s", acnhapi.GetLinksFromPage("List of villager names in other languages"))
}

func ByTypePage(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, "This is %s villagers page", params.ByName("specie"))

}
