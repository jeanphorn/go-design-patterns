package main

import (
	"fmt"
)

type Facade struct {
	M Music
	V Video
	C Count
}

func (this *Facade) GetRecommandVideos() error {
	this.V.GetVideos()
	this.C.GetCountByID(111)

	return nil
}

type Music struct {
}

func (this *Music) GetMusic() error {
	fmt.Println("get music material")
	// logic code here
	return nil
}

type Video struct {
	vid int64
}

func (this *Video) GetVideos() error {
	fmt.Println("get videos1")
	return nil
}

type Count struct {
	PraiseCnt  int64 //点赞数
	CommentCnt int64 //评论数
	CollectCnt int64 //收藏数
}

func (this *Count) GetCountByID(id int64) (*Count, error) {
	fmt.Println("get video counts")
	return this, nil
}

func main() {
	f := &Facade{}
	f.GetRecommandVideos()
}
