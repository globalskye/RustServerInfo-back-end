package repository

import (
	"github.com/jlaffaye/ftp"
	"io/ioutil"
)

func GetBytesFromFile(path string, ftp *ftp.ServerConn) ([]byte, error) {
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
