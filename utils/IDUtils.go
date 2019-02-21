package utils

import (
	"math/rand"
	"strconv"
	"time"
)

type IDUtils struct {
}

func (b *IDUtils) GenItemId() int64 {
	o := rand.New(rand.NewSource(time.Now().UnixNano())).Int63()
	return o

}
func (b *IDUtils) GenImageName() string {
	o := rand.New(rand.NewSource(time.Now().UnixNano())).Int63()
	s := strconv.FormatInt(o, 10)

	return s

}
