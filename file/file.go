package main

import (
	"fmt"
	"os"
)

func main() {
	// createFile()
	readFile()
}

func readFile() {
	userFile := "dash.txt"
	file, err := os.Open(userFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()
	buf := make([]byte, 1024*8)
	for {
		n, _ := file.Read(buf)
		if n == 0 {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}

// 创建文件，写文件
func createFile() {
	userFile := "dash.txt"
	fout, err := os.Create(userFile)
	if err != nil {
		fmt.Println(err)
	}

	defer fout.Close()
	for i := 0; i < 10; i++ {
		s := fmt.Sprintf("Dash Test haha %v\n", i)
		fout.WriteString(s)
		fout.Write([]byte("Just a test" + "\n"))
	}

}

func mkdir() {
	// 创建文件夹
	os.Mkdir("dash", 0777)
	os.MkdirAll("dash/path/file", 0744)
	err := os.Remove("dash")
	if err != nil {
		fmt.Println(err)
	}

	// 删除文件夹
	os.RemoveAll("dash")
}
