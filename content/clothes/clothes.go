package clothes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type Clothing struct {
	ID           int
	Name         string
	JapaneseName string
	Gender       string
	Species      string
	Personality  string
	Games        []int
	Birthday     *time.Time
}

func ClothesPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "This is clothing page")

}

func ClothesByTypePage(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, "This is %s Clothes page", params.ByName("clothe_type"))

}
