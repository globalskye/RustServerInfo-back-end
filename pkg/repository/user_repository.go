package repository

type UserRepository struct {
	tool *RepoTools
}

func NewUserRepository(tool *RepoTools) *UserRepository {
	return &UserRepository{tool: tool}
}
func (u UserRepository) GetAllUsersFiles() (map[string][]byte, error) {
	result := make(map[string][]byte, 4)
	rustUsersBytes, err := GetBytesFromFile("rust_users.txt", u.tool.ftp)
	if err != nil {
		return nil, err
	}
	rustTopFarmBytes, err := GetBytesFromFile("oxide/data/TopFarmer.json", u.tool.ftp)
	if err != nil {
		return nil, err
	}
	rustTopOnlineBytes, err := GetBytesFromFile("oxide/data/TopOnline.json", u.tool.ftp)
	if err != nil {
		return nil, err
	}
	rustTopRaidBytes, err := GetBytesFromFile("oxide/data/TopRaiders.json", u.tool.ftp)
	if err != nil {
		return nil, err
	}
	result["users"] = rustUsersBytes
	result["topfarm"] = rustTopFarmBytes
	result["toponline"] = rustTopOnlineBytes
	result["topraid"] = rustTopRaidBytes
	return result, nil
}
