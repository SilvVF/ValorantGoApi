package types

import "time"

type ValData struct {
	Segments []struct {
		Type       string `json:"type"`
		Attributes struct {
			SeasonID string `json:"seasonId"`
			Playlist string `json:"playlist"`
		} `json:"attributes"`
		Metadata struct {
			Name         string    `json:"name"`
			PlaylistName string    `json:"playlistName"`
			StartTime    time.Time `json:"startTime"`
			EndTime      time.Time `json:"endTime"`
			Schema       string    `json:"schema"`
		} `json:"metadata"`
		ExpiryDate time.Time `json:"expiryDate"`
		Stats      struct {
			MatchesPlayed struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"matchesPlayed"`
			MatchesWon struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"matchesWon"`
			MatchesLost struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"matchesLost"`
			MatchesTied struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"matchesTied"`
			MatchesWinPct struct {
				Rank            interface{} `json:"rank"`
				Percentile      float32     `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Value           float32     `json:"value"`
				DisplayValue    string      `json:"displayValue"`
				DisplayType     string      `json:"displayType"`
			} `json:"matchesWinPct"`
			MatchesDuration struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"matchesDuration"`
			TimePlayed struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"timePlayed"`
			RoundsPlayed struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"roundsPlayed"`
			RoundsWon struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"roundsWon"`
			RoundsLost struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"roundsLost"`
			RoundsWinPct struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"roundsWinPct"`
			RoundsDuration struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"roundsDuration"`
			Score struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"score"`
			ScorePerMatch struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"scorePerMatch"`
			ScorePerRound struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"scorePerRound"`
			Kills struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"kills"`
			KillsPerRound struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"killsPerRound"`
			KillsPerMatch struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"killsPerMatch"`
			Deaths struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"deaths"`
			DeathsPerRound struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"deathsPerRound"`
			DeathsPerMatch struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"deathsPerMatch"`
			Assists struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"assists"`
			AssistsPerRound struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"assistsPerRound"`
			AssistsPerMatch struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"assistsPerMatch"`
			KDRatio struct {
				Rank            interface{} `json:"rank"`
				Percentile      float32     `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Value           float32     `json:"value"`
				DisplayValue    string      `json:"displayValue"`
				DisplayType     string      `json:"displayType"`
			} `json:"kDRatio"`
			KDARatio struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"kDARatio"`
			KADRatio struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"kADRatio"`
			Damage struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"damage"`
			DamageDelta struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"damageDelta"`
			DamagePerRound struct {
				Rank            interface{} `json:"rank"`
				Percentile      float32     `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Value           float32     `json:"value"`
				DisplayValue    string      `json:"displayValue"`
				DisplayType     string      `json:"displayType"`
			} `json:"damagePerRound"`
			DamagePerMatch struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"damagePerMatch"`
			DamagePerMinute struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"damagePerMinute"`
			DamageReceived struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"damageReceived"`
			Headshots struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"headshots"`
			HeadshotsPerRound struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"headshotsPerRound"`
			HeadshotsPercentage struct {
				Rank            interface{} `json:"rank"`
				Percentile      float32     `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Value           float32     `json:"value"`
				DisplayValue    string      `json:"displayValue"`
				DisplayType     string      `json:"displayType"`
			} `json:"headshotsPercentage"`
			GrenadeCasts struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"grenadeCasts"`
			GrenadeCastsPerRound struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"grenadeCastsPerRound"`
			GrenadeCastsPerMatch struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"grenadeCastsPerMatch"`
			Ability1Casts struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"ability1Casts"`
			Ability1CastsPerRound struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"ability1CastsPerRound"`
			Ability1CastsPerMatch struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"ability1CastsPerMatch"`
			Ability2Casts struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"ability2Casts"`
			Ability2CastsPerRound struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"ability2CastsPerRound"`
			Ability2CastsPerMatch struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"ability2CastsPerMatch"`
			UltimateCasts struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"ultimateCasts"`
			UltimateCastsPerRound struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"ultimateCastsPerRound"`
			UltimateCastsPerMatch struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"ultimateCastsPerMatch"`
			DealtHeadshots struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"dealtHeadshots"`
			DealtBodyshots struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"dealtBodyshots"`
			DealtLegshots struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"dealtLegshots"`
			ReceivedHeadshots struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"receivedHeadshots"`
			ReceivedBodyshots struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"receivedBodyshots"`
			ReceivedLegshots struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"receivedLegshots"`
			EconRating struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"econRating"`
			EconRatingPerMatch struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"econRatingPerMatch"`
			EconRatingPerRound struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"econRatingPerRound"`
			Suicides struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"suicides"`
			FirstBloods struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"firstBloods"`
			FirstBloodsPerMatch struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"firstBloodsPerMatch"`
			FirstDeaths struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"firstDeaths"`
			LastDeaths struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"lastDeaths"`
			Survived struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"survived"`
			Traded struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"traded"`
			KAST struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"kAST"`
			MostKillsInMatch struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"mostKillsInMatch"`
			Flawless struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"flawless"`
			Thrifty struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"thrifty"`
			Aces struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"aces"`
			TeamAces struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"teamAces"`
			Clutches struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"clutches"`
			ClutchesLost struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"clutchesLost"`
			Clutches1V1 struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"clutches1v1"`
			Clutches1V2 struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"clutches1v2"`
			Clutches1V3 struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"clutches1v3"`
			Clutches1V4 struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"clutches1v4"`
			Clutches1V5 struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"clutches1v5"`
			ClutchesLost1V1 struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"clutchesLost1v1"`
			ClutchesLost1V2 struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"clutchesLost1v2"`
			ClutchesLost1V3 struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"clutchesLost1v3"`
			ClutchesLost1V4 struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"clutchesLost1v4"`
			ClutchesLost1V5 struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"clutchesLost1v5"`
			Kills1K struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"kills1K"`
			Kills2K struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"kills2K"`
			Kills3K struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"kills3K"`
			Kills4K struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"kills4K"`
			Kills5K struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"kills5K"`
			Kills6K struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"kills6K"`
			Plants struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"plants"`
			PlantsPerMatch struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"plantsPerMatch"`
			PlantsPerRound struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"plantsPerRound"`
			AttackKills struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"attackKills"`
			AttackDeaths struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"attackDeaths"`
			AttackKDRatio struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"attackKDRatio"`
			AttackAssists struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"attackAssists"`
			AttackRoundsWon struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"attackRoundsWon"`
			AttackRoundsLost struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"attackRoundsLost"`
			AttackRoundsWinPct struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"attackRoundsWinPct"`
			AttackScore struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"attackScore"`
			AttackDamage struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"attackDamage"`
			AttackHeadshots struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"attackHeadshots"`
			AttackTraded struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"attackTraded"`
			AttackSurvived struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"attackSurvived"`
			AttackFirstBloods struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"attackFirstBloods"`
			AttackFirstDeaths struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"attackFirstDeaths"`
			AttackKAST struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"attackKAST"`
			Defuses struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"defuses"`
			DefusesPerMatch struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"defusesPerMatch"`
			DefusesPerRound struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"defusesPerRound"`
			DefenseKills struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"defenseKills"`
			DefenseDeaths struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"defenseDeaths"`
			DefenseKDRatio struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"defenseKDRatio"`
			DefenseAssists struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"defenseAssists"`
			DefenseRoundsWon struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"defenseRoundsWon"`
			DefenseRoundsLost struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"defenseRoundsLost"`
			DefenseRoundsWinPct struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"defenseRoundsWinPct"`
			DefenseScore struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"defenseScore"`
			DefenseDamage struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"defenseDamage"`
			DefenseHeadshots struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"defenseHeadshots"`
			DefenseTraded struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"defenseTraded"`
			DefenseSurvived struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"defenseSurvived"`
			DefenseFirstBloods struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"defenseFirstBloods"`
			DefenseFirstDeaths struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"defenseFirstDeaths"`
			DefenseKAST struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				DisplayName     string  `json:"displayName"`
				DisplayCategory string  `json:"displayCategory"`
				Category        string  `json:"category"`
				Value           float32 `json:"value"`
				DisplayValue    string  `json:"displayValue"`
				DisplayType     string  `json:"displayType"`
			} `json:"defenseKAST"`
			Rank struct {
				//Rank            interface{} `json:"rank"`
				//Percentile      interface{} `json:"percentile"`
				//DisplayName     string      `json:"displayName"`
				//DisplayCategory interface{} `json:"displayCategory"`
				//Category        string      `json:"category"`
				Metadata struct {
					IconURL  string `json:"iconUrl"`
					TierName string `json:"tierName"`
				} `json:"metadata"`
				//Value        interface{} `json:"value"`
				//DisplayValue string      `json:"displayValue"`
				//DisplayType  string      `json:"displayType"`
			} `json:"rank"`
			PeakRank struct {
				//Rank            string `json:"rank"`
				//Percentile      string `json:"percentile"`
				//DisplayName     string `json:"displayName"`
				//DisplayCategory string `json:"displayCategory"`
				//Category        string `json:"category"`
				Metadata struct {
					IconURL  string `json:"iconUrl"`
					TierName string `json:"tierName"`
					ActID    string `json:"actId"`
					ActName  string `json:"actName"`
				} `json:"metadata"`
				//Value        interface{} `json:"value"`
				//DisplayValue string      `json:"displayValue"`
				//DisplayType  string      `json:"displayType"`
			} `json:"peakRank"`
		} `json:"stats"`
	} `json:"segments"`
}
