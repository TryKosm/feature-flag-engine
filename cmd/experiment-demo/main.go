package main

import (
	"fmt"

	"feature-flag-engine/internal/engine"
)

func main() {
	flag := engine.Flag{Key: "new-homepage", RolloutPercent: 30}
	fmt.Println(engine.Enabled(flag, "user-101"))
}
