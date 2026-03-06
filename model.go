package restlogger

import "time"

type LogRow struct {
	UUID    string                     `db:"uuid"`
	AppName string                     `db:"app_name" json:"app_name"`
	StaticData map[string]interface{}  `db:"static_data" json:"static_data"`
	DinamicData map[string]interface{} `db:"dinamic_data" json:"dinamic_data"`
	UpdDttm time.Time                  `db:"upd_dttm"`
	UpdCnt  int64                      `db:"upd_cnt"`
	ReadFlg bool                       `db:"read_flg"`
}