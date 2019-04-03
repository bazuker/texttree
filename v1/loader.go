package texttree

// AUTHOR: Daniil Furmanov
// LICENSE: MIT
// DESCRIPTION: TextTree is a file buffer that stores files content
// 				in memory and allow access to it by path.
//				It is useful for working with localization trees.

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

const DefaultMaxFileSize = 16 << 10 // 16 KB

type TextTree struct {
	cache    map[string]*Entity
	basePath string
}

type Entity struct {
	Content  string
	Filename string
}

// Creates a new text tree by loading all the files from a specified directory
// that are smaller than maximum file size
func NewTextTree(path string, maxFileSize int64) (*TextTree, error) {
	cache := make(map[string]*Entity)
	// recursively walk through the path
	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// check the file size
			if !info.IsDir() && info.Size() < maxFileSize {
				data, err := ioutil.ReadFile(path)
				if err != nil {
					return err
				}
				newEntity := &Entity{Content: string(data), Filename: info.Name()}
				dir, filename := filepath.Split(path)
				key := dir + filename[0:len(filename)-len(filepath.Ext(filename))]
				if entity, ok := cache[key]; ok {
					// if entity already exists, keep the key extension
					delete(cache, key)
					cache[path] = newEntity
					cache[dir+entity.Filename] = entity
				} else {
					cache[key] = newEntity
				}
			}
			return nil
		})
	if err != nil {
		return nil, err
	}
	// create a text tree object with buffered data
	return &TextTree{
		cache:    cache,
		basePath: path,
	}, nil
}

// Creates an array of file paths
func (tt *TextTree) Entities() (entities []string) {
	entities = make([]string, len(tt.cache))
	i := 0
	for key := range tt.cache {
		entities[i] = key
		i++
	}
	return
}

// Returns an entity
func (tt *TextTree) Get(path string) *Entity {
	return tt.cache[path]
}

// Returns an entity if it exists
func (tt *TextTree) GetIfExists(path string) (*Entity, bool) {
	e, ok := tt.cache[path]
	return e, ok
}

// Returns an entity's content
func (tt *TextTree) GetString(path string) string {
	return tt.cache[path].Content
}

// Returns an entity's content if it exists
func (tt *TextTree) GetStringIfExists(path string) (string, bool) {
	e, ok := tt.cache[path]
	return e.Content, ok
}

// Returns the path of a loaded directory
func (tt *TextTree) GetBasePath() string {
	return tt.basePath
}
