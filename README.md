#tbw

Interface and a API wrapper for [termbox-go](https://github.com/nsf/termbox-go).

##Use case

Testing and abstraction.

##Example

```go
import (
    tbw "github.com/morcmarc/termbox-wrapper"
)

type MyApp struct {
    TBInstance tbw.TermboxApi
}

// In source file
func main() {
    // NewTermboxWrapper() will return a dummy struct that delegates all calls,
    // just like if you'd be using termbox directly.
    tbwrapper := tbw.NewTermboxWrapper()
    ma := &MyApp{
      TBInstance: tbwrapper
    }
}

// In tests
// You can mock out the termbox depenceny and add your own method
// implementations.
type TBMock struct {
    tbw.TermboxApi
}

func (tbm TBMock) Init() error {
    ...
}

...
```