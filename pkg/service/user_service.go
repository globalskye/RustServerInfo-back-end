package service

import (
	"fmt"
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/model"
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/repository"
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

const TimeLayout = "01/02/2006 15:04:05"

func (u UserService) GetAllUsers() ([]model.User, error) {
	counter, err := u.repo.GetAllUsersFiles()
	if err != nil {
		return nil, err
	}
	unmarshalUserBytes(counter["users"])

	return nil, nil
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
				t, _ := time.Parse(TimeLayout, subArr[j])
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
	fmt.Println(users[0])

	//_ := "TITLE=RustExtended.Core\nVERSION=1.0.2.5\nTIME=327466816\n\n[76561197965680350]\nUSERNAME=HATO\nPASSWORD=\nCOMMENTS=\nHWID=C1DD8D9D-7"
	//fmt.Println(datas[1])
	return nil
}
