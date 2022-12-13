package repository

import "github.com/jlaffaye/ftp"

type Repository struct {
}

func NewRepository(ftp *ftp.ServerConn) *Repository {
	return &Repository{}
}
