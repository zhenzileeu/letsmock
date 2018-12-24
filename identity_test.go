package letsmock

import (
	"testing"
	"fmt"
)

var m = IdentityMock{}

func TestIdentityMock_MockIdCardNum(t *testing.T) {
	fmt.Println("mock id card num: ", m.MockIdCardNum())
}

