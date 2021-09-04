package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"io"
	"bufio"
	"strings"
)

// IsExist @title IsExist
// @description checks whether a file or directory exists
// @author DM
// @time 2021/4/9 13:36
// @param f
// @return bool It returns false when the file or directory does not exist
func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}

// IsFile @title IsFile
// @description checks whether the path is a file
// @author DM
// @time 2021/4/9 13:37
// @param f
// @return bool it returns false when it's a directory or does not exist
func IsFile(f string) bool {
	fi, e := os.Stat(f)
	if e != nil {
		return false
	}
	return !fi.IsDir()
}

// IsDir @title IsDir
// @description 判断所给路径是否为文件夹
// @author DM
// @time 2021/4/9 13:50
// @param path
// @return bool
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

//ListFiles 列出目录下所有文件（非递归）
func ListFiles(path string, pred ...func(os.FileInfo)bool) (infos [] os.FileInfo) {
	infos = make([]os.FileInfo, 0)
	finfos, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(pred) < 1 {
		infos = finfos
		return
	}

	p := pred[0]
	for _, info := range finfos {
		if p(info) {
			infos = append(infos, info)
		}
	}
	return
}

//ReadListFile 读取列表文件
func ReadListFile(fname string) (strs []string, err error) {
	fp, err := os.Open(fname)
	if err != nil {
		return
	}
	defer fp.Close()

	reader := bufio.NewReader(fp)
	for {
		line, _, err := reader.ReadLine()
		if err == nil {
			strs = append(strs, strings.TrimSpace(string(line)))
		} else if err == io.EOF {
			break
		} else {
			return nil, err
		}
	}
	if err == io.EOF {
		err = nil
	}
	return
}