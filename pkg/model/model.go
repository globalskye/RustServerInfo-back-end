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
	Online           float32            `json:"online" bson:"online"`
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
	Members         []Player           `json:"members" bson:"members"`
}

type VkPost struct {
	Response struct {
		Count int `json:"count"`
		Items []struct {
			Id          int    `json:"id"`
			FromId      int    `json:"from_id"`
			OwnerId     int    `json:"owner_id"`
			Date        int    `json:"date"`
			PostponedId int    `json:"postponed_id,omitempty"`
			MarkedAsAds int    `json:"marked_as_ads"`
			CanDelete   int    `json:"can_delete"`
			IsFavorite  bool   `json:"is_favorite"`
			PostType    string `json:"post_type"`
			Text        string `json:"text"`
			CanEdit     int    `json:"can_edit,omitempty"`
			CreatedBy   int    `json:"created_by,omitempty"`
			CanPin      int    `json:"can_pin"`
			IsPinned    int    `json:"is_pinned,omitempty"`
			Attachments []struct {
				Type  string `json:"type"`
				Photo struct {
					AlbumId   int    `json:"album_id"`
					Date      int    `json:"date"`
					Id        int    `json:"id"`
					OwnerId   int    `json:"owner_id"`
					AccessKey string `json:"access_key"`
					PostId    int    `json:"post_id"`
					Sizes     []struct {
						Height int    `json:"height"`
						Type   string `json:"type"`
						Width  int    `json:"width"`
						Url    string `json:"url"`
					} `json:"sizes"`
					Text    string `json:"text"`
					UserId  int    `json:"User_id"`
					HasTags bool   `json:"has_tags"`
				} `json:"photo,omitempty"`
				Video struct {
					AccessKey     string `json:"access_key"`
					CanComment    int    `json:"can_comment"`
					CanEdit       int    `json:"can_edit"`
					CanLike       int    `json:"can_like"`
					CanRepost     int    `json:"can_repost"`
					CanSubscribe  int    `json:"can_subscribe"`
					CanAddToFaves int    `json:"can_add_to_faves"`
					CanAdd        int    `json:"can_add"`
					CanAttachLink int    `json:"can_attach_link"`
					Comments      int    `json:"comments"`
					Date          int    `json:"date"`
					Description   string `json:"description"`
					Duration      int    `json:"duration"`
					Image         []struct {
						Url         string `json:"url"`
						Width       int    `json:"width"`
						Height      int    `json:"height"`
						WithPadding int    `json:"with_padding"`
					} `json:"image"`
					Id         int    `json:"id"`
					OwnerId    int    `json:"owner_id"`
					Title      string `json:"title"`
					IsFavorite bool   `json:"is_favorite"`
					TrackCode  string `json:"track_code"`
					Type       string `json:"type"`
					Views      int    `json:"views"`
					LocalViews int    `json:"local_views"`
					Platform   string `json:"platform"`
				} `json:"video,omitempty"`
			} `json:"attachments"`
			PostSource struct {
				Type string `json:"type"`
			} `json:"post_source"`
			Comments struct {
				CanPost       int  `json:"can_post"`
				CanClose      int  `json:"can_close"`
				Count         int  `json:"count"`
				GroupsCanPost bool `json:"groups_can_post"`
			} `json:"comments"`
			Likes struct {
				CanLike    int `json:"can_like"`
				Count      int `json:"count"`
				UserLikes  int `json:"user_likes"`
				CanPublish int `json:"can_publish"`
			} `json:"likes"`
			Reposts struct {
				Count        int `json:"count"`
				WallCount    int `json:"wall_count"`
				MailCount    int `json:"mail_count"`
				UserReposted int `json:"user_reposted"`
			} `json:"reposts"`
			Views struct {
				Count int `json:"count"`
			} `json:"views"`
			Donut struct {
				IsDonut bool `json:"is_donut"`
			} `json:"donut"`
			ShortTextRate float64 `json:"short_text_rate"`
			Hash          string  `json:"hash"`
			Edited        int     `json:"edited,omitempty"`
		} `json:"items"`
	} `json:"response"`
}
type User struct {
	Id       primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Username string             `json:"username" binding:"required"`
	Balance  int                `bson:"balance" json:"balance"`
	Password string             `json:"password" binding:"required"`
	Role     string             `json:"role" `
}
