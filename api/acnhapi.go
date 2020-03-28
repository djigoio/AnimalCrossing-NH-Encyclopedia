package acnhapi

// ADD IF JSON FILE WITH INFO EMPTY, THEN GET API CALL?
import (
	"cgt.name/pkg/go-mwclient"
	"github.com/antonholmquist/jason"
)

func GetAPIData(parameters map[string]string) *jason.Object {

	w, _ := mwclient.New("https://nookipedia.com/w/api.php", "Nookpedia")
	response, _ := w.Get(parameters)
	return response
}

func GetLinksFromPage(PageName string) *jason.Object {

	parameters := map[string]string{
		"action":  "query",
		"format":  "json",
		"prop":    "links",
		"titles":  PageName,
		"pllimit": "500",
	}

	return GetAPIData(parameters)
}
