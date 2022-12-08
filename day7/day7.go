package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
    "strings"
)

type dir struct {
    depth int
    files []int
    dirs []string
}

func getFiles(in []string, n string) []int {

    var files []int
    indir := false
    for _,line := range in {
        if line == fmt.Sprintf("$ cd %s",n){
            indir = true
            continue
        }
        if indir == true && strings.HasPrefix(line,"$ cd") {
            break
        }
        if strings.HasPrefix(line,"$ ls") {
            continue
        }
        
        if indir {
            if !strings.HasPrefix(line,"dir") {
                ss,_,_ := strings.Cut(line," ")
                s,_ := strconv.Atoi(ss)
                files = append(files,s)
            }
        }
    }

    return files
}

func getDirs(in []string, n string) []string {

    var dirs []string
    indir := false
    for _,line := range in {
        if line == fmt.Sprintf("$ cd %s",n){
            indir = true
            continue
        }
        if indir == true && strings.HasPrefix(line,"$ cd") {
            break
        }
        if strings.HasPrefix(line,"$ ls") {
            continue
        }
        
        if indir {
            if strings.HasPrefix(line,"dir") {
                _,d,_ := strings.Cut(line," ")
                dirs = append(dirs,d)
            }
        }
    }

    return dirs
}

func getDepth(in []string) map[string]int {

    depthMap := make(map[string]int)

    var fileDepth []string
    for _,line := range in {
        if strings.HasPrefix(line,"$ cd") {
            switch _,d,_ := strings.Cut(line,"d "); d {
            case "..":
                depthMap[fileDepth[len(fileDepth)-1]] = len(fileDepth) - 1
                fileDepth = fileDepth[:len(fileDepth)-1]
            default:
                fileDepth = append(fileDepth,d)
            }
            depthMap[fileDepth[len(fileDepth)-1]] = len(fileDepth) -1 
        }
    }
    return depthMap
}

func getBelowFiles(d dir, dir_list map[string]dir) []int {

    var newFiles []int

    fmt.Println(d,"has subdirs",d.dirs)
    for _,v := range d.dirs {
        fmt.Println("getting the files from",v)
        newFiles = append(newFiles,dir_list[v].files...)
        fmt.Println(newFiles)
    }

    return newFiles
}

func main() {
	file, err := os.Open("day7/day7.input")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

    dir_list := make(map[string]dir)
    dir_list["/"] = dir{}
    for _,line := range lines {
        if strings.HasPrefix(line,"dir") {
            _,d,_ := strings.Cut(line," ")
            dir_list[d] = dir{}
        }                
    }
    dm := getDepth(lines)
    lowest := 0
    for _,d := range dm {
        if d > lowest {
           lowest = d 
        }
    }

    for k,v := range dir_list {
        dd := getDirs(lines,k)
        v.dirs = dd
        ff := getFiles(lines,k)
        v.files = ff
        v.depth = dm[k]
        dir_list[k] = v
    }
   
    fmt.Println("lowest group is",lowest)
    for i:=lowest-1;i>=0;i-- {
        fmt.Println("getting group",i)
        for k,v := range dir_list {
            if v.depth == i {
                fmt.Println(v)
                newFiles := getBelowFiles(v,dir_list)
                v.files = append(v.files,newFiles...)
                dir_list[k] = v
                fmt.Println(k,dir_list[k].files)
            }
        }
    }

    final := 0
    for _,d := range dir_list {
        st := 0
        for _,f := range d.files {
            st += f
        }
        if st <= 100000 {
            final += st
            fmt.Println(d,st)
        }
    }

    fmt.Println(final)

}
