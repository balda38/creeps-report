package constants

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/balda38/creeps-report/opendotaclient/types"
)

type SeriesType struct {
	ShortName            string `json:"short_name"`
	LongName             string `json:"long_name"`
	RequiredNumberOfWins int    `json:"required_number_of_wins"`
	MaxNumberOfMatches   int    `json:"max_number_of_matches"`
}

func GetHeroes() map[string]types.Hero {
	heroesFile, err := os.Open("constants/heroes.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer heroesFile.Close()

	var heroes map[string]types.Hero
	err = json.NewDecoder(heroesFile).Decode(&heroes)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil
	}

	return heroes
}

func GetSeriesTypes() map[string]SeriesType {
	seriesTypesFile, err := os.Open("constants/series_types.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer seriesTypesFile.Close()

	var seriesTypes map[string]SeriesType
	err = json.NewDecoder(seriesTypesFile).Decode(&seriesTypes)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil
	}

	return seriesTypes
}
