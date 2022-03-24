package internal

import "gorm.io/gorm"

type RepositoryInterface interface {
	GetPosts() []Post
	CreatePost(req UpsertPostRequest) (*Post, error)
	FindPost(id uint) (*Post, error)
	UpdatePost(post *Post, req UpsertPostRequest) error
	DeletePost(post *Post) error
	GetComments(post *Post) []Comment
	CreateComment(post *Post, req UpsertCommentRequest) (*Comment, error)
	FindComment(post *Post, id uint) (*Comment, error)
	UpdateComment(comment *Comment, req UpsertCommentRequest) error
	DeleteComment(comment *Comment) error
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) RepositoryInterface {
	return &Repository{db}
}

func (r *Repository) GetPosts() []Post {
	var p []Post
	r.db.Find(&p)

	return p
}

func (r *Repository) CreatePost(req UpsertPostRequest) (*Post, error) {
	p := Post{
		Title:  req.Title,
		Detail: req.Detail,
	}

	if result := r.db.Create(&p); result.Error != nil {
		return nil, result.Error
	}

	return &p, nil
}

func (r *Repository) FindPost(id uint) (*Post, error) {
	var p Post
	if result := r.db.First(&p, id); result.Error != nil {
		return nil, result.Error
	}

	return &p, nil
}

func (r *Repository) UpdatePost(post *Post, req UpsertPostRequest) error {
	n := Post{
		Title:  req.Title,
		Detail: req.Detail,
	}

	result := r.db.Model(post).Updates(n)
	return result.Error
}

func (r *Repository) DeletePost(post *Post) error {
	result := r.db.Delete(&Post{}, post.ID)

	return result.Error
}

func (r *Repository) GetComments(post *Post) []Comment {
	return post.Comments
}

func (r *Repository) CreateComment(post *Post, req UpsertCommentRequest) (*Comment, error) {
	c := Comment{
		PostID:  post.ID,
		Message: req.Message,
	}

	if result := r.db.Create(&c); result.Error != nil {
		return nil, result.Error
	}

	return &c, nil
}

func (r *Repository) FindComment(post *Post, id uint) (*Comment, error) {
	var c Comment
	if result := r.db.Where("post_id", post.ID).First(&c, id); result.Error != nil {
		return nil, result.Error
	}

	return &c, nil
}

func (r *Repository) UpdateComment(comment *Comment, req UpsertCommentRequest) error {
	n := Comment{
		Message: req.Message,
	}

	result := r.db.Model(comment).Updates(n)
	return result.Error
}

func (r *Repository) DeleteComment(comment *Comment) error {
	result := r.db.Delete(&Comment{}, comment.ID)

	return result.Error
}
