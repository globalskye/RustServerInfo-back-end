package service

import (
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/model"
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/repository"
	"strconv"
	"strings"
	"time"
)

type ClanService struct {
	repo repository.ClanI
}

func NewClanService(repo repository.ClanI) *ClanService {
	return &ClanService{repo: repo}
}

func (u ClanService) GetAllClans() ([]model.Clan, error) {
	counter, err := u.repo.GetAllClansFiles()
	if err != nil {
		return nil, err
	}
	clans := unmarshalClanBytes(counter["clans"])
	return clans, err
}
func unmarshalClanBytes(bytes []byte) []model.Clan {
	str := string(bytes)
	arr := strings.Split(str, "\r\n\r\n")

	var clans []model.Clan
	for i := 1; i < len(arr); i++ {
		subArr := strings.Split(arr[i], "\r\n")
		var clan model.Clan
		for j := 1; j < len(subArr); j++ {
			if strings.Contains(subArr[j], "NAME=") {
				subArr[j] = strings.ReplaceAll(subArr[j], "NAME=", "")
				clan.Name = subArr[j]
				continue
			}
			if strings.Contains(subArr[j], "ABBREV=") {
				subArr[j] = strings.ReplaceAll(subArr[j], "ABBREV=", "")
				clan.Abbr = subArr[j]
				continue
			}
			if strings.Contains(subArr[j], "LEADER=") {
				subArr[j] = strings.ReplaceAll(subArr[j], "LEADER=", "")
				s, _ := strconv.Atoi(subArr[j])
				clan.LeaderSteamId = s
				continue
			}
			if strings.Contains(subArr[j], "CREATED=") {
				subArr[j] = strings.ReplaceAll(subArr[j], "CREATED=", "")
				t, _ := time.Parse(TimeLayout, subArr[j])
				clan.Created = t
				continue
			}
			if strings.Contains(subArr[j], "BALANCE=") {
				subArr[j] = strings.ReplaceAll(subArr[j], "BALANCE=", "")
				s, _ := strconv.Atoi(subArr[j])
				clan.Balance = s
				continue
			}
			if strings.Contains(subArr[j], "TAX=") {
				subArr[j] = strings.ReplaceAll(subArr[j], "TAX=", "")
				s, _ := strconv.Atoi(subArr[j])
				clan.Tax = s
				continue
			}
			if strings.Contains(subArr[j], "LEVEL=") {
				subArr[j] = strings.ReplaceAll(subArr[j], "LEVEL=", "")
				s, _ := strconv.Atoi(subArr[j])
				clan.Level = s
				continue
			}
			if strings.Contains(subArr[j], "EXPERIENCE=") {
				subArr[j] = strings.ReplaceAll(subArr[j], "EXPERIENCE=", "")
				s, _ := strconv.Atoi(subArr[j])
				clan.Experience = s
				continue
			}
			if strings.Contains(subArr[j], "MEMBER=") {
				subArr[j] = strings.ReplaceAll(subArr[j], "MEMBER=", "")
				a := strings.Split(subArr[j], ",")
				s, _ := strconv.Atoi(a[0])
				clan.MembersSteamIds = append(clan.MembersSteamIds, s)
				continue
			}

		}
		clans = append(clans, clan)
	}
	return clans
}
