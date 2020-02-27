package cache

import (
	"fmt"
	"testing"
)

func TestNewCache(t *testing.T) {
	c := NewCache("127.0.0.1:6379", "")
	fmt.Println(c.Many([]string{"hello", "age", "name"}))
}
