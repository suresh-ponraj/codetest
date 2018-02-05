package lottery

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Get method will filter the json data based on the ticket type
// and key value and return the response so that it can be showed to the
// customer.
func Get() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		ticketType := r.FormValue("type")
		ticketKey := r.FormValue("key")

		fileContents, _ := ioutil.ReadFile("./json/response.json")

		jsonBody := bytes.NewReader(fileContents)

		var lottery Ticket
		json.NewDecoder(jsonBody).Decode(&lottery)

		var filteredLottery Ticket
		for _, result := range lottery.Results {
			if result.Type == ticketType && result.Key == ticketKey {
				filteredLottery.Results = append(filteredLottery.Results, result)
			}
		}

		response, _ := json.Marshal(&filteredLottery)
		w.Write(response)
	})
}
