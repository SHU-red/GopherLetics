package fancyfs

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/test"
)

func TestDetailsForFolder(t *testing.T) {
	test.NewApp()
	testdata := storage.NewFileURI("testdata")
	ff, err := DetailsForFolder(testdata)

	assert.Nil(t, err)
	path := ff.BackgroundURI.Path()
	assert.Equal(t, ".background.svg", path[len(path)-15:])
}

func TestDetailsForFolder_nil(t *testing.T) {
	ff, err := DetailsForFolder(nil)

	assert.Nil(t, ff)
	assert.Nil(t, err)
}
