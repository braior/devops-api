package cmd

import "fmt"

// GetUsers 根据name获取token
func (t *Token) GetUsers() (map[string][]byte, error) {
	result, err := t.TokenDB.GetAll()
	if err != nil {
		return nil, fmt.Errorf("get users error: %s", err)
	}
	return result, nil
}
