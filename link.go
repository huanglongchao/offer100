/**
go 实现链表

及fmt.Printf 占位符
%v	值的默认格式。当打印结构体时，“加号”标记（%+v）会添加字段名
%#v　相应值的Go语法表示
%T	相应值的类型的Go语法表示
%%	字面上的百分号，并非值的占位符
 */
package main

import (
	"errors"
	"fmt"
)

type Post struct{
	body string
	publishDate int64 //Unix timestamp
	next *Post
}

type Feed struct {
	length int //we'll use it later
	start *Post
}

func (f *Feed) Append(newPost *Post)  {
	if f.length == 0 {
		f.start = newPost
	} else {
		currentPost := f.start
		for currentPost.next != nil {
			currentPost = currentPost.next
		}
		currentPost.next = newPost
	}
	f.length++
}

func (f *Feed) Remove(publicDate int64){
	if f.length == 0{
		panic(errors.New("Feed is empty"))
	}
	var previousPost *Post
	currentPost := f.start

	for currentPost.publishDate != publicDate {
		if currentPost.next == nil{
			panic(errors.New("No such Post found."))
		}
		previousPost = currentPost
		currentPost = currentPost.next
	}
	previousPost.next = currentPost.next
	f.length--
}

func (f *Feed) Insert(newPost *Post){
	if f.length == 0{
		f.start = newPost
	} else {
		var previousPost *Post
		currentPost := f.start

		for currentPost.publishDate < newPost.publishDate {
			previousPost = currentPost
			currentPost = previousPost.next
		}
		previousPost.next = newPost
		newPost.next = currentPost
	}
	f.length++
}

func main(){
	f := &Feed{}
	p1 := Post{
		body: "Lorem ipsum",
	}
	f.Append(&p1)

	fmt.Printf("Length: %v\n",f.length)
	fmt.Printf("First: %+v\n", f.start)

	p2 := Post{
		body: "Dolor sit amet",
	}
	f.Append(&p2)

	fmt.Printf("Length: %#v\n",f.length)
	fmt.Printf("First: %#v\n", f.start)
	fmt.Printf("Second: %Tv\n",f.start.next)
}