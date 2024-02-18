package helpers

import (
	"fmt"
	"testing"
)

func TestRandomString(t *testing.T) {
	fmt.Println(RandomString(16))
	fmt.Println(RandomString(16))
}
