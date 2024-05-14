package server

import (
	"context"
	"fmt"

	"github.com/ondbyte/cloud_bees/blog"
)

/* type Post struct {
	ID              uint
	Title           string
	Content         string
	Author          string
	PublicationDate time.Time
	Tags            []string
} */

type blogServiceServer struct {
	posts map[string]*blog.Post
	idGen uint
	blog.UnimplementedBlogServiceServer
}

func NewBlogServiceServer() *blogServiceServer {
	return &blogServiceServer{
		posts: make(map[string]*blog.Post),
		idGen: 0,
	}
}

// CreatePost implements blog.BlogServiceServer.
func (b *blogServiceServer) CreatePost(ctx context.Context, req *blog.CreatePostRequest) (*blog.CreatePostResponse, error) {
	b.idGen++
	newPost := &blog.Post{
		PostId:          fmt.Sprint(b.idGen),
		Title:           req.Title,
		Content:         req.Content,
		Author:          req.Author,
		PublicationDate: req.PublicationDate,
		Tags:            req.Tags,
	}
	b.posts[newPost.PostId] = newPost
	return &blog.CreatePostResponse{
		Post: newPost,
	}, nil
}

// DeletePost implements blog.BlogServiceServer.
func (b *blogServiceServer) DeletePost(ctx context.Context, req *blog.DeletePostRequest) (*blog.DeletePostResponse, error) {
	_, ok := b.posts[req.PostId]
	if !ok {
		return &blog.DeletePostResponse{
			ErrorMessage: fmt.Sprintf("post with id '%v' doesnt exist", req.PostId),
		}, nil
	}
	delete(b.posts, req.PostId)
	return &blog.DeletePostResponse{
		Success: true,
	}, nil
}

// ReadPost implements blog.BlogServiceServer.
func (b *blogServiceServer) ReadPost(ctx context.Context, req *blog.ReadPostRequest) (*blog.ReadPostResponse, error) {
	post, ok := b.posts[req.PostId]
	if !ok {
		return &blog.ReadPostResponse{
			ErrorMessage: fmt.Sprintf("post with id '%v' doesnt exist", req.PostId),
		}, nil
	}
	return &blog.ReadPostResponse{
		Post: post,
	}, nil
}

// UpdatePost implements blog.BlogServiceServer.
func (b *blogServiceServer) UpdatePost(ctx context.Context, req *blog.UpdatePostRequest) (*blog.UpdatePostResponse, error) {
	post, ok := b.posts[req.PostId]
	if !ok {
		return &blog.UpdatePostResponse{
			ErrorMessage: fmt.Sprintf("post with id '%v' doesnt exist", req.PostId),
		}, nil
	}
	b.posts[post.PostId] = &blog.Post{
		PostId:          post.PostId,
		Title:           req.Title,
		Content:         req.Content,
		Author:          req.Author,
		PublicationDate: post.PublicationDate,
		Tags:            req.Tags,
	}
	return &blog.UpdatePostResponse{
		Post: b.posts[post.PostId],
	}, nil
}
