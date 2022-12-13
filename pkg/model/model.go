package model

import "time"

type User struct {
	Id               int                  `json:"id"`
	SteamId          int                  `json:"steamId"`
	Name             string               `json:"name"`
	Hwid             string               `json:"hwid"`
	Rank             int                  `json:"rank"`
	FristConnectTime time.Time            `json:"fristConnectTime"`
	LastConnectTime  time.Time            `json:"lastConnectTime"`
	Balance          int                  `json:"balance"`
	KilledPlayers    int                  `json:"killedPlayers"`
	KilledMutants    int                  `json:"killedMutants"`
	KilledAnimals    int                  `json:"killedAnimals"`
	Deaths           int                  `json:"deaths"`
	Kits             map[string]time.Time `json:"kits"`
}
