// Copyright (c) 2022 Silverton Data, Inc.
// You may use, distribute, and modify this code under the terms of the AGPLv3 license, a copy of
// which may be found at https://github.com/silverton-io/honeypot/blob/main/LICENSE

package envelope

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/google/uuid"
)

type EventMeta struct {
	Protocol  string    `json:"protocol,omitempty"`
	Uuid      uuid.UUID `json:"uuid,omitempty"`
	Vendor    string    `json:"vendor,omitempty"`
	Namespace string    `json:"namespace,omitempty"`
	Version   string    `json:"version,omitempty"`
	Format    string    `json:"format,omitempty"`
	Schema    string    `json:"schema,omitempty"`
}

func (e EventMeta) Value() (driver.Value, error) {
	b, err := json.Marshal(e)
	return string(b), err
}

func (e EventMeta) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), &e)
}
