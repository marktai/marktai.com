package posts

import (
	"errors"
	"github.com/go-fsnotify/fsnotify"
	"github.com/naoina/toml"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

// Info of a post
type PostInfo struct {
	Title    string
	Ignore   bool
	Subtitle string
	Author   string
	Created  time.Time
	Modified time.Time
}

// Post (PostInfo and the actual text content)
type Post struct {
	PostInfo
	Image   string
	Content []string
}

// Returns just info of a post
func (p Post) Info() PostInfo {
	return p.PostInfo
}

// for sorting
type postSlice []*Post

func (p postSlice) Len() int {
	return len(p)
}

func (p postSlice) Less(i, j int) bool {
	return p[i].Created.After(p[j].Created)
}

func (p postSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

var postMap map[string]*Post

// Gets post by title
func GetPost(id string) (*Post, error) {
	post, ok := postMap[id]
	if !ok {
		return nil, errors.New(id + " not found")
	}
	return post, nil
}

// Returns a list of all the posts sorted by time modified
func GetPostList() []string {
	posts := make(postSlice, len(postMap))
	i := 0
	for _, post := range postMap {
		posts[i] = post
		i += 1
	}

	sort.Stable(posts)

	postTitles := make([]string, len(posts))
	for i, post := range posts {
		postTitles[i] = post.Title
	}

	return postTitles
}

// Reads the folder and parses all the posts
// Also sets up a watcher so the posts will automatically
//   be updated if any of the files change
func Init(path string) error {
	err := readFolder(path)
	if err != nil {
		return err
	}
	doneChan := make(chan bool)
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}

	go func() {
		<-doneChan
		watcher.Close()
	}()

	err = watcher.Add(path)
	if err != nil {
		return err
	}
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if strings.HasSuffix(event.Name, ".toml") {
					log.Println("event:", event)
					readFolder(path)
					log.Println("Updated posts")
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	return nil
}

// Reads the folder and parses all the files into posts
func readFolder(path string) error {

	tempPostMap := make(map[string]*Post)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println("Can't read at %s", path)
		return err
	}

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".toml") || file.IsDir() || file.Name()[0] == '.' {
			continue
		}
		post, err := readPostFile(path + "/" + file.Name())
		if err != nil {
			log.Println(err)
			continue
		}
		if post.Ignore {
			continue
		}
		if _, ok := tempPostMap[post.Title]; !ok {
			tempPostMap[post.Title] = post
		} else {
			log.Println(post.Title + " encountered twice")
		}

	}

	postMap = tempPostMap

	return nil
}

// Reads a file and returns a post
func readPostFile(path string) (*Post, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	var post Post
	if err := toml.Unmarshal(buf, &post); err != nil {
		return nil, err
	}

	if post.Title == "" {
		return nil, errors.New("No Title specified")
	}
	return &post, nil
}
