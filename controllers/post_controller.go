package controllers

import (
	"github.com/gin-gonic/gin"
	"go-ginapp/data/request"
	"go-ginapp/services/post"
	"go-ginapp/utils"
	"net/http"
	"strconv"
)

type PostController interface {
	GetAllPosts(c *gin.Context)
	GetPostByID(c *gin.Context)
	GetPostsByUser(c *gin.Context)
	CreatePost(c *gin.Context)
	UpdatePost(c *gin.Context)
	DeletePost(c *gin.Context)
}

type postController struct {
	postService post.PostService
}

func NewPostController(postService post.PostService) PostController {
	return &postController{postService}
}

func (p *postController) GetAllPosts(c *gin.Context) {
	response, err := p.postService.GetAllPosts()
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Get all post failed", err)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Get all post success", response)
}

func (p *postController) GetPostByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("postId"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid post id", err)
		return
	}
	response, err := p.postService.GetPostById(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Get posts failed", err)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Get posts success", response)
}

func (p *postController) GetPostsByUser(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Param("userId"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid user id", err)
		return
	}
	response, err := p.postService.GetPostByAuthor(uint(userId))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Get posts with author failed", err)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "get post with author success", response)
}

func (p *postController) CreatePost(c *gin.Context) {
	userId := c.GetUint("userId")
	var createPostRq request.CreatePostRequest
	if err := c.ShouldBindJSON(&createPostRq); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request", err)
		return
	}
	response, err := p.postService.CreatePost(createPostRq, userId)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Create post failed", err)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Create posts success", response)
}

func (p *postController) UpdatePost(c *gin.Context) {
	postId, err := strconv.ParseUint("postId", 10, 32)
	userId := c.GetUint("userId")
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid post id", err)
		return
	}
	var updatePostRq request.UpdatePostRequest
	if err := c.ShouldBindJSON(&updatePostRq); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request", err)
		return
	}
	response, err := p.postService.UpdatePost(updatePostRq, uint(postId), userId)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Update post failed", err)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Update post success", response)
}

func (p *postController) DeletePost(c *gin.Context) {
	postId, err := strconv.ParseUint("postId", 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid post id", err)
		return
	}
	err = p.postService.DeletePost(uint(postId))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Delete post failed", err)
		return
	}
}
