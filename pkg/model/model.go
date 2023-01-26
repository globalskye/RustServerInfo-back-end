package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Player struct {
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
	Farm             PlayerTopFarm      `json:"farm" bson:"farm"`
	Online           int                `json:"online" bson:"online"`
	Raid             float32            `json:"raid" bson:"raid"`
}

type PlayerTopFarm struct {
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
	Id        primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Count     int                `bson:"count" json:"count"`
	Nicknames []string           `json:"nicknames" bson:"nicknames"`
}
type Clan struct {
	Id              primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Name            string             `json:"name" bson:"name"`
	Abbr            string             `json:"abbr" bson:"abbr"`
	LeaderSteamId   int                `json:"leaderSteamId" bson:"leaderSteamId"`
	Leader          Player             `json:"leader" bson:"leader"`
	Created         time.Time          `json:"created" bson:"created"`
	Balance         int                `json:"balance" bson:"balance"`
	Tax             int                `json:"tax" bson:"tax"`
	Level           int                `json:"level" bson:"level"`
	Experience      int                `json:"experience" bson:"experience"`
	MembersSteamIds []int              `json:"membersSteamIds" bson:"membersSteamIds"`

	Members []Player `json:"members" bson:"members"`
}
type VkPost struct {
	Items []VkPostItem `bson:"items" json:"items"`
}
type VkPostItem struct {
	Id        int    `json:"id" bson:"id"`
	OwnerId   int    `json:"owner_id" bson:"ownerId"`
	Date      int    `json:"date" bson:"date"`
	Text      string `json:"text" bson:"text"`
	ImageLink string `bson:"imageLink" json:"imageLink"`
	VideoLink string `bson:"videoLink" json:"videoLink"`
}

type User struct {
	Id       primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Username string             `json:"username" binding:"required"`
	Balance  int                `bson:"balance" json:"balance"`
	Password string             `json:"password" binding:"required"`
	Role     string             `json:"role" bson:"role"`
	Cart     []CartItem         `json:"cart" json:"cart"`
	Stock    []StockItem        `json:"stock" json:"stock"`
}
type CartItem struct {
	Item  DonatItem `json:"item"`
	Count int       `json:"count"`
}
type StockItem struct {
	Item  DonatItem `json:"item"`
	Count int       `json:"count"`
}

type DonatItem struct {
	Id          primitive.ObjectID `json:"id" bson:"id"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Category    string             `json:"category" bson:"category"`
	ImageUrl    string             `json:"imageUrl" bson:"imageUrl"`
	Price       int                `json:"price" bson:"price"`
	Rank        int                `json:"rank" bson:"rank"`
	Attachments []DonatAttachments `json:"attachments" bson:"attachments"`
}
type DonatAttachments struct {
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	ImageUrl    string `json:"imageUrl" bson:"imageUrl"`
}
