package common

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/chanyipiaomiao/phonedata"
)

// QueryPhone 获取电话号码归属地信息
func QueryPhone(phone string) (map[string]string, error) {
	// 通过beego解析phonedat文件路径
	dbPath := beego.AppConfig.String("phone::dbpath")
	if dbPath == "" {
		return nil, fmt.Errorf("not found phone dat file")
	}

	p, err := phonedata.NewPhoneQuery(dbPath)
	if err != nil {
		return nil, err
	}

	m, err := p.Query(phone)
	if err != nil {
		return nil, err
	}

	return m, nil
}
