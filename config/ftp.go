package config

import (
	"errors"
	"fmt"
	"github.com/jlaffaye/ftp"
	log "github.com/sirupsen/logrus"
	"time"
)

func SetupConnectionFTP(username string, password string, host string, port int) (*ftp.ServerConn, error) {
	c, err := ftp.Dial(fmt.Sprintf("%s:%v", host, port), ftp.DialWithTimeout(5*time.Second))

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Cannot connect FTP %s", err))
	}

	err = c.Login(username, password)
	if err != nil {
		return nil, err
	}

	log.Info("Connected to FTP ", fmt.Sprintf("%s:%v", host, port))

	return c, nil
}