package domain_tree

import (
	"bufio"
	"fmt"
	"github.com/YafimK/go-succinct-data-structure-trie/succinct_tree"
	"github.com/YafimK/go-succinct-data-structure-trie/tree_proto"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func Reverse(in string) string {
	var sb strings.Builder
	runes := []rune(in)
	for i := len(runes) - 1; 0 <= i; i-- {
		sb.WriteRune(runes[i])
	}
	return sb.String()
}

func ReadWordsFromFile(sourceFilePath string) ([]string, error) {
	file, err := os.Open(sourceFilePath)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		word := scanner.Text()
		if word = strings.TrimSpace(word); len(word) == 0 {
			continue
		}
		word = strings.ToLower(word)
		word = Reverse(word)
		lines = append(lines, word)
	}
	return lines, nil
}

func ConstructTree(allowedChars string, words []string) (*succinct_tree.Trie, error) {
	succinct_tree.SetAllowedCharacters(allowedChars)
	te := succinct_tree.Trie{}
	te.Init()

	for _, word := range words {
		te.Insert(word)
	}

	return &te, nil
}

func SerializeTree(allowedChars, name, treeData, rankData string, nodeCount uint64) ([]byte, error) {
	e := &tree_proto.TreeEntry{
		AllowedChars: allowedChars,
		WordListName: name,
		Tree:         treeData,
		Rank:         rankData,
		NodeCount:    nodeCount,
	}

	data, err := proto.Marshal(e)
	if err != nil {
		err = fmt.Errorf("marshaling error: %w", err)
	}
	return data, err
}

func WriteNewDomainTree(allowedChars, name, sourceFile, outputPath string) {

	words, err := ReadWordsFromFile(sourceFile)
	if err != nil {
		log.Fatalf("failed loading loading file %q: %v", sourceFile, err)
	}
	tree, err := ConstructTree(allowedChars, words)
	if err != nil {
		log.Fatalf("failed constructing tree from file %q: %v", sourceFile, err)
	}
	teData := tree.Encode()
	rd := succinct_tree.CreateRankDirectory(teData, tree.GetNodeCount()*2+1, succinct_tree.L1, succinct_tree.L2)

	serializedTree, err := SerializeTree(allowedChars, name, teData, rd.GetData(), uint64(tree.GetNodeCount()))
	if err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile(outputPath, serializedTree, 0644); err != nil {
		log.Fatalln("Failed to write tree:", err)
	}
}

func LoadTree(sourceFile []byte) (*succinct_tree.FrozenTrie, error) {
	treeEntry := &tree_proto.TreeEntry{}
	err := proto.Unmarshal(sourceFile, treeEntry)
	if err != nil {
		return nil, fmt.Errorf("tree unmarshaling error: %w", err)
	}
	succinct_tree.SetAllowedCharacters(treeEntry.AllowedChars)
	ft := succinct_tree.FrozenTrie{}
	ft.Init(treeEntry.GetTree(), treeEntry.GetRank(), uint(treeEntry.GetNodeCount()))
	return &ft, nil
}
