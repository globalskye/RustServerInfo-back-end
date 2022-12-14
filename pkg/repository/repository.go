package repository

import (
	"github.com/jlaffaye/ftp"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	UserI
}

func NewRepository(r *RepoTools) *Repository {
	return &Repository{
		UserI: NewUserRepository(r),
	}
}

type UserI interface {
	GetAllUsersFiles() (map[string][]byte, error)
	GetAllClansFiles() (map[string][]byte, error)
}

type RepoTools struct {
	ftp   *ftp.ServerConn
	mongo *mongo.Database
}

func NewRepoTools() *RepoTools {
	ftpConn, err := NewFtpConnect()
	if err != nil {
		logrus.Fatalf("failed to initialize ftp: %s", err.Error())
	}
	mongoConn, err := NewMongoConnect()
	if err != nil {
		logrus.Fatalf("failed to initialize mongoDb: %s", err.Error())
	}
	return &RepoTools{
		ftp:   ftpConn,
		mongo: mongoConn,
	}
}
