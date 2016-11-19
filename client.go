package tvmaze

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Search(query string) []Result {
	url := fmt.Sprintf("http://api.tvmaze.com/search/shows?q=%s", query)
	response, err := http.Get(url)
	if err != nil {
		log.Fatalln("Error: Fetching url failed")
		return nil
	}

	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln("Error: Error reading contents")
		return nil
	}

	var results []Result

	err = json.Unmarshal(content, &results)
	if err != nil {
		log.Fatalln("Error: Error decoding JSON")
		return nil
	}

	return results
}
