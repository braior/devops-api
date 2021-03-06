package common

import (
	"github.com/chanyipiaomiao/hltool"
)

// BackupBolDB 备份数据库文件
func BackupBolDB(filepath string) error {
	btb, err := hltool.NewBoltDB(DBPath, "token")
	if err != nil {
		return err
	}

	err = btb.Backup(filepath)
	if err != nil {
		return err
	}
	return nil
}
