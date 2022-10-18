package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/tidwall/gjson"
)

type playerData struct {
	Summary map[string]interface{} `json:"summary"`
	Stats   map[int64]interface{}  `json:"stats"`
}

func loadPlayerDataMap(profileIDList []int) map[int]playerData {
	var playerDataMap = make(map[int]playerData)
	{
		var tmp []string
		for _, profileID := range profileIDList {
			tmp = append(tmp, strconv.Itoa(profileID))
		}
		_url := `https://coh2-api.reliclink.com/community/leaderboard/GetPersonalStat?title=coh2&profile_ids=[` + strings.Join(tmp, ",") + `]`
		resp, err := http.Get(_url)
		if err != nil {
			return nil
		}
		bts, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			return nil
		}

		metadata := gjson.ParseBytes(bts)
		statGroups := metadata.Get("statGroups").Array()
		for _, profileID := range profileIDList {
			var summary interface{}
			leaderboards := map[int64]interface{}{}
			for _, g := range statGroups {
				members := g.Get("members").Array()
				if len(members) == 1 && members[0].Get("profile_id").Int() == int64(profileID) {
					summary = members[0].Value()
					groupID := g.Get("id").Int()
					// fmt.Println()
					// fmt.Println(groupID)
					// fmt.Println(`leaderboardStats.#(statgroup_id==` + strconv.FormatInt(groupID, 10) + `)#`)
					// fmt.Println(len(metadata.Get(`leaderboardStats.#(statgroup_id==` + strconv.FormatInt(groupID, 10) + `)#`).Array()))
					// fmt.Println()
					for _, d := range metadata.Get(`leaderboardStats.#(statgroup_id==` + strconv.FormatInt(groupID, 10) + `)#`).Array() {
						leaderboards[d.Get("leaderboard_id").Int()] = d.Value()
					}
					break
				}
			}
			playerDataMap[profileID] = playerData{
				Summary: summary.(map[string]interface{}),
				Stats:   leaderboards,
			}
		}
	}

	{
		steamNameList := []string{}
		for _, v := range playerDataMap {
			steamNameList = append(steamNameList, `"`+v.Summary["name"].(string)+`"`)
		}
		_url := `https://coh2-api.reliclink.com/community/external/proxysteamuserrequest?title=coh2&profileNames=[` + strings.Join(steamNameList, ",") + `]&request=/ISteamUser/GetPlayerSummaries/v0002/`
		resp, err := http.Get(_url)
		if err != nil {
			return nil
		}
		bts, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			return nil
		}
		metadata := gjson.ParseBytes(bts)
		players := metadata.Get("steamResults.response.players").Array()
		for profileID, playerData := range playerDataMap {
			for _, p := range players {
				if strings.HasSuffix(playerData.Summary["name"].(string), p.Get("steamid").String()) {
					playerDataMap[profileID].Summary["avatar"] = p.Get("avatarfull").String()
				}
			}
		}
	}

	return playerDataMap
}

func loadPlayerSummariesOriginal(steamNameList []string) ([]byte, error) {
	var tmp []string
	for _, steamName := range steamNameList {
		tmp = append(tmp, `"`+steamName+`"`)
	}
	_url := `https://coh2-api.reliclink.com/community/external/proxysteamuserrequest?title=coh2&profileNames=[` + strings.Join(tmp, ",") + `]&request=/ISteamUser/GetPlayerSummaries/v0002/`
	resp, err := http.Get(_url)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}
