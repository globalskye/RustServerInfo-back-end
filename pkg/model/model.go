package model

import "time"

type User struct {
	Id               int       `json:"id"`
	SteamId          int       `json:"steamId"`
	Name             string    `json:"name"`
	Hwid             string    `json:"hwid"`
	Rank             int       `json:"rank"`
	FirstConnectTime time.Time `json:"firstConnectTime"`
	LastConnectTime  time.Time `json:"lastConnectTime"`
	Balance          int       `json:"balance"`
	KilledPlayers    int       `json:"killedPlayers"`
	KilledMutants    int       `json:"killedMutants"`
	KilledAnimals    int       `json:"killedAnimals"`
	Deaths           int       `json:"deaths"`
	Kits             []KitInfo `json:"kits"`
}
type KitInfo struct {
	Name       string    `json:"name"`
	Countdown  time.Time `json:"countdown"`
	Disposable bool      `json:"disposable"`
}
