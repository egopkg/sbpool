package sbpool_test

import (
	"strings"
	"testing"

	"github.com/ergopkg/sbpool"
)

var sb *strings.Builder

func TestAcquireStringsBuilder(t *testing.T) {
	sb = sbpool.AcquireStringsBuilder()
	sb.WriteString("hello")
	sb.WriteString(" ")
	sb.WriteString("world")

	if sb.String() != "hello world" {
		t.Fatal("wrong resulted string")
	}
}

func TestReleaseStringsBuilder(t *testing.T) {
	sbpool.ReleaseStringsBuilder(sb)
}
