package villagers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"cgt.name/pkg/go-mwclient"
	"github.com/joho/godotenv"
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

func GetAPIData() {

	// Initialize a *Client with New(), specifying the wiki's API URL
	// and your HTTP User-Agent. Try to use a meaningful User-Agent.
	w, err := mwclient.New("https://nookipedia.com/w/api.php", "Nookpedia")
	if err != nil {
		panic(err)
	}
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	// Log in.
	api_username, _ := os.LookupEnv("USERNAME")
	api_password, _ := os.LookupEnv("PASSWORD")

	err = w.Login(api_username, api_password)
	if err != nil {
		panic(err)
	}

	// Specify parameters to send.
	parameters := map[string]string{
		"action":   "query",
		"list":     "recentchanges",
		"rclimit":  "2",
		"rctype":   "edit",
		"continue": "",
	}

	// Make the request.
	resp, err := w.Get(parameters)
	if err != nil {
		panic(err)
	}

	// Print the *jason.Object
	fmt.Println(resp)

}

func MainPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "This is villagers page")
	GetAPIData()
}

func ByTypePage(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, "This is %s villagers page", params.ByName("specie"))

}
