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
	Content   string
	Filename  string
	Directory bool
}

// Creates a new text tree by loading all the files from a specified directory
// that are smaller than maximum file size
func NewTextTree(dirPath string, maxFileSize int64) (*TextTree, error) {
	cache := make(map[string]*Entity)
	// remove last back slash
	dirPathLen := len(dirPath)
	if dirPath[dirPathLen-1] == '/' {
		dirPathLen--
		dirPath = dirPath[:dirPathLen]
	}
	// recursively walk through the dirPath
	err := filepath.Walk(dirPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// skip the root directory
			if len(path) <= dirPathLen {
				return nil
			}
			shortPath := path[dirPathLen+1:]
			// check if it is a directory
			if info.IsDir() {
				key := shortPath
				if info.Name()[len(info.Name())-1] != '/' {
					key += "/"
				}
				cache[key] = &Entity{
					Content:   info.Name(),
					Filename:  info.Name(),
					Directory: true,
				}
				return nil
			}
			// check the file size
			if info.Size() < maxFileSize {
				data, err := ioutil.ReadFile(path)
				if err != nil {
					return err
				}
				newEntity := &Entity{
					Content:   string(data),
					Filename:  info.Name(),
					Directory: false,
				}
				dir, filename := filepath.Split(shortPath)
				key := dir + filename[0:len(filename)-len(filepath.Ext(filename))]
				if entity, ok := cache[key]; ok {
					// if entity already exists, keep the key extension
					delete(cache, key)
					cache[dir+info.Name()] = newEntity
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
		basePath: dirPath,
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

// Verifies if subdirectory exists
func (tt *TextTree) SubExists(sub string) bool {
	_, exists := tt.cache[sub+"/"]
	return exists
}

// Gets an entity with a sub path
func (tt *TextTree) GetSub(sub, path string) *Entity {
	return tt.cache[sub+"/"+path]
}

// Gets an entity
func (tt *TextTree) Get(path string) *Entity {
	return tt.cache[path]
}

// Gets an entity if it exists
func (tt *TextTree) GetIfExists(path string) (*Entity, bool) {
	e, ok := tt.cache[path]
	return e, ok
}

// Gets an entity with a sub path if it exists
func (tt *TextTree) GetSubIfExists(sub, path string) (*Entity, bool) {
	e, ok := tt.cache[sub+"/"+path]
	return e, ok
}

// Gets an entity's content with a sub path
func (tt *TextTree) GetStringSub(sub, path string) string {
	return tt.cache[sub+"/"+path].Content
}

// Gets an entity's content
func (tt *TextTree) GetString(path string) string {
	return tt.cache[path].Content
}

// Gets an entity's content if it exists
func (tt *TextTree) GetStringIfExists(path string) (result string, ok bool) {
	e, ok := tt.cache[path]
	if ok {
		result = e.Content
	}
	return
}

// Gets an entity's content with a sub path if it exists
func (tt *TextTree) GetStringSubIfExists(sub, path string) (result string, ok bool) {
	e, ok := tt.cache[sub+"/"+path]
	if ok {
		result = e.Content
	}
	return
}

// Gets the path of a loaded directory
func (tt *TextTree) GetBasePath() string {
	return tt.basePath
}
