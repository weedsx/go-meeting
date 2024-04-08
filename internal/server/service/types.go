package service

import "time"

type UserLoginRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password"`
}

type MeetingListRequest struct {
	Page    int    `json:"page" form:"page"`
	Size    int    `json:"size" form:"size"`
	Keyword string `json:"keyword" form:"keyword"`
}

type MeetingListReply struct {
	Identity string    `json:"identity"`
	Name     string    `json:"name,omitempty"`
	BeginAt  time.Time `json:"begin_at"`
	EndAt    time.Time `json:"end_at"`
}
