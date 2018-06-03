// this is doc
package main

// this is doc
import (
	"fmt"
	"log"
	"net/http"
	"regexp"
)

//this is doc
func handler(w http.ResponseWriter, r *http.Request) {
	regex := regexp.MustCompile("^jingjun.org/$")
	path := r.URL.Path[1:]
	match := regex.FindAllStringSubmatch(path, -1)
	if match != nil {
		fmt.Fprintf(w, "Hello %s\n", match[0][0])
		return
	}
	fmt.Fprintf(w, "Hello youtube %s\n", path)
}

type ByteSize float64

const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

type ByteSlice []byte

func (p *ByteSlice) Write(data []byte) (n int, err error) {
	slice := *p
	l := len(data)

	fmt.Printf("%d\n", cap(slice))

	if l+len(data) > cap(slice) {
		newslice := make([]byte, 2*(len(slice)+l))

		copy(newslice, newslice)
		slice = newslice
	}
	fmt.Printf("%d\n", cap(slice))
	for i, value := range data {
		slice[i+l] = value
	}
	*p = slice
	return len(data), nil

}
func (b ByteSize) String() string {
	switch {
	case b >= YB:
		return fmt.Sprintf("%.2fYB", b/YB)
	case b >= ZB:
		return fmt.Sprintf("%.2fZB", b/ZB)
	case b >= EB:
		return fmt.Sprintf("%.2fEB", b/EB)
	case b >= PB:
		return fmt.Sprintf("%.2fPB", b/PB)
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%.2fB", b)
}
func main() {
	fmt.Println("%v", TB)
	student := map[string]int{

		"jingjun":  24,
		"xiaoming": 34,
	}
	x := []int{1, 23, 3}
	y := []int{4, 5, 6}
	x = append(x, y...)
	fmt.Printf("%v\n", x)
	for key, value := range student {
		fmt.Printf("%s=%d\n", key, value)
	}

	var b ByteSlice
	fmt.Fprintf(&b, "This hour has %d days\n", 7)
	fmt.Printf("%v\n", b)

	a := ByteSize(1e13)
	fmt.Println(a)
	fmt.Println(TB)
	http.HandleFunc("/", handler)
	// http.ListenAndServe(":8080", nil)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
