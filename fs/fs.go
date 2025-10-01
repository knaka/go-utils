// Package fs is utilities for filesystem.
package fs

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	//revive:disable-next-line:dot-imports
	. "github.com/knaka/go-utils"
)

// IsRootDir returns true if the given directory is the root directory.
func IsRootDir(dir string) bool {
	dirPath, err := Realpath(dir)
	if err != nil {
		return false
	}
	return dirPath == filepath.Dir(dirPath)
}

// IsSubDir returns true if subDir is a subdirectory of parentDir.
func IsSubDir(subDir string, parentDir string) (bool, error) {
	subDir, err := Realpath(subDir)
	if err != nil {
		return false, err
	}
	parentDir, err = Realpath(parentDir)
	if err != nil {
		return false, err
	}
	return strings.HasPrefix(subDir, parentDir), nil
}

// Realpath returns the canonical absolute path of the given value.
func Realpath(s string) (ret string, err error) {
	ret, err = filepath.Abs(s)
	if err != nil {
		return
	}
	ret, err = filepath.EvalSymlinks(ret)
	if err != nil {
		return
	}
	ret = filepath.Clean(ret)
	return
}

func copyFile(src, dst string) (err error) {
	reader := V(os.Open(src))
	defer (func() { V0(reader.Close()) })()
	writer := V(os.Create(dst))
	defer (func() { Ignore(writer.Close()) })()
	V0(io.Copy(writer, reader))
	V0(writer.Close())
	statSrc := V(os.Stat(src))
	V0(os.Chmod(dst, statSrc.Mode()))
	V0(os.Chtimes(dst, statSrc.ModTime(), statSrc.ModTime()))
	return
}

func copyDir(srcDir string, dstDir string) (err error) {
	return filepath.Walk(srcDir, func(srcObjPath string, srcObjStat fs.FileInfo, errGiven error) (err error) {
		if errGiven != nil {
			return errGiven
		}
		dstObj := filepath.Join(dstDir, strings.TrimPrefix(srcObjPath, srcDir))
		if srcObjStat.IsDir() {
			err = os.MkdirAll(dstObj, srcObjStat.Mode())
			if err != nil {
				return
			}
		} else if !srcObjStat.Mode().IsRegular() {
			switch srcObjStat.Mode().Type() & os.ModeType {
			case os.ModeSymlink:
				linkDst, err2 := os.Readlink(srcObjPath)
				if err2 != nil {
					return err2
				}
				err = os.Symlink(linkDst, dstObj)
				if err != nil {
					return
				}
			default:
				err = nil
			}
		} else {
			return copyFile(srcObjPath, dstObj)
		}
		err = os.Chtimes(dstObj, srcObjStat.ModTime(), srcObjStat.ModTime())
		return err
	})
}

// Copy copies a file or a directory.
func Copy(src string, dst string) (err error) {
	if stat, err2 := os.Stat(src); err2 != nil {
		return err2
	} else if stat.IsDir() {
		if _, err2 := os.Stat(dst); err2 == nil {
			Ignore(os.RemoveAll(dst))
		}
		return copyDir(src, dst)
	}
	if _, err2 := os.Stat(dst); err2 == nil {
		Ignore(os.RemoveAll(dst))
	}
	return copyFile(src, dst)
}

// Move moves a file or a directory.
func Move(src string, dst string) (err error) {
	err = Copy(src, dst)
	if err != nil {
		return
	}
	return os.RemoveAll(src)
}

// Touch creates an empty file if it doesn't exist, or updates its modification time if it does.
func Touch(path string) (err error) {
	_, err = os.Stat(path)
	if err != nil {
		// If not exists, create an empty file.
		if os.IsNotExist(err) {
			file, err2 := os.Create(path)
			if err2 != nil {
				return err2
			}
			return file.Close()
		}
		return err
	}
	// If exists, update the timestamp.
	now := time.Now()
	return os.Chtimes(path, now, now)
}
