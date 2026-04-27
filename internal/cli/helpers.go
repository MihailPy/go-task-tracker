package cli

import (
	"fmt"
	"strconv"
)

func parseTaskID(arg string) (int, error) {
	id, err := strconv.Atoi(arg)
	if err != nil {
		return 0, fmt.Errorf("неверный формат ID: %s (должно быть число)", arg)
	}
	return id, nil
}
