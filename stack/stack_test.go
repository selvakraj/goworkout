package stack_test

import (
	"fmt"
	"testing"

	"github.com/selvakraj/goworkout/stack"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestPush(t *testing.T) {

	t.Log("Given the need to test Push functionality.")
	{
		const items = 5
		t.Logf("\tTest 0:\tWhen pushing %d items", items)
		{
			var s stack.Stack

			var orgData string
			for i := 0; i < items; i++ {
				name := fmt.Sprintf("Name%d", i)
				orgData = name + orgData
				s.Push(&stack.Data{Name: name})
			}
			fmt.Println("Stack", s)

			if s.Count() != items {
				t.Logf("\t%s\tTest 0:\tShould be able to push %d items.", failed, items)
				t.Fatalf("\t\tTest 0:\tGot %d, Expected %d.", s.Count(), items)
			}
			t.Logf("\t%s\tTest 0:\tShould be able to push %d items.", succeed, items)
		}
	}

}
