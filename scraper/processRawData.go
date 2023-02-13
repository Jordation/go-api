package scraper

import "fmt"

var statCategories = []string{"Rating", "ACS", "Kills", "Deaths", "Assists", "", "KAST", "ADR", "HS%", "FK", "FD", "", "Player"}

func getMappedData(dataset []string, player string) map[string]string {
	md := make(map[string]string)
	cd := append(dataset, player)

	for i, v := range statCategories {
		switch v {
		case "":
		default:
			md[v] = cd[i]
		}
	}

	return md
}

func labelAllData(data []string, players []string) []map[string]string {
	var labeledData []map[string]string
	for i := 0; i < len(data); i += 12 {
		// sending in data[i:i+12] was sending a pointer i think and causing the player to be added to the entire dataset
		dg := make([]string, 12)
		copy(dg, data[i:i+12])

		labeledData = append(labeledData, getMappedData(dg, players[i/12]))
	}
	return labeledData
}

//ProcessRawData process the raw scraped data
func ProcessRawData(rawData RawMatchData) {

	// validate shape of data

	fmt.Println("Hello")
}
