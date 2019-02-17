package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("Newからの戻り値がNil")
	} else {
		tracer.Trace("Hello trace")
		if buf.String() != "Hello trace\n" {
			t.Errorf("'%s'という誤った文字列", buf.String())
		}
	}
}
