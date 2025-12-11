package flow

import (
	"fmt"
	"testing"
	"time"
)

func TestC(t *testing.T) {

	svc, err := CreateCacheService("redis://localhost:6379")

	fmt.Print(err)

	time.Sleep((3 * time.Second))

	svc.Stop()

	time.Sleep((5 * time.Second))
}
