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

type PostInfo struct {
	Title    string
	Subtitle string
	Author   string
	Created  time.Time
	Modified time.Time
}

type Post struct {
	Title    string
	Subtitle string
	Author   string
	Created  time.Time
	Modified time.Time
	Content  []string
}

func (p Post) Info() PostInfo {
	return PostInfo{p.Title, p.Subtitle, p.Author, p.Created, p.Modified}
}

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

func GetPost(id string) (*Post, error) {
	post, ok := postMap[id]
	if !ok {
		return nil, errors.New(id + " not found")
	}
	return post, nil
}

func GetPostList() []string {
	posts := make(postSlice, len(postMap))
	i := 0
	for _, post := range postMap {
		posts[i] = post
		i += 1
	}

	sort.Sort(posts)

	postTitles := make([]string, len(posts))
	for i, post := range posts {
		postTitles[i] = post.Title
	}

	return postTitles
}

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

func readFolder(path string) error {

	tempPostMap := make(map[string]*Post)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println("Can't read at %s", path)
		return err
	}

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".toml") || file.IsDir() {
			continue
		}
		post, err := readPostFile(path + "/" + file.Name())
		if err != nil {
			log.Println(err)
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
