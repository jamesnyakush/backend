package model

import "github.com/gofrs/uuid"

type Notice struct {
	NoticeId uuid.UUID `json:"notice_id"`
}
