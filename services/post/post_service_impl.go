package post

import (
	"errors"
	"go-ginapp/data/request"
	"go-ginapp/data/response"
	"go-ginapp/models"
	"go-ginapp/repositories"
	"go-ginapp/utils"
)

type PostServiceImpl struct {
	postRepo repositories.PostRepository
	userRepo repositories.UserRepository
}

func NewPostServiceImpl(postRepo repositories.PostRepository,
	userRepo repositories.UserRepository,
) PostService {
	return &PostServiceImpl{postRepo: postRepo, userRepo: userRepo}
}

func (p *PostServiceImpl) GetAllPosts() ([]response.PostResponse, error) {
	posts, err := p.postRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return ToListPostResponse(posts)
}

func (p *PostServiceImpl) GetPostById(id uint) (response.PostResponse, error) {
	post, err := p.postRepo.FindById(id)
	if err != nil {
		return response.PostResponse{}, err
	}
	postResponse, err := utils.ToPostResponse(post)
	if err != nil {
		return response.PostResponse{}, err
	}
	return postResponse, nil
}

func (p *PostServiceImpl) GetPostByAuthor(userId uint) ([]response.PostResponse, error) {
	posts, err := p.postRepo.FindByUser(userId)
	if err != nil {
		return nil, err
	}
	return ToListPostResponse(posts)
}

func (p *PostServiceImpl) CreatePost(request request.CreatePostRequest, userId uint) (response.PostResponse, error) {
	post := models.Post{
		Title:   request.Title,
		Content: request.Content,
		UserId:  userId,
	}
	post, err := p.postRepo.Create(post)
	if err != nil {
		return response.PostResponse{}, err
	}
	postResponse, err := utils.ToPostResponse(post)
	if err != nil {
		return response.PostResponse{}, err
	}
	return postResponse, nil
}

func (p *PostServiceImpl) UpdatePost(postRequest request.UpdatePostRequest, postId uint, userId uint) (response.PostResponse, error) {
	post, err := p.postRepo.FindById(postId)
	if err != nil {
		return response.PostResponse{}, err
	}
	if post.UserId != userId {
		return response.PostResponse{}, errors.New("you does not have permission to change this post")
	}
	post.Title = postRequest.Title
	post.Content = postRequest.Content

	post, err = p.postRepo.Update(post)
	if err != nil {
		return response.PostResponse{}, err
	}
	postResponse, err := utils.ToPostResponse(post)
	if err != nil {
		return response.PostResponse{}, err
	}
	return postResponse, nil
}

func (p *PostServiceImpl) DeletePost(id uint) error {
	return p.postRepo.Delete(id)
}

func ToListPostResponse(posts []models.Post) ([]response.PostResponse, error) {
	var postResponses []response.PostResponse
	for _, post := range posts {
		postResponse, err := utils.ToPostResponse(post)
		if err != nil {
			return nil, err
		}
		postResponses = append(postResponses, postResponse)
	}
	return postResponses, nil
}
