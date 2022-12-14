package repository

type ClanRepository struct {
	tool *RepoTools
}

func NewClanRepository(tool *RepoTools) *ClanRepository {
	return &ClanRepository{tool: tool}
}
func (u ClanRepository) GetAllClansFiles() (map[string][]byte, error) {
	result := make(map[string][]byte, 1)
	rustClansBytes, err := GetBytesFromFile("rust_clans.txt", u.tool.ftp)
	if err != nil {
		return nil, err
	}
	result["clans"] = rustClansBytes
	return result, err
}
