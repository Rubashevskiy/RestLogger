package restlogger

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Heandler struct {
	repo *pgxpool.Pool
}

func NewHeandler(connString string) (*Heandler, error) {
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}
	defer pool.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5* time.Second)
	defer cancel()

	if err := pool.Ping(ctx); err != nil {
		return nil, err
	} else {
		return &Heandler{repo: pool}, nil
	}
}

func (h *Heandler) UpsertLog(w http.ResponseWriter, r *http.Request) {
    var log_data LogRow
    if err := json.NewDecoder(r.Body).Decode(&log_data); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    upsert := `INSERT INTO logs.app_logs(uuid, app_name, static_data, dynamic_data, upd_dttm, upd_cnt, read_flg)
			   VALUES (@uuid, @app_name, @static_data, @dynamic_data, Now(), 1, false)
			   ON CONFLICT (app_name, static_data) 
			   DO UPDATE SET 
    		       dynamic_data = EXCLUDED.dynamic_data,
    			   upd_dttm = EXCLUDED.upd_dttm,
    			   upd_cnt = logs.app_logs.upd_cnt + 1
    			   read_flg = EXCLUDED.read_flg;
	`
	args := pgx.NamedArgs{
		"uuid"        : uuid.New().String(),
		"app_name"    : log_data.AppName, 
		"static_data" : log_data.StaticData,
		"dynamic_data": log_data.DynamicData,
	}
	if _, err := h.repo.Exec(context.Background(), upsert, args); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
        return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Heandler) Status() (bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 5* time.Second)
	defer cancel()

	if err := h.repo.Ping(ctx); err != nil {
		return false
	} else {
		return true
	}	
}