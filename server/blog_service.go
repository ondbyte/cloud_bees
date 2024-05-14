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

type BlogServiceServer struct {
	posts map[string]*blog.Post
	idGen uint
}

// CreatePost implements blog.BlogServiceServer.
func (b *BlogServiceServer) CreatePost(ctx context.Context, req *blog.CreatePostRequest) (*blog.CreatePostResponse, error) {
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
func (b *BlogServiceServer) DeletePost(ctx context.Context, req *blog.DeletePostRequest) (*blog.DeletePostResponse, error) {
	_, ok := b.posts[req.PostId]
	if !ok {
		return &blog.DeletePostResponse{
			ErrorMessage: fmt.Sprintf("post with "),
		}
	}
	delete(b.posts, req.PostId)
}

// ReadPost implements blog.BlogServiceServer.
func (b *BlogServiceServer) ReadPost(context.Context, *blog.ReadPostRequest) (*blog.ReadPostResponse, error) {
	panic("unimplemented")
}

// UpdatePost implements blog.BlogServiceServer.
func (b *BlogServiceServer) UpdatePost(context.Context, *blog.UpdatePostRequest) (*blog.UpdatePostResponse, error) {
	panic("unimplemented")
}

// mustEmbedUnimplementedBlogServiceServer implements blog.BlogServiceServer.
func (b *BlogServiceServer) mustEmbedUnimplementedBlogServiceServer() {
	panic("unimplemented")
}
