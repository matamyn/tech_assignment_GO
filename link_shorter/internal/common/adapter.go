package common

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateShortLinkKey() string {
	rand.Seed(time.Now().UnixNano())
	short_link := strconv.FormatInt(rand.Int63n(1<<40), 32)
	return short_link
}
