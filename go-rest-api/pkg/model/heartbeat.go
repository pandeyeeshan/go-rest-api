package model

type HeartbeatResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}