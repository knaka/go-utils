package fs

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/knaka/go-testutils/fsassert"
	"github.com/stretchr/testify/assert"

	//revive:disable-next-line:dot-imports
	. "github.com/knaka/go-utils"
)

func TestCopy(t *testing.T) {
	tempDir := Value(os.MkdirTemp("", "copy_dir_test"))

	srcDir := filepath.Join("testdata", "dir1")
	dstDir := filepath.Join(tempDir, "dirX")
	Must(Copy(srcDir, dstDir))
	fsassert.DirsAreEqual(t, srcDir, dstDir)

	srcFile := filepath.Join("testdata", "dir1", "bar.txt")
	dstFile := filepath.Join(tempDir, "bar.txt")
	Must(Copy(srcFile, dstFile))
	fsassert.FilesAreEqual(t, srcFile, dstFile)
}

func TestRealpath(t *testing.T) {
	type args struct {
		s string
	}
	wd := Value(os.Getwd())
	wd = Value(filepath.Abs(wd))
	wd = Value(filepath.EvalSymlinks(wd))
	wd = filepath.Clean(wd)
	tests := []struct {
		name    string
		args    args
		wantRet string
		wantErr bool
	}{
		{"Test",
			args{filepath.Join(".", "testdata", "dir2", "link1", "bar.txt")},
			filepath.Join(wd, "testdata", "dir1", "bar.txt"),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRet, err := Realpath(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Realpath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRet != tt.wantRet {
				t.Errorf("Realpath() gotRet = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}

func TestTouch(t *testing.T) {
	tempDir := Value(os.MkdirTemp("", "touch_test"))
	filePath := filepath.Join(tempDir, "foo.txt")

	Must(Touch(filePath))
	stat, err := os.Stat(filePath)
	assert.Nil(t, err)
	assert.True(t, stat.Size() == 0)

}

func TestMove(t *testing.T) {
	type args struct {
		src string
		dst string
	}
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		//{
		//	"Test",
		//	args{
		//		"/tmp/test.txt",
		//		"/Users/knaka/test.txt",
		//	},
		//	assert.NoError,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantErr(t, Move(tt.args.src, tt.args.dst), fmt.Sprintf("Move(%v, %v)", tt.args.src, tt.args.dst))
		})
	}
}
