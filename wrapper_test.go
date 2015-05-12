package tbw

import (
	"testing"
)

type TestApp struct {
	TB TermboxApi
}

func NewTestApp(tba TermboxApi) *TestApp {
	return &TestApp{
		TB: tba,
	}
}

func TestTermboxWrapperImplementsAllMethods(t *testing.T) {
	tbw := &TermboxWrapper{}
	_ = NewTestApp(tbw)
}
