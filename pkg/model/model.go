package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	Id               primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	SteamId          int                `json:"steamId" bson:"steamId"`
	ClanName         string             `json:"clanName" bson:"clanName"`
	Name             string             `json:"name" bson:"name"`
	Hwid             string             `json:"hwid" bson:"hwid"`
	Rank             int                `json:"rank" bson:"rank"`
	FirstConnectTime time.Time          `json:"firstConnectTime" bson:"firstConnectTime"`
	LastConnectTime  time.Time          `json:"lastConnectTime" bson:"lastConnectTime"`
	Balance          int                `json:"balance" bson:"balance"`
	KilledPlayers    int                `json:"killedPlayers" bson:"killedPlayers"`
	KilledMutants    int                `json:"killedMutants" bson:"killedMutants"`
	KilledAnimals    int                `json:"killedAnimals" bson:"killedAnimals"`
	Deaths           int                `json:"deaths" bson:"deaths"`
	Kits             []KitInfo          `json:"kits" bson:"kits"`
	Farm             UserTopFarm        `json:"farm" bson:"farm"`
	Online           float32            `json:"online" bson:"online"`
	Raid             float32            `json:"raid" bson:"raid"`
}
type UserTopFarm struct {
	Id      primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	SteamId int                `json:"steamId" bson:"steamId"`
	Wood    int                `json:"wood" bson:"wood"`
	Metal   int                `json:"metal" bson:"metal"`
	Sulfur  int                `json:"sulfur" bson:"sulfur"`
	Leather int                `json:"leather" bson:"leather"`
	Cloth   int                `json:"cloth" bson:"cloth"`
	Fat     int                `json:"fat" bson:"fat"`
}
type KitInfo struct {
	Name       string    `json:"name" bson:"name"`
	Countdown  time.Time `json:"countdown" bson:"countdown"`
	Disposable bool      `json:"disposable" bson:"disposable"`
}
type Online struct {
	Id    primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Count int                `bson:"count" json:"count"`
	Users []User             `json:"users" bson:"users"`
}
type Clan struct {
	Id              primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Name            string             `json:"name" bson:"name"`
	Abbr            string             `json:"abbr" bson:"abbr"`
	LeaderSteamId   int                `json:"leaderSteamId" bson:"leaderSteamId"`
	Created         time.Time          `json:"created" bson:"created"`
	Balance         int                `json:"balance" bson:"balance"`
	Tax             int                `json:"tax" bson:"tax"`
	Level           int                `json:"level" bson:"level"`
	Experience      int                `json:"experience" bson:"experience"`
	MembersSteamIds []int              `json:"membersSteamIds" bson:"membersSteamIds"`
}
