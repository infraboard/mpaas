package task

import (
	"fmt"
	"time"
)

func NewCallbackStatus(describe string) *CallbackStatus {
	return &CallbackStatus{
		Description: describe,
	}
}

func (h *CallbackStatus) StartSend() {
	h.StartAt = time.Now().UnixMilli()
}

func (h *CallbackStatus) SendFailed(format string, a ...interface{}) {
	if h.StartAt != 0 {
		h.Cost = time.Now().UnixMilli() - h.StartAt
	}
	h.Message = fmt.Sprintf(format, a...)
}

func (h *CallbackStatus) SendSuccess(message string) {
	if h.StartAt != 0 {
		h.Cost = time.Now().UnixMilli() - h.StartAt
	}
	h.Success = true
	h.Message = message
}
