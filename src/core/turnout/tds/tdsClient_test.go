package tds

import (
	"io"
	"testing"
)

func TestThattdsClientImplementsIOCloser(t *testing.T) {
	var _ io.Closer = &tdsClient{}
}
