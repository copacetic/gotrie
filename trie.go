package main

import (
    "bufio"
    "fmt"
    "os"
)

type TrieNode struct {
    val      string
    children map[uint8]*TrieNode
}

func NewTrieNode() *TrieNode {
    tn := TrieNode{
        children: make(map[uint8]*TrieNode),
        val:      "",
    }
    return &tn
}

func (root *TrieNode) Contains(test_me string) bool {
    node := root
    strlen := len(test_me)
    for i := 0; i < strlen; i++ {
        nextNode, ok := node.children[test_me[i]]
        if ok {
            node = nextNode
        } else {
            return false
        }
    }
    if node.val == test_me {
        return true
    }
    return false
}

func (root *TrieNode) Insert(insert_me string) bool {
    node := root
    strlen := len(insert_me)
    i := 0
    for ; i < strlen; i++ {
        nextNode, ok := node.children[insert_me[i]]
        if ok {
            node = nextNode
        } else {
            break
        }
    }

    for ; i < strlen; i++ {
        temp := NewTrieNode()
        node.children[insert_me[i]] = temp
        node = node.children[insert_me[i]]
    }
    node.val = insert_me
    return true
}

func build_tree(dict []string) *TrieNode {
    root := NewTrieNode()
    for _, value := range dict {
        root.Insert(value)
    }
    return root
}

func readLines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    fmt.Println("Loaded file!")
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}
func main() {
    lines, err := readLines("english.dict")
    if err == nil {
        fmt.Println(lines[2000])
        tree := build_tree(lines)
        fmt.Println("hello is English? ", tree.Contains("hello"))
        fmt.Println("aardvark is English? ", tree.Contains("aardvark"))
        fmt.Println("haygoolig is English? ", tree.Contains("haygoolig"))
    }
}
