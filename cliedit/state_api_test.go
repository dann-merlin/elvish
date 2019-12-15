package cliedit

import (
	"testing"

	"github.com/elves/elvish/cli"
	"github.com/elves/elvish/cli/el/codearea"
)

func TestInsertAtDot(t *testing.T) {
	f := setup()
	defer f.Cleanup()

	cli.SetCodeBuffer(f.Editor.app, codearea.Buffer{Content: "ab", Dot: 1})
	evals(f.Evaler, `edit:insert-at-dot XYZ`)

	testCodeBuffer(t, f.Editor, codearea.Buffer{Content: "aXYZb", Dot: 4})
}

func TestReplaceInput(t *testing.T) {
	f := setup()
	defer f.Cleanup()

	cli.SetCodeBuffer(f.Editor.app, codearea.Buffer{Content: "ab", Dot: 1})
	evals(f.Evaler, `edit:replace-input XYZ`)

	testCodeBuffer(t, f.Editor, codearea.Buffer{Content: "XYZ", Dot: 3})
}

func TestDot(t *testing.T) {
	f := setup()
	defer f.Cleanup()

	cli.SetCodeBuffer(f.Editor.app, codearea.Buffer{Content: "code", Dot: 4})
	evals(f.Evaler, `edit:-dot = 0`)

	testCodeBuffer(t, f.Editor, codearea.Buffer{Content: "code", Dot: 0})
}

func TestCurrentCommand(t *testing.T) {
	f := setup()
	defer f.Cleanup()

	evals(f.Evaler, `edit:current-command = code`)

	testCodeBuffer(t, f.Editor, codearea.Buffer{Content: "code", Dot: 4})
}

func testCodeBuffer(t *testing.T, ed *Editor, wantBuf codearea.Buffer) {
	t.Helper()
	if buf := cli.CodeBuffer(ed.app); buf != wantBuf {
		t.Errorf("content = %v, want %v", buf, wantBuf)
	}
}