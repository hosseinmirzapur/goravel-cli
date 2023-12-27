package transaction

import (
	"encoding/json"
	"fmt"

	"github.com/hosseinmirzapur/goravel-cli/prisma/logger"
)

type Result struct {
	cache []byte
}

func (r *Result) Get(c <-chan []byte, v interface{}) error {
	var res []byte
	if r.cache != nil {
		res = r.cache
	} else {
		data, ok := <-c
		if !ok {
			return fmt.Errorf("result not fetched")
		}
		res = data
		r.cache = data
	}
	logger.Debug.Printf("tx result: %s", res)
	if err := json.Unmarshal(res, &v); err != nil {
		return err
	}
	return nil
}
