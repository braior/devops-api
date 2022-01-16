package controllers

import (
	"github.com/braior/brtool"
	"github.com/sirupsen/logrus"
)

// Get md5 ...
func (m *MD5Controller) Get() {
	rawString := m.GetString("rawStr")
	rawStringMD5 := brtool.GetMD5(rawString)
	data := map[string]string{
		"rawString":    rawString,
		"rawStringMD5": rawStringMD5,
	}
	m.Json("Get String MD5", "", 0, logrus.InfoLevel, data, true)
}
