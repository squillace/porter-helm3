package helm3

import (
	"strings"
	"testing"

	"github.com/deislabs/porter/pkg/porter/version"
	"github.com/deislabs/porter/pkg/printer"
	"github.com/squillace/porter-helm3/pkg"
	"github.com/stretchr/testify/require"
)

func TestPrintVersion(t *testing.T) {
	pkg.Commit = "abc123"
	pkg.Version = "v1.2.3"

	m := NewTestMixin(t)

	opts := version.Options{}
	err := opts.Validate()
	require.NoError(t, err)
	m.PrintVersion(opts)

	gotOutput := m.TestContext.GetOutput()
	wantOutput := "helm3 v1.2.3 (abc123) by ralph squillace"
	if !strings.Contains(gotOutput, wantOutput) {
		t.Fatalf("invalid output:\nWANT:\t%q\nGOT:\t%q\n", wantOutput, gotOutput)
	}
}

func TestPrintJsonVersion(t *testing.T) {
	pkg.Commit = "abc123"
	pkg.Version = "v1.2.3"

	m := NewTestMixin(t)

	opts := version.Options{}
	opts.RawFormat = string(printer.FormatJson)
	err := opts.Validate()
	require.NoError(t, err)
	m.PrintVersion(opts)

	gotOutput := m.TestContext.GetOutput()
	wantOutput := `{
  "name": "helm3",
  "version": "v1.2.3",
  "commit": "abc123",
  "author": "ralph squillace"
}
`
	if !strings.Contains(gotOutput, wantOutput) {
		t.Fatalf("invalid output:\nWANT:\t%q\nGOT:\t%q\n", wantOutput, gotOutput)
	}
}
