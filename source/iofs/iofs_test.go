package iofs_test

import (
	"testing"

	"lilac.ooo/calliope/source/iofs"
	st "lilac.ooo/calliope/source/testing"
)

func Test(t *testing.T) {
	// reuse the embed.FS set in example_test.go
	d, err := iofs.New(fs, "testdata/migrations")
	if err != nil {
		t.Fatal(err)
	}

	st.Test(t, d)
}
