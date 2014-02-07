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
    prefix_length, node := root.LongestPrefix(test_me)
    if prefix_length == len(test_me) && node.val == test_me {
        return true
    }
    return false
}

func (root *TrieNode) Preorder() {
    for _, child := range root.children {
        if child.val != "" {
            fmt.Println(child.val)
        }
        child.Preorder()
    }
}

func (root *TrieNode) LongestPrefix(prefix string) (int, *TrieNode) {
    node := root
    strlen := len(prefix)
    i := 0
    for ; i < strlen; i++ {
        nextNode, ok := node.children[prefix[i]]
        if ok {
            node = nextNode
        } else {
            break
        }
    }
    return i, node
}

func (root *TrieNode) WithPrefix(prefix string) {
    prefix_length, node := root.LongestPrefix(prefix)
    if prefix_length == len(prefix) {
        node.Preorder()
    }
}

func (root *TrieNode) Insert(insert_me string) {
    prefix_length, node := root.LongestPrefix(insert_me)

    for ; prefix_length < len(insert_me); prefix_length++ {
        temp := NewTrieNode()
        node.children[insert_me[prefix_length]] = temp
        node = node.children[insert_me[prefix_length]]
    }
    node.val = insert_me
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
    if err != nil {
        panic(err)
    }
    tree := build_tree(lines)
    findLongestPrefix := func(word string) {
        i, _ := tree.LongestPrefix(word)
        fmt.Printf("Longest prefix match for %s? %s\n", word, word[:i])
    }
    findWordsWithPrefix := func(word string) {
        fmt.Printf("Words that start with %s?\n", word)
        tree.WithPrefix(word)
    }
    fmt.Println("hello is English? ", tree.Contains("hello"))
    fmt.Println("aardvark is English? ", tree.Contains("aardvark"))
    fmt.Println("haygoolig is English? ", tree.Contains("haygoolig"))
    findLongestPrefix("exterosis")
    findLongestPrefix("whitney")
    findLongestPrefix("alakas")
    findWordsWithPrefix("arbitrag")
    findWordsWithPrefix("zy")
}
