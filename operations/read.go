package operations

import (
	"fmt"
	"os"
)

func ReadAll(condition bool, db string) {
	buffers, err := os.ReadFile(db)
	if err != nil {
		fmt.Println(err)
	}
}

func ReadOne(condition bool) {}

func ReadWhere(condition bool, limit int) {}
