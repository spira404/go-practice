package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)


var level int = 0
var abslevels = make(map[int]bool) // abs

func dirTree(out io.Writer, path string, printFiles bool) error{

	elements, _ := ioutil.ReadDir(path)
	foldersnum := 0
	
	for _, element := range elements {
		if element.IsDir(){
			foldersnum += 1
		}
	}

	for el_num, element := range elements {

		// folders
		if element.IsDir() {

			// print connecting pipes
			for lev_num := 0; lev_num < level; lev_num++{
				if _, ok := abslevels[lev_num]; ok {
					io.WriteString(out, "│\t")
					continue
				}
				io.WriteString(out, "\t")
			}
			
			foldersnum -= 1
			// detect and print ending pipes
			if (el_num == len(elements) - 1) || (!printFiles && foldersnum == 0){
				io.WriteString(out, "└───")
			}else{
				io.WriteString(out, "├───")
			}

			io.WriteString(out, element.Name() + "\n")
			if (el_num == len(elements) - 1) || (!printFiles && foldersnum == 0){
				delete(abslevels, level) // abs
			}else{
				abslevels[level] = true // abs
			}
			// recursive part
			level++ 
			folder := filepath.Join(path, element.Name())
			dirTree(out, folder, printFiles)
			level--
			continue
		}


		// files
		if printFiles{

			// print connecting pipes
			for lev_num := 0; lev_num < level; lev_num++{
				if _, ok := abslevels[lev_num]; ok {
					io.WriteString(out, "│\t")
					continue
				}
				io.WriteString(out, "\t")
			}

			// detect and print ending pipes
			if el_num == len(elements) - 1 {
				io.WriteString(out, "└───")
			}else{
				io.WriteString(out, "├───")
			}

			// print out filename and filesize
			filesize := element.Size()
			if filesize == 0 {
				io.WriteString(out, element.Name() + " (empty)\n")
			}else{
				io.WriteString(out, element.Name() + fmt.Sprintf(" (%vb)\n", filesize))
			}
		}
	}
	return nil
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)

	if err != nil {
		panic(err.Error())
	}
}
