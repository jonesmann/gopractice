package main

import (
	"time"
	"fmt"
	"sort"
	"errors"
	"io/ioutil"
	"os"
	"sync"
	"path/filepath"
	"crypto/md5"
)

type md5info struct{
	path string
	sum [md5.Size]byte
	err error
}
func sumFiles(done chan struct{}, root string)(<-chan md5info, <- chan error){
	c := make(chan md5info)
	errc :=make(chan error)

	go func(){
		var wg sync.WaitGroup
		err := filepath.Walk(root, func(path string, fileinfo os.FileInfo, err error) error{
			if err != nil{
				return err
			}
			if ! fileinfo.Mode().IsRegular(){
				return nil
			}
			wg.Add(1)
			go func(){
				data, err := ioutil.ReadFile(path)
				select {
				case c<- md5info{path, md5.Sum(data),err}:
				case <-done:
				}
				wg.Done()
			}()
			select{
			case <-done:
				return errors.New("walk cancel")
			default:
				return nil
			}
		})
		go func(){
			wg.Wait()
			close(c)
		}()
		errc <- err
	}()
	return c, errc
}

func Md5All(root string)(map[string][md5.Size]byte, error){
	done := make(chan struct{})
	defer close(done)

	c ,errc := sumFiles(done, root)

	m := make(map[string][md5.Size]byte)
	for r := range c{
		if r.err !=nil{
			return nil, r.err
		}
		m[r.path] = r.sum

	}
	if err := <- errc; err != nil {
		return nil,err
	}
	return m, nil
}

func main(){
	start := time.Now()
	m,err := Md5All(os.Args[1])
	if err != nil{
		return
	}
	var paths []string
	for path := range m{
		paths = append(paths, path)
	}
	sort.Strings(paths)
	for _, path := range paths{
		fmt.Println(path, m[path])
	}
	d := time.Since(start)
	fmt.Println(d)
}