package decoration

import (
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type Decoration struct {
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
	fmt.Fprintf(w, "This is decoration page")

}

func ByTypePage(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, "This is %s decoration page", params.ByName("decoration_type"))

}
