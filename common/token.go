package common

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/chanyipiaomiao/hltool"
)

const (
	// 存放token的表名
	tokenTableName = "token"
)

// Token 结构体
type Token struct {
	TokenDb    *hltool.BoltDB
	SignString string
}

// NewToken 返回Token对象
func NewToken() (*Token, error) {
	tokenDb, err := hltool.NewBoltDB(DBPath, tokenTableName)
	if err != nil {
		return nil, err
	}

	signString := beego.AppConfig.String("security::jwtokenSignString")
	if signString == "" {
		return nil, fmt.Errorf("warning: in conf file jwtokenSignString must not null")
	}

	return &Token{
		TokenDb:    tokenDb,
		SignString: signString,
	}, nil
}

// GetToken 根据name获取token
func (t *Token) GetToken(name string) (map[string][]byte, error) {
	result, err := t.TokenDb.Get([]string{name})
	if err != nil {
		return nil, fmt.Errorf("get token < %s > error: %s", name, err)
	}
	return result, nil
}

// IsExistToken token 是否存在
// name token的名称
func (t *Token) IsExistToken(name string) (bool, error) {
	result, err := t.GetToken(name)
	if err != nil {
		return false, err
	}

	if _, ok := result[name]; !ok {
		return false, nil
	}

	if string(result[name]) != "" {
		return true, fmt.Errorf("exist < %s > token", name)
	}

	return false, nil
}

// IsTokenValid token是否有效
func (t *Token) IsTokenValid(token string) (bool, error) {
	jwt := hltool.NewJWToken(t.SignString)
	parseToken, err := jwt.ParseJWToken(token)
	if err != nil {
		return false, err
	}

	tokenName := parseToken["name"].(string)
	dbToken, err := t.GetToken(tokenName)
	if err != nil {
		return false, err
	}

	if _, ok := dbToken[tokenName]; !ok {
		return false, fmt.Errorf("token is not exist")
	}

	if string(dbToken[tokenName]) == token {
		return true, nil
	}

	return false, fmt.Errorf("token is not valid")
}
