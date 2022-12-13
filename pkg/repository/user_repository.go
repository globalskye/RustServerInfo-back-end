package repository

import (
	"github.com/jlaffaye/ftp"
	"io/ioutil"
)

type UserRepository struct {
	tool *RepoTools
}

func NewUserRepository(tool *RepoTools) *UserRepository {
	return &UserRepository{tool: tool}
}
func (u UserRepository) GetAllUsersFiles() (map[string][]byte, error) {
	result := make(map[string][]byte, 4)
	rustUsersBytes, err := getBytesFromFile("rust_users.txt", u.tool.ftp)
	if err != nil {
		return nil, err
	}
	rustTopFarmBytes, err := getBytesFromFile("oxide/data/TopFarmer.json", u.tool.ftp)
	if err != nil {
		return nil, err
	}
	rustTopOnlineBytes, err := getBytesFromFile("oxide/data/TopOnline.json", u.tool.ftp)
	if err != nil {
		return nil, err
	}
	rustTopRaidBytes, err := getBytesFromFile("oxide/data/TopRaiders.json", u.tool.ftp)
	if err != nil {
		return nil, err
	}
	result["users"] = rustUsersBytes
	result["topfarm"] = rustTopFarmBytes
	result["toponline"] = rustTopOnlineBytes
	result["topraid"] = rustTopRaidBytes
	return result, nil
}
func getBytesFromFile(path string, ftp *ftp.ServerConn) ([]byte, error) {
	r, err := ftp.Retr(path)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return buf, err
}
