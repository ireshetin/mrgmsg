package Helper

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func ParseJsonIntoStruct(c []byte, v interface{}) error {

	c = bytes.Trim(c, "")

	err := json.Unmarshal(c, v)
	if err != nil {
		return fmt.Errorf("cant parse json (parseJsonIntoStruct): ", err)
	}

	return nil
}
