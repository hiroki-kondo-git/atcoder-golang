package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	initialBufSize = 1e4
	maxBufSize     = 1e6 + 7
)

var buf []byte = make([]byte, initialBufSize)
var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(buf, maxBufSize)
}

func main() {
	x := nextInt()
	if x%100 == 0 && x != 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

/*==========================================
 *             Library
 *==========================================*/

// 一行ずつ読み込む
func nextLine() string {
	sc.Scan()
	return sc.Text()
}

// スペース区切りで読み込む
func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
