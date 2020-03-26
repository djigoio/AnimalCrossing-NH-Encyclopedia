package fruits

import (
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type Fruit struct {
	ID           int
	Name         string
	JapaneseName string
	Gender       string
	Species      string
	Personality  string
	Games        []int
	Birthday     *time.Time
}

func FruitsPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "This is fruits page")

}
