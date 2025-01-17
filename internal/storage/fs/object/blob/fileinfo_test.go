package blob

import (
	"io/fs"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestFileInfo(t *testing.T) {
	modTime := time.Now()
	fi := NewFileInfo("f.txt", 100, modTime)
	require.Equal(t, fs.FileMode(0), fi.Type())
	require.Equal(t, "f.txt", fi.Name())
	require.Equal(t, int64(100), fi.Size())
	require.Equal(t, modTime, fi.ModTime())
	require.Equal(t, false, fi.isDir)
	info, err := fi.Info()
	require.NoError(t, err)
	require.Equal(t, fi, info)
	require.Nil(t, fi.Sys())
}

func TestFileInfoIsDir(t *testing.T) {
	fi := FileInfo{}
	fi.SetDir(true)
	require.Equal(t, true, fi.isDir)
}
