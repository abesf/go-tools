package rand

import (
	"math/rand"
	"time"
)

func rundNum( int) int  {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
}