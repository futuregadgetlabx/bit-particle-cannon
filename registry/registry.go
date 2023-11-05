package registry

import (
	"errors"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

var Users = make(map[string]string)
var dataFile string

const (
	DataFileRelease = "/data/users.data"
	DataFileDev     = "users.data"
)

func Load() {
	if os.Getenv("ENV") == "dev" {
		dataFile = DataFileDev
	} else {
		dataFile = DataFileRelease
	}
	_, err := os.Stat(dataFile)
	if err != nil && os.IsNotExist(err) {
		_, err := os.Create(dataFile)
		if err != nil {
			panic(err)
		}
	}
	f, err := os.ReadFile(dataFile)
	if err != nil {
		return
	}

	lines := strings.Split(string(f), "\n")
	for _, l := range lines {
		if !strings.ContainsRune(l, ':') {
			continue
		}
		kv := strings.Split(l, ":")
		Users[kv[0]] = kv[1]
	}

	logrus.Info("User data loaded.")
}

func Add(larkID, creds string) error {
	if !strings.HasPrefix(creds, "LEETCODE_SESSION") {
		return errors.New("invalid creds :" + creds)
	}
	Users[larkID] = creds

	// save to file
	f, err := os.OpenFile(dataFile, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logrus.WithError(err).Error("open data file error.")
		return err
	}
	defer f.Close()
	// refresh all Users
	var data []string
	for k, v := range Users {
		data = append(data, k+":"+v)
	}
	_, err = f.WriteString(strings.Join(data, "\n"))
	if err != nil {
		logrus.WithError(err).Error("write data file error.")
		return err
	}
	return nil
}
