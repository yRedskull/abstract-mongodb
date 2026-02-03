package mongodb

import (
	"fmt"
)

func ValFindParams(findParams FindParams) error {
	if findParams.Collection == "" {
		return fmt.Errorf("findParams.Collection vazio")
	}

	if len(findParams.Filter) == 0 {
		return fmt.Errorf("findParams.Filter vazio")
	}

	return nil
}

