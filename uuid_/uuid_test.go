package uuid_

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func TestUUID(t *testing.T) {
	uuid := uuid.New()
	key := uuid.String()
	fmt.Println(key)
}
