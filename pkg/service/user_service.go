package service

import (
	"encoding/json"
	"errors"
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/model"
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/repository"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

type UserService struct {
	repo repository.UserI
}

func NewUserService(repo repository.UserI) *UserService {
	return &UserService{repo: repo}
}

type topFarm struct {
	SteamId int `json:"steamId"`
	Wood    int `json:"Дерева"`
	Metal   int `json:"Метала"`
	Sulfur  int `json:"Серы"`
	Leather int `json:"Кожы"`
	Cloth   int `json:"Ткани"`
	Fat     int `json:"Жира"`
}

const TimeLayout = "01/02/2006 15:04:05"

func (u UserService) GetOnline() ([]string, error) {
	resp, err := http.Get("https://rage.hostfun.ru/players.php")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode > 400 {
		return nil, errors.New("Cannot get online")
	}

	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Cannot read response.body from https://rage.hostfun.ru/players.php")
	}

	s := string(bytes)
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "\r\n", "")
	arr := strings.Split(s, "<divclass='collapse'id='server_62_122_214_162_27019'><divclass='cardcard-body'style='display:inline-flex;'><p>")
	subArr := strings.Split(arr[1], "</p>")
	users := strings.Split(subArr[0], ",")

	return users, err
}
func (u UserService) GetAllUsers() ([]model.User, error) {
	counter, err := u.repo.GetAllUsersFiles()
	if err != nil {
		return nil, err
	}
	users := unmarshalUserBytes(counter["users"])
	sort.Slice(users, func(i, j int) bool {
		return users[i].SteamId < users[j].SteamId
	})
	farm := unmarshalTopFarmBytes(counter["topfarm"])
	sort.Slice(farm, func(i, j int) bool {
		return farm[i].SteamId < farm[j].SteamId
	})
	online := unmarshalTopOnlineBytes(counter["toponline"])
	raid := unmarshalTopRaidBytes(counter["topraid"])
	for i, u := range users {
		for _, f := range farm {
			if u.SteamId == f.SteamId {
				users[i].Farm = model.UserTopFarm{
					SteamId: f.SteamId,
					Cloth:   f.Cloth,
					Fat:     f.Fat,
					Leather: f.Leather,
					Metal:   f.Metal,
					Sulfur:  f.Sulfur,
					Wood:    f.Wood,
				}
			}
		}
		users[i].Online = online[u.SteamId]
		users[i].Raid = raid[u.SteamId]
	}
	return users, err
}
func unmarshalTopFarmBytes(bytes []byte) []topFarm {
	var topFarm []topFarm
	json.Unmarshal(bytes, &topFarm)
	return topFarm
}
func unmarshalTopOnlineBytes(bytes []byte) map[int]float32 {
	var topOnline map[int]float32
	json.Unmarshal(bytes, &topOnline)

	return topOnline
}
func unmarshalTopRaidBytes(bytes []byte) map[int]float32 {
	var topRaid map[int]float32
	json.Unmarshal(bytes, &topRaid)

	return topRaid
}

func unmarshalUserBytes(bytes []byte) []model.User {
	str := string(bytes)
	arr := strings.Split(str, "\r\n\r\n")

	var users []model.User

	for i := 1; i < len(arr)-1; i++ {
		subArr := strings.Split(arr[i], "\r\n")
		var user model.User
		for j := 0; j < len(subArr); j++ {
			if j == 0 {
				subArr[j] = strings.ReplaceAll(subArr[j], "[", "")
				subArr[j] = strings.ReplaceAll(subArr[j], "]", "")
				s, _ := strconv.Atoi(subArr[j])
				user.SteamId = s
				continue
			}
			if strings.Contains(subArr[j], "USERNAME=") {
				subArr[j] = strings.ReplaceAll(subArr[j], "USERNAME=", "")
				user.Name = subArr[j]
				continue
			}
			if strings.Contains(subArr[j], "HWID=") {
				subArr[j] = strings.ReplaceAll(subArr[j], "HWID=", "")
				user.Hwid = subArr[j]
				continue
			}
			if strings.Contains(subArr[j], "RANK=") {
				subArr[j] = strings.ReplaceAll(subArr[j], "RANK=", "")
				s, _ := strconv.Atoi(subArr[j])
				user.Rank = s
				continue
			}
			if strings.Contains(subArr[j], "COUNTDOWN=kit.") {
				subArr[j] = strings.ReplaceAll(subArr[j], "COUNTDOWN=kit.", "")
				a := strings.Split(subArr[j], ",")
				if a[1] == "0" {
					user.Kits = append(user.Kits, model.KitInfo{
						Name:       a[0],
						Disposable: true,
					})
					continue
				}
				t, _ := time.Parse(TimeLayout, a[1])
				user.Kits = append(user.Kits, model.KitInfo{
					Name:      a[0],
					Countdown: t,
				})
			}
			if strings.Contains(subArr[j], "LASTCONNECTDATE=") {
				subArr[j] = strings.ReplaceAll(subArr[j], "LASTCONNECTDATE=", "")
				t, _ := time.Parse(TimeLayout, subArr[j])
				user.LastConnectTime = t
				continue
			}
			if strings.Contains(subArr[j], "FIRSTCONNECTDATE=") {
				subArr[j] = strings.ReplaceAll(subArr[j], "FIRSTCONNECTDATE=", "")
				t, _ := time.Parse(TimeLayout, subArr[j])
				user.FirstConnectTime = t
				continue
			}
			if strings.Contains(subArr[j], "BALANCE=") {
				subArr[j] = strings.ReplaceAll(subArr[j], "BALANCE=", "")
				s, _ := strconv.Atoi(subArr[j])
				user.Balance = s
				continue
			}
			if strings.Contains(subArr[j], "KILLEDPLAYERS=") {
				subArr[j] = strings.ReplaceAll(subArr[j], "KILLEDPLAYERS=", "")
				s, _ := strconv.Atoi(subArr[j])
				user.KilledPlayers = s
				continue
			}
			if strings.Contains(subArr[j], "KILLEDMUTANTS=") {
				subArr[j] = strings.ReplaceAll(subArr[j], "KILLEDMUTANTS=", "")
				s, _ := strconv.Atoi(subArr[j])
				user.KilledMutants = s
				continue
			}
			if strings.Contains(subArr[j], "KILLEDANIMALS=") {
				subArr[j] = strings.ReplaceAll(subArr[j], "KILLEDANIMALS=", "")
				s, _ := strconv.Atoi(subArr[j])
				user.KilledAnimals = s
				continue
			}
			if strings.Contains(subArr[j], "DEATHS=") {
				subArr[j] = strings.ReplaceAll(subArr[j], "DEATHS=", "")
				s, _ := strconv.Atoi(subArr[j])
				user.Deaths = s
				continue
			}
		}
		users = append(users, user)
	}
	return users
}
