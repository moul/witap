package witap

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var inputs = []string{
	`
d  S------+    b
          |
          |
   c      +--->a
`, `
S-----+---a->c
      |
      V
      b
`, `
a S      s
  |      |
  V      V
  b      c
`, `
d s<+S+--V
    |||  Q
    -++
`, `
d s-+   +-S  +--+
    +-->b |  |  |
     |  | +--+  |
     +--+ A<----+
`, `
S-----+
|     +-^
+---+->B
    +---^
`,
}
var outputs = []string{
	"a", "b", "b", "Q", "A", "B",
}

func ExampleResolver_a() {
	resolver := NewResolverFromString(inputs[0])
	fmt.Println(resolver)
	// fixme: add output
}

func TestResolver(t *testing.T) {
	Convey("Testing resolver", t, func() {
		for idx, input := range inputs {
			Convey(fmt.Sprintf("input %d", idx+1), func() {
				resolver := NewResolverFromString(input[1 : len(input)-1])
				output, err := resolver.Run()
				So(err, ShouldBeNil)
				So(output, ShouldEqual, outputs[idx])
			})
		}
	})
}