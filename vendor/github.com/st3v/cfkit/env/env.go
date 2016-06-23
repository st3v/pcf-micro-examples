package env

import (
	"fmt"
	"os"
)

func Addr() string {
	return fmt.Sprintf(":%s", os.Getenv("PORT"))
}
