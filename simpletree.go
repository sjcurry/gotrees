package simpletree

import (
	"fmt"
	"os"
	"strings"
)

type Tree struct {
	value string
	left  *Tree
	right *Tree
}

type Forest []*Tree

func Initforest(maxValueLength int) Forest {
	var f Forest
	var i int = 1

	for i < maxValueLength {
		f = append(f, nil)
		i += 1
	}

	return f
}

func AddIfNew(frst Forest, value string) (found bool) {
	tinst := frst[len(value)]

	if tinst == nil {

		tinst = new(Tree)
		tinst.value = value
		fmt.Fprintf(os.Stdout, "%s\n", value)
		return true
	} else {

		x := strings.Compare(tinst.value, value)
		if x == 0 {
			return false
		}
		if x == 1 {
			return AddIfNew(frst, tinst.left.value)
		} else {
			return AddIfNew(frst, tinst.right.value)
		}
	}
}

func PrintTree(forest Forest, length int){

	var ptree func(tree *Tree) 
	ptree = func(tree *Tree) {
		if tree == nil {
			return
		} else{
			ptree(tree.left)
			fmt.Println(tree.value)
			ptree(tree.right)
		}

	}
	if forest[length] == nil { return}

	fmt.Printf("\nPrinting tree with value length %d\n", length)
	ptree(forest[length])
	fmt.Println("=========================================")
}

