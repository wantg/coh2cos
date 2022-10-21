package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// App struct
type App struct {
	ctx    context.Context
	matchs []match
	db     *gorm.DB
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	dbPath := filepath.Join(Config.ExeDirPath, "cohcos.db")
	_, err := os.Stat(dbPath)
	a.db, _ = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if errors.Is(err, os.ErrNotExist) {
		structures := []interface{}{&player{}, &match{}}
		a.db.AutoMigrate(structures...)
	}
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) StartLogListener() {
	keywords := []string{
		" Setting player (",
		"GAME -- Win Condition Name",
		" Player: ",
		// "Party::SetStatus",
	}
	logListenerStart(Config.LogPath, keywords, a.catchedLog)
	a.sendAllMatchsToUI()
}

func (a *App) catchedLog(logCreatedAt string, logText string) {
	if strings.Contains(logText, "Setting player (") {
		// 生成对局、计算玩家总数
		_tmp := strings.Split(logText, "Setting player (")[1]
		_endPost := strings.Index(_tmp, ")")
		_tmp = _tmp[:_endPost]
		playerIndex, _ := strconv.Atoi(_tmp)
		if playerIndex == 0 {
			a.matchs = append(a.matchs, match{MatchID: md5Encode(logCreatedAt + logText), StartedAt: time.Now()})
		}
		a.matchs[len(a.matchs)-1].PlayerCount = playerIndex + 1
	} else if strings.Contains(logText, "GAME -- Win Condition Name") {
	} else if strings.Contains(logText, " Player: ") {
		// 将玩家信息逐个添加到对局的players数组里
		_tmp := strings.Split(strings.Split(logText, " Player: ")[1], " ")
		slot, _ := strconv.Atoi(_tmp[0])
		profileID, _ := strconv.Atoi(_tmp[len(_tmp)-3])
		side, _ := strconv.Atoi(_tmp[len(_tmp)-2])
		faction := _tmp[len(_tmp)-1]
		alias := strings.Join(_tmp[1:len(_tmp)-3], " ")
		lastMatch := a.matchs[len(a.matchs)-1]
		player := player{
			MatchID:   lastMatch.MatchID,
			Slot:      slot,
			ProfileID: profileID,
			Side:      side,
			Faction:   faction,
			Alias:     alias,
		}
		a.matchs[len(a.matchs)-1].Players = append(lastMatch.Players, player)
		lastMatch = a.matchs[len(a.matchs)-1]
		if len(lastMatch.Players) == lastMatch.PlayerCount {
			a.saveMatch(lastMatch)
		}
	}
	// a.writeLog(logText)
}

func (a *App) saveMatch(m match) {
	// bts, _ := json.MarshalIndent(m, "", "  ")
	// fmt.Println("========")
	// fmt.Println(string(bts))
	var count int64
	a.db.Model(&match{}).Where("match_id = ?", m.MatchID).Count(&count)
	if count == 0 {
		a.db.Create(&m)

		profileIDList := []int{}
		for _, player := range m.Players {
			if player.ProfileID > 0 {
				profileIDList = append(profileIDList, player.ProfileID)
			}
		}
		playerDataMap := loadPlayerDataMap(profileIDList)

		for _, player := range m.Players {
			if playerData, ok := playerDataMap[player.ProfileID]; ok {
				bts, _ := json.Marshal(&playerData.Stats)
				player.Stats = string(bts)
				bts, _ = json.Marshal(&playerData.Summary)
				player.Summary = string(bts)
			}
			a.db.Create(&player)
		}
	}
	a.sendAllMatchsToUI()
}

func (a *App) sendAllMatchsToUI() {
	// fmt.Println("sendAllMatchsToUI")
	matchs := []match{}
	a.db.Order("id desc").Find(&matchs)
	// fmt.Println(&matchs)
	runtime.EventsEmit(a.ctx, "match-updated", &matchs)
}

func (a *App) UpdateConfig(c map[string]interface{}) {
	if strings.Contains(c["lang"].(string), "zh") {
		runtime.WindowSetTitle(a.ctx, getString("zh", "WINDOW_TITLE"))
	}
}

func (a *App) LoadMatchData(matchID string) *match {
	match := match{}
	a.db.Where("match_id = ?", matchID).First(&match)
	a.db.Where("match_id = ?", matchID).Find(&match.Players)
	return &match
}

func (a *App) writeLog(logText string) {
	timeNow := time.Now()
	logWithTimeFlag := "[" + timeNow.Format("2006-01-02 15:04:05") + "] " + logText
	runtimeLogPath := filepath.Join(Config.ExeDirPath, "runtime.log")
	f, err := os.OpenFile(runtimeLogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(logWithTimeFlag); err != nil {
		log.Println(err)
	}
}
