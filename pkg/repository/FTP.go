package repository

import (
	"github.com/jlaffaye/ftp"
	"os"
)

func NewFtpConnect() (*ftp.ServerConn, error) {
	conn, err := ftp.Dial(os.Getenv("FTP_URI"))
	if err != nil {
		return nil, err
	}
	err = conn.Login(os.Getenv("FTP_LOGIN"), os.Getenv("FTP_PASSWORD"))
	if err != nil {
		return nil, err
	}

	return conn, err

}
