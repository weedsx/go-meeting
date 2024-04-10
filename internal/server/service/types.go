package service

import (
	"time"
)

// swagger:parameters UserLoginRequest
type UserLoginRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password"`
}

// swagger:parameters MeetingListRequest
type MeetingListRequest struct {
	Page    int    `json:"page" form:"page"`
	Size    int    `json:"size" form:"size"`
	Keyword string `json:"keyword" form:"keyword"`
}

// swagger:parameters MeetingListReply
type MeetingListReply struct {
	Identity string    `json:"identity"`
	Name     string    `json:"name,omitempty"`
	BeginAt  time.Time `json:"begin_at"`
	EndAt    time.Time `json:"end_at"`
}

// swagger:parameters MeetingCreateRequest
type MeetingCreateRequest struct {
	Name    string `json:"name,omitempty"`
	BeginAt int64  `json:"begin_at"`
	EndAt   int64  `json:"end_at"`
}

// swagger:parameters MeetingEditRequest
type MeetingEditRequest struct {
	Identity string `json:"identity"`
	*MeetingCreateRequest
}

// swagger:parameters WsP2PConnectionRequest
type WsP2PConnectionRequest struct {
	RoomIdentity string `json:"room_identity" uri:"room_identity"`
	UserIdentity string `json:"user_identity" uri:"user_identity"`
}
