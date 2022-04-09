package ntree

import (
	"fmt"
	"testing"
)

func TestPath(t *testing.T) {
	var tree *NTree = GetNAryTree()
	tree.EnsurePath("a.b.c")
	tree.PreOrder(tree.Root, 1)
	tree.EnsurePath("a.b.c.d.e.f")
	tree.PreOrder(tree.Root, 1)
	tree.EnsurePath("a.x.y")
	tree.PreOrder(tree.Root, 1)
	tree.EnsurePath("a.z.y")
	tree.PreOrder(tree.Root, 1)

}

func TestConvert(t *testing.T) {

	t.Fail()

	var tree *NTree = GetNAryTree()
	/*
			   10
			  /   \
			 /     \
			/       \
		   8         5
		  /|\      /|\ \
		 / | \    / | \ \
		2  1  6  7  8 3  4
		  / \     \     /| \
		 9  11     -1  2 1  3
		-----------------------
		Constructing N-Arr tree
	*/
	// First element of tree
	tree.Root = NewTreeNode("10")
	tree.Root.addChild(NewTreeNode("8"))
	tree.Root.addChild(NewTreeNode("5"))
	// Add child node [2,1,6] in node (8)

	c := tree.Root.GetChildren()
	c[0].addChild(NewTreeNode("2"))
	c[0].addChild(NewTreeNode("1"))
	c[0].addChild(NewTreeNode("6"))
	// Add child node [9,11] in node (1)
	c[1].addChild(NewTreeNode("9"))
	c[1].addChild(NewTreeNode("11"))
	// Add child node [7  8 3  4] in node (5)
	c[1].addChild(NewTreeNode("7"))
	c[1].addChild(NewTreeNode("8"))
	c[1].addChild(NewTreeNode("3"))
	c[1].addChild(NewTreeNode("4"))
	// Add child node [-7] in node (4)
	c[1].GetChildren()[0].addChild(NewTreeNode("-1"))
	// Add child node [2,1,3] in node (7)
	c[1].GetChildren()[3].addChild(NewTreeNode("2"))
	c[1].GetChildren()[3].addChild(NewTreeNode("1"))
	c[1].GetChildren()[3].addChild(NewTreeNode("3"))
	fmt.Print(" Preorder : \n")
	// Print tree element
	tree.PreOrder(tree.Root, 1)
	fmt.Print(" Sorted : \n")
	// /tree.sortTree()
	tree.PreOrder(tree.Root, 1)

}
