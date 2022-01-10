package models

import "time"

type Logs struct {
	id       string `json:"id;primaryKey;autoIncrement"`
	Method   string `json:"method"`
	Params   string `json:"params"`
	accessAt time.Time
}

func NewLogs(
	new_method string,
	new_params string,
	new_time time.Time,
) Logs {
	return Logs{
		Method:   new_method,
		Params:   new_params,
		accessAt: time.Now().UTC(),
	}
}
