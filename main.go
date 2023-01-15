package main

import (
	"ValorantGqlApi/types"
	"encoding/json"
	"github.com/Danny-Dasilva/CycleTLS/cycletls"
	"log"
	"strings"
)

const (
	GET = "GET"
	// POST = "POST"
)

func getPlayerData(name string, tagline string) types.ValData {

	buildTrackerPageUrl := func(n string, t string) string {
		return "https://tracker.gg/valorant/profile/riot/silv%230000/overview"
	}

	url := buildTrackerPageUrl(name, tagline)

	client := cycletls.Init()

	response, err := client.Do(url, cycletls.Options{
		Body:      "",
		Ja3:       "771,4865-4867-4866-49195-49199-52393-52392-49196-49200-49162-49161-49171-49172-51-57-47-53-10,0-23-65281-10-11-35-16-5-51-43-13-45-28-21,29-23-24-25-256-257,0",
		UserAgent: "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:87.0) Gecko/20100101 Firefox/87.0",
	}, GET)
	if err != nil {
		log.Print("Request Failed: " + err.Error())
	}
	data, err := responseToValData(response)
	if err != nil {
		log.Println(err)
	}
	return data
	//for _, v := range data.Segments {
	//	var k string = strconv.FormatFloat(float64(v.Stats.Kills.Value), 'f', -1, 32)
	//	var hs string = strconv.FormatFloat(float64(v.Stats.HeadshotsPercentage.Value), 'f', -1, 32)
	//	var d string = strconv.FormatFloat(float64(v.Stats.Damage.Value), 'f', -1, 32)
	//
	//	fmt.Println(v.Stats.PeakRank.Metadata.TierName, v.Stats.Rank.Metadata.TierName)
	//	fmt.Print(" Kills=" + k + " ")
	//	fmt.Print(" Damage=" + d + " ")
	//	fmt.Println(" HS%=" + hs + " ")
	//}
}

func responseToValData(response cycletls.Response) (data types.ValData, err error) {
	start := strings.Index(response.Body, "\"segments\"")
	end := strings.Index(response.Body, "\"availableSegments\"")

	jsonData := "{" + response.Body[start:end-1] + "}"

	var playerData types.ValData
	err = json.Unmarshal([]byte(jsonData), &playerData)

	if err != nil {
		return types.ValData{}, err
	}
	return playerData, err
}
