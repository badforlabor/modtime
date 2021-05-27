/**
 * Auth :   liubo
 * Date :   2021/5/27 14:50
 * Comment:
 */

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	var t = time.Now()

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		os.Chtimes(path, t, t)
		return nil
	})
	fmt.Println("work done.")
}


