package test

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"testing"
)

func TestGenerateUUID(t *testing.T) {
	v4 := uuid.NewV4().String()
	fmt.Println(v4)
}
