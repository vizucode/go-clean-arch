package helpers

import (
	"strconv"
	"time"
)

func TFCode(numberUnique uint) string {
	unique := strconv.Itoa(int(numberUnique))
	epoch := strconv.Itoa(int(time.Now().Unix()))

	return "TF-" + epoch + unique
}
