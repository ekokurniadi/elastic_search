package helper

import (
	"database/sql/driver"
	"encoding/json"
)

type attributes map[string]interface{}

func (a attributes) Value() (driver.Value, error) {
	x := make(map[string]interface{})
	for key, val := range a {
		x[key] = val
	}

	return json.Marshal(x)
}
