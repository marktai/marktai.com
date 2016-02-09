package posts_test

import (
	"fmt"
	"os"
	"posts"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	posts.Init("samplePosts")
	code := m.Run()
	os.Exit(code)
}

func TestGetPost(t *testing.T) {
	baconPost, err := posts.GetPost("Bacon Ipsum")
	if err != nil {
		t.Error(err.Error())
	}
	if baconPost == nil {
		t.Error("Expected non nil Bacon Ipsum post")
	}

	if baconPost.Title != "Bacon Ipsum" {
		t.Error("Expected Title as Bacon Ipsum")
	}
	if baconPost.Ignore != false {
		t.Error("Expected Ignore as false")
	}
	if baconPost.Subtitle != "Better than that Lorem nonsense" {
		t.Error("Expected Subtitle as Better than that Lorem nonsense")
	}
	if baconPost.Author != "Mark Tai" {
		t.Error("Expected Author as Mark Tai")
	}
	if baconPost.Image != "img/bacon-bg.jpg" {
		t.Error("Expected Image as img/bacon-bg.jpg")
	}

	timeConstant := "2006-01-02T15:04:05-07:00"
	expectedTime, _ := time.Parse(timeConstant, "2000-07-27T07:32:00-08:00")
	if !baconPost.Created.Equal(expectedTime) {
		t.Error("Expected Created as 2000-07-27T07:32:00-08:00")
	}
	expectedTime, _ = time.Parse(timeConstant, "2000-05-27T07:32:00-08:00")
	if !baconPost.Modified.Equal(expectedTime) {
		t.Error("Expected Modified as 2000-05-27T07:32:00-08:00")
	}
	if len(baconPost.Content) != 1 {
		t.Error("Expected Content to have length 1")
	} else if baconPost.Content[0] != "Here is bacon text" {
		t.Error("Expected Content[0] as Here is bacon text")
	}

	bacon2Post, err := posts.GetPost("Bacon Ipsum2")
	if err != nil {
		t.Error(err.Error())
	}
	if bacon2Post == nil {
		t.Error("Expected non nil Bacon Ipsum2 post")
	}
	if len(bacon2Post.Content) != 2 {
		t.Error("Expected Content to have length 2")
	}
	if bacon2Post.Content[1] != "And another one!" {
		t.Error("Expected Content[1] as And another one!")
	}

	_, err = posts.GetPost("ignore")
	if err == nil {
		t.Error("Expected error getting ignored post")
	}

	_, err = posts.GetPost("bad")
	if err == nil {
		t.Error("Expected error getting bad post")
	}
}

func TestGetPostList(t *testing.T) {
	postList := posts.GetPostList()

	expectedPostList := []string{"Bacon Ipsum2", "Bacon Ipsum"}

	if len(postList) != 2 {
		t.Error("Expected post list to be length 2")
	}

	for i, _ := range postList {
		if postList[i] != expectedPostList[i] {
			t.Error(fmt.Sprintf("Expected postList[%d] as %s", i, expectedPostList[i]))
		}
	}
}

// func GetPost(id string) (*Post, error) {
// 	post, ok := postMap[id]
// 	if !ok {
// 		return nil, errors.New(id + " not found")
// 	}
// 	return post, nil
// }
//
// func GetPostList() []string {
// 	posts := make(postSlice, len(postMap))
// 	i := 0
// 	for _, post := range postMap {
// 		posts[i] = post
// 		i += 1
// 	}
//
// 	sort.Stable(posts)
//
// 	postTitles := make([]string, len(posts))
// 	for i, post := range posts {
// 		postTitles[i] = post.Title
// 	}
//
// 	return postTitles
// }
//
// func Init(path string) error {
// 	err := readFolder(path)
// 	if err != nil {
// 		return err
// 	}
// 	doneChan := make(chan bool)
// 	watcher, err := fsnotify.NewWatcher()
// 	if err != nil {
// 		return err
// 	}
//
// 	go func() {
// 		<-doneChan
// 		watcher.Close()
// 	}()
//
// 	err = watcher.Add(path)
// 	if err != nil {
// 		return err
// 	}
// }
