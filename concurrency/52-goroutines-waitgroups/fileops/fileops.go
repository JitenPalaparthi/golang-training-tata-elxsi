package fileops

import (
	"bufio"
	"log"
	. "os" // using . can call functions and types directlyh without using package name
)

var FileName string // this is a package level variable

func init() {
	println("init func- from fileops- >1")
}

func init() {
	println("init func- from fileops ->2", FileName)
}

func ReadFileIntoSlice() ([]string, error) {
	lines := make([]string, 0)
	if file, err := Open(FileName); err != nil { // os.Open
		log.Println(err.Error())
		return nil, err
	} else {
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() { // as long as the scanner.Scan gives true
			line := scanner.Text()
			lines = append(lines, line)
		}
		return lines, nil
	}

}
