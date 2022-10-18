package main

import (
	"path/filepath"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Model struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"-" `
	// UpdatedAt time.Time      `json:"-" `
	// DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type match struct {
	Model
	MatchID     string    `json:"matchID"     gorm:"index"`
	StartedAt   time.Time `json:"startedAt"`
	PlayerCount int       `json:"playerCount"`
	Players     []player  `json:"players"     gorm:"-"`
}

type player struct {
	Model
	MatchID   string `json:"matchID"   gorm:"index"`
	Slot      int    `json:"slot"`
	ProfileID int    `json:"profileID" gorm:"index"`
	Side      int    `json:"side"`
	Faction   string `json:"faction"`
	Alias     string `json:"alias"`
	Summary   string `json:"summary"`
	Stats     string `json:"stats"`
}

func main() {
	// executablePath, _ := os.Executable()
	// exeDirPath := filepath.Join(filepath.Dir(executablePath))

	exeDirPath := "./"
	db, _ := gorm.Open(sqlite.Open(filepath.Join(exeDirPath, "./build/cohcos-init.db")), &gorm.Config{})

	structures := []interface{}{&player{}, &match{}}
	db.AutoMigrate(structures...)
}
