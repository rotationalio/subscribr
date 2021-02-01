/*
Package subscribr is a super awesome utility for capturing emails for people who are
interested in what we might do in the future!
*/
package subscribr

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("First init!")
}

func New() error {
	return nil
}
