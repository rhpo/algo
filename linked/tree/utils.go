package tree

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var nodeIDCounter int

func GenerateUniqueID() int {
	nodeIDCounter++
	return nodeIDCounter
}

func Scan(T *Node, slice *[]string, nodeMap map[*Node]int) {
	if _, ok := nodeMap[T]; !ok {
		nodeMap[T] = GenerateUniqueID()
	}

	*slice = append(*slice, fmt.Sprintf("%d [label=\"%d\"]", nodeMap[T], T.Data))

	if T.Left != nil || T.Right != nil {
		if T.Left != nil {
			if _, ok := nodeMap[T.Left]; !ok {
				nodeMap[T.Left] = GenerateUniqueID()
			}
			*slice = append(*slice, fmt.Sprintf("%d -> %d", nodeMap[T], nodeMap[T.Left]))
			Scan(T.Left, slice, nodeMap)
		} else {
			id = GenerateUniqueID()
			*slice = append(*slice, fmt.Sprintf("null%d [shape=\"point\"];", T.Data*2+id))
			*slice = append(*slice, fmt.Sprintf("%d -> null%d;", nodeMap[T], T.Data*2+id))
		}

		if T.Right != nil {
			if _, ok := nodeMap[T.Right]; !ok {
				nodeMap[T.Right] = GenerateUniqueID()
			}

			*slice = append(*slice, fmt.Sprintf("%d -> %d", nodeMap[T], nodeMap[T.Right]))
			Scan(T.Right, slice, nodeMap)
		} else {
			id = GenerateUniqueID()
			*slice = append(*slice, fmt.Sprintf("null%d [shape=\"point\"];", T.Data*2+1+id))
			*slice = append(*slice, fmt.Sprintf("%d -> null%d;", nodeMap[T], T.Data*2+1+id))
		}
	}
}

func Convert(T *Tree, outputfile string) {
	file, err := os.Create(outputfile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := "digraph tree {" + "\n"
	slice := []string{}
	nodeMap := make(map[*Node]int)
	Scan(T.Root, &slice, nodeMap)
	result += strings.Join(slice, "\n")
	result += "\n}"
	file.WriteString(result)
}

var id int = 0

func showTree(T *Tree) {
	id++

	outputfile := fmt.Sprintf("%d.dot", id)
	outputImageFile := fmt.Sprintf("%d.png", id)

	Convert(T, outputfile)

	if !isDotCommandInstalled() {
		err := os.Remove(outputfile)
		if err != nil {
			log.Fatal(err)
		}
		log.Fatal("Graphviz (dot command) is not installed, please install it in your system.")
	} else {
		cmd := exec.Command("dot", "-Tpng", outputfile, "-o", outputImageFile)
		er := cmd.Run()
		if er != nil {
			log.Fatal(er)
		}

		cmd = exec.Command("open", outputImageFile)
		er = cmd.Run()

		// delete the file
		// err := os.Remove(outputfile)

		// if err != nil {
		// 	log.Fatal(err)
		// }

		// err = os.Remove(outputImageFile)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// if er != nil {
		// 	log.Fatal(er)
		// }
	}
}
