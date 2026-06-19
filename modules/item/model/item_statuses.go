package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type ItemStatus int

const (
	ItemStatusENABLE ItemStatus = iota
	ItemStatusDISABLE
	ItemStatusDELETED
)

var allItemStatus = [3]string{"ENABLE", "DISABLE", "DELETED"}

func (item *ItemStatus) String() string {
	if int(*item) < 0 || int(*item) >= len(allItemStatus) {
		return "UNKNOWN"
	}
	return allItemStatus[*item]
}

func parseItemStatus(s string) (ItemStatus, error) {
	for i, v := range allItemStatus {
		if v == s {
			return ItemStatus(i), nil
		}
	}
	return 0, fmt.Errorf("invalid ItemStatus: %s", s)
}

func (item ItemStatus) Value() (driver.Value, error) {
	return item.String(), nil
}

func (item ItemStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, item.String())), nil
}

func (item *ItemStatus) UnmarshalJSON(data []byte) error {
	var strValue string
	if err := json.Unmarshal(data, &strValue); err != nil {
		return err
	}

	parsed, err := parseItemStatus(strValue)
	if err != nil {
		return err
	}

	*item = parsed

	return nil
}

func (item *ItemStatus) Scan(value interface{}) error {
	if value == nil {
		*item = ItemStatusENABLE
		return nil
	}

	var strValur string
	switch v := value.(type) {
	case string:
		strValur = v
	case []byte:
		strValur = string(v)
	default:
		return errors.New("invalid type for ItemStatus")
	}

	parsed, err := parseItemStatus(strValur)
	if err != nil {
		return err
	}

	*item = parsed

	return nil
}