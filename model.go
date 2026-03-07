package restlogger

import "time"

type LogRow struct {
	UUID    string                     `db:"uuid" json:"uuid,omitempty"`
	AppName string                     `db:"app_name" json:"app_name"`
	StaticData map[string]interface{}  `db:"static_data" json:"static_data"`
	DynamicData map[string]interface{} `db:"dynamic_data" json:"dynamic_data"`
	UpdDttm time.Time                  `db:"upd_dttm" json:"upd_dttm,omitempty"`
	UpdCnt  int64                      `db:"upd_cnt"  json:"upd_cnt,omitempty"`
	ReadFlg bool                       `db:"read_flg" json:"read_flg,omitempty"`
}