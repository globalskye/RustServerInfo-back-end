package service

import (
	"fmt"
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/model"
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/repository"
	"strconv"
	"strings"
)

type UserService struct {
	repo repository.UserI
}

func NewUserService(repo repository.UserI) *UserService {
	return &UserService{repo: repo}
}

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

	for i := 1; i < len(arr); i++ {
		subArr := strings.Split(arr[i], "\r\n")
		for j := 0; i < len(subArr); j++ {
			var user model.User
			if j == 0 {
				subArr[j] = strings.ReplaceAll(subArr[j], "[", "")
				subArr[j] = strings.ReplaceAll(subArr[j], "]", "")
				s, _ := strconv.Atoi(subArr[j])
				user.SteamId = s
				users = append(users, user)
				break
			}

		}

	}
	fmt.Println(users[0])

	//_ := "TITLE=RustExtended.Core\nVERSION=1.0.2.5\nTIME=327466816\n\n[76561197965680350]\nUSERNAME=HATO\nPASSWORD=\nCOMMENTS=\nHWID=C1DD8D9D-7"
	//fmt.Println(datas[1])
	return nil
}
