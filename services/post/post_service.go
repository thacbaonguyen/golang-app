package post

import (
	"go-ginapp/data/request"
	"go-ginapp/data/response"
)

type PostService interface {
	GetAllPosts() ([]response.PostResponse, error)
	GetPostById(id uint) (response.PostResponse, error)
	GetPostByAuthor(userId uint) ([]response.PostResponse, error)
	CreatePost(request request.CreatePostRequest, userId uint) (response.PostResponse, error)
	UpdatePost(postRequest request.UpdatePostRequest, postId uint) (response.PostResponse, error)
	DeletePost(id uint) error
}
