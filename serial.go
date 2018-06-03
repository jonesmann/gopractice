package main

import (
	"time"
	"sort"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"os"
	"crypto/md5"
)

func main() {
	start := time.Now()
	m, err := MkAllMd5(os.Args[1])
	var paths []string
	for path := range m {
		paths=append(paths, path)
	}
	if err != nil{
		return
	}
	sort.Strings(paths)
	for _, path := range paths{
		fmt.Println(path, m[path])
	}
	d := time.Since(start)
	fmt.Println(d)
}

func MkAllMd5(root string) (map[string][md5.Size]byte, error){
	m := make(map[string][md5.Size]byte)
	err:=filepath.Walk(root, func(path string, fileinfo os.FileInfo, err error) error{
		if err != nil{
			return err
		}
		if  ! fileinfo.Mode().IsRegular(){
			return nil
		}
		data, err := ioutil.ReadFile(path)
		if err != nil{
			return err
		}
		m[path] = md5.Sum(data)
		return nil
	})

	return m, err
}
