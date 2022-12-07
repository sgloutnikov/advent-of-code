package main

import (
	aoc_common "advent-of-code-2022"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type directory struct {
	Name        string
	Directories []*directory
	Files       []*file
	Parent      *directory
	Size        int64
}

type file struct {
	Name string
	Size int64
}

func main() {
	part1()
}

func part1() {
	fileTree := getFileTree()
	fmt.Println(fileTree)
	// TODO: Recurse tree and sum directory sizes less than 100000
}

func getFileTree() *directory {
	f := aoc_common.OpenFile("07/inputtest")
	defer f.Close()

	root := directory{Name: "/"}
	currentDir := &root

	fileScanner := bufio.NewScanner(f)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		cmdLine := strings.Split(line, " ")

		if strings.HasPrefix(cmdLine[0], "$") {
			switch cmdLine[1] {
			case "cd":
				if cmdLine[2] == ".." {
					currentDir = currentDir.Parent
				} else {
					for _, dir := range currentDir.Directories {
						if dir.Name == cmdLine[2] {
							currentDir = dir
						}
					}
				}
			default:
			}
		} else {
			// listing dirs/files
			switch cmdLine[0] {
			case "dir":
				dir := directory{Name: cmdLine[1], Parent: currentDir}
				currentDir.Directories = append(currentDir.Directories, &dir)
			default:
				size, _ := strconv.ParseInt(cmdLine[0], 10, 64)
				newFile := file{Name: cmdLine[1], Size: size}

				currentDir.Files = append(currentDir.Files, &newFile)
				updateDirSize(currentDir, size)
			}
		}
	}
	return &root
}

func updateDirSize(dir *directory, size int64) {
	dir.Size += size
	if dir.Parent != nil {
		updateDirSize(dir.Parent, size)
	}
}
