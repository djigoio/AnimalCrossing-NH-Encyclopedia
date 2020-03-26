package minerals

import (
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type Minerals struct {
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
	fmt.Fprintf(w, "This is minerals page")

}
