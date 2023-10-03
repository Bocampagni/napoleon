package main

import (
	"fmt"
	"math/rand"
    "time"
	"github.com/bocampagni/napoleon/assets"
)

func main() {
	s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
	fmt.Println(assets.Quotes[r1.Intn(len(assets.Quotes))])
}