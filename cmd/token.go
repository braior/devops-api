package cmd

import (
	"fmt"
	"time"

	"github.com/braior/brtool"
	"github.com/spf13/viper"
)

const (
	// 定义存放token的表明
	tokenTableName = "token"
)

// Token struct
type Token struct {
	TokenDB    *brtool.BoltDB
	SignString string
}

// NewToken 返回Token对象
func NewToken() (*Token, error) {
	tokenDB, err := brtool.NewBoltDB(viper.GetString("db.boltDB.dbPath"), tokenTableName)

	if err != nil {
		return nil, err
	}

	signString := viper.GetString("security.jwtokenSignString")
	if signString == "" {
		return nil, fmt.Errorf("warning: in conf file jwtokenSignString must not null")
	}
	return &Token{TokenDB: tokenDB, SignString: signString}, nil
}

// GetToken 根据name获取token
func (t *Token) GetToken(name string) (map[string][]byte, error) {
	result, err := t.TokenDB.Get([]string{name})
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
	jwt := brtool.NewJWToken(t.SignString)
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

	return false, fmt.Errorf("token is not vaild")
}

// IsRootToken 是否是root token,root token 不能被用来请求
func (t *Token) IsRootToken(token string) (bool, error) {
	jwt := brtool.NewJWToken(t.SignString)
	parseToken, err := jwt.ParseJWToken(token)
	if err != nil {
		return false, err
	}
	tokenName := parseToken["name"].(string)
	return tokenName == "root", nil
}

// DeleteToken 删除Token
// name token名称
func (t *Token) DeleteToken(rootToken, name string) error {
	if name == "root" {
		return fmt.Errorf("can't delete root token")
	}

	if rootToken == "" {
		return fmt.Errorf("need root token")
	}

	if ok, err := t.IsTokenValid(rootToken); !ok {
		return err
	}

	r, _ := t.IsExistToken(name)
	if r {
		err := t.TokenDB.Delete([]string{name})
		if err != nil {
			return fmt.Errorf("delete token < %s > error: %s", name, err)
		}
		fmt.Printf("delete token <%s> ok.\n", name)
		return nil
	}
	return fmt.Errorf("token < %s > not exist", name)
}

// AddToken 生成一个root token 用于管理其他的token
// rootToken root token 创建其他token 需要root token
// name token的名称: root token名为: root , 其他token: 指定的名称
func (t *Token) AddToken(rootToken, name string) error {
	if name != "root" && rootToken == "" {
		return fmt.Errorf("error: need root token, please use --root-token appoint root token")
	}
	if rootToken != "" {
		if ok, err := t.IsTokenValid(rootToken); !ok {
			return err
		}
	}

	tokenValue := map[string]interface{}{
		"name": name,
		"updateTime": func() int64 {
			return time.Now().Unix()
		}(),
	}

	jwt := brtool.NewJWToken(t.SignString)
	token, err := jwt.GenJWToken(tokenValue)
	if err != nil {
		return err
	}

	t.TokenDB.Set(map[string][]byte{
		name: []byte(token),
	})

	fmt.Printf("warning: For < %s > token only shows once, keep in mind!!! \n", name)
	fmt.Printf("\t %s \n", token)
	return nil
}

// AddRootToken 创建一个root token
func (t *Token) AddRootToken() error {
	r, err := t.IsExistToken("root")
	if r {
		return fmt.Errorf("%s, you can use --refresh-root-token refresh root token", err)
	}
	return t.AddToken("", "root")
}

// ForceRefresh: 是否强制刷新 root token
func (t *Token) ForceRefresh() error {
	return t.AddToken("", "root")

}
