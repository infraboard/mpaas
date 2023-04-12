package task

import (
	"fmt"
	"strings"
	"time"
)

func NewCallbackStatus(describe string) *CallbackStatus {
	return &CallbackStatus{
		Description: describe,
		Events:      []*Event{},
	}
}

func (h *CallbackStatus) StartSend() {
	h.StartAt = time.Now().UnixMilli()
}

func (h *CallbackStatus) MakeStatusUseEvent() {
	msg := h.ErrorEventMessage()
	if msg != "" {
		h.SendFailed(msg)
	} else {
		h.SendSuccess("")
	}
}

func (h *CallbackStatus) ErrorEventMessage() string {
	items := h.ErrorEvent()
	if len(items) == 0 {
		return ""
	}

	errors := []string{}
	for i := range items {
		item := items[i]
		errors = append(errors, item.Message)
	}

	return strings.Join(errors, ",")
}

func (h *CallbackStatus) ErrorEvent() (items []*Event) {
	for i := range h.Events {
		e := h.Events[i]
		if e.Level.Equal(EVENT_LEVEL_ERROR) {
			items = append(items, e)
		}
	}
	return
}

func (h *CallbackStatus) AddEvent(events ...*Event) {
	h.Events = append(h.Events, events...)
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
