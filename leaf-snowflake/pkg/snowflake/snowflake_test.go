package snowflake_test

import (
	"fmt"
	"testing"

	"leaf-snowflake/pkg/snowflake"
)

func TestGenerateID(t *testing.T) {
	node1 := snowflake.NewNode(0)
	node2 := snowflake.NewNode(1)
	for i := 0; i < 3; i++ {
		id1, _ := node1.NextId()
		id2, _ := node2.NextId()
		fmt.Println(id1, id2)
	}
}
