package rand

import (
	"math/rand"
	"time"
)
//todo
func rundNum( int) int  {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
}