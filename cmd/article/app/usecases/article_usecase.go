package usecase

import (
	"context"
	"database/sql"
	"time"

	"github.com/azcov/sagara_crud/cmd/article/app/domain"
	appConfig "github.com/azcov/sagara_crud/cmd/article/config"
	"github.com/google/uuid"
)

type articleUsecase struct {
	repository domain.ArticleRepository
	config     *appConfig.Config
}

func NewUsecase(cfg *appConfig.Config, repo domain.ArticleRepository) domain.ArticleUsecase {
	return &articleUsecase{
		repository: repo,
		config:     cfg,
	}
}

func (u *articleUsecase) CreateArticle(ctx context.Context, req *domain.CreateArticleRequest, userID uuid.UUID) error {
	params := &domain.InsertArticleParams{
		Title:       req.Title,
		Description: req.Description,
		CreatedAt:   time.Now().Unix(),
		CreatedBy:   userID,
	}
	if err := u.repository.InsertArticle(ctx, params); err != nil {
		return err
	}
	return nil
}

func (u *articleUsecase) GetArticle(ctx context.Context, req *domain.GetArticleRequest) (*domain.GetArticleResponse, error) {
	data, err := u.repository.GetArticleByArticleID(ctx, req.ArticleID)
	if err != nil {
		return nil, err
	}

	res := &domain.GetArticleResponse{
		ArticleID:   data.ID.String(),
		Title:       data.Title,
		Description: data.Description,
		CreatedAt:   data.CreatedAt,
		CreatedBy:   data.CreatedBy.String(),
	}

	if data.UpdatedAt.Valid {
		res.UpdatedAt = &data.UpdatedAt.Int64
	}
	if data.UpdatedBy.Valid {
		updatedBy := data.UpdatedBy.UUID.String()
		res.UpdatedBy = &updatedBy
	}
	return res, nil
}

func (u *articleUsecase) GetArticles(ctx context.Context, req *domain.GetArticlesRequest, userID uuid.UUID) ([]*domain.GetArticleResponse, error) {
	params := &domain.GetArticlesByUserIDParams{
		UserID:   userID,
		Search:   req.Search,
		OrderBy:  req.SortBy,
		SqlLimit: req.Limit,
	}
	data, err := u.repository.GetArticlesByUserID(ctx, params, req.Page)
	if err != nil {
		return nil, err
	}

	res := []*domain.GetArticleResponse{}
	for i := range data {
		article := &domain.GetArticleResponse{
			ArticleID:   data[i].ID.String(),
			Title:       data[i].Title,
			Description: data[i].Description,
			CreatedAt:   data[i].CreatedAt,
			CreatedBy:   data[i].CreatedBy.String(),
		}

		if data[i].UpdatedAt.Valid {
			article.UpdatedAt = &data[i].UpdatedAt.Int64
		}
		if data[i].UpdatedBy.Valid {
			updatedBy := data[i].UpdatedBy.UUID.String()
			article.UpdatedBy = &updatedBy
		}
		res = append(res, article)
	}

	return res, nil
}

func (u *articleUsecase) UpdateArticle(ctx context.Context, req *domain.UpdateArticleRequest, userID uuid.UUID) error {
	params := &domain.UpdateArticleParams{
		Title:       req.Title,
		Description: req.Description,
		UpdatedAt:   sql.NullInt64{Int64: time.Now().Unix(), Valid: true},
		UpdatedBy:   uuid.NullUUID{UUID: userID, Valid: true},
		ArticleID:   req.ArticleID,
	}
	if err := u.repository.UpdateArticle(ctx, params); err != nil {
		return err
	}
	return nil
}

func (u *articleUsecase) DeleteArticle(ctx context.Context, req *domain.DeleteArticleRequest, userID uuid.UUID) error {
	params := &domain.DeleteArticleParams{
		DeletedAt: sql.NullInt64{Int64: time.Now().Unix(), Valid: true},
		DeletedBy: uuid.NullUUID{UUID: userID, Valid: true},
		ArticleID: req.ArticleID,
	}
	if err := u.repository.DeleteArticle(ctx, params); err != nil {
		return err
	}
	return nil
}

func (u *articleUsecase) ForceDeleteArticle(ctx context.Context, req *domain.ForceDeleteArticleRequest) error {
	if err := u.repository.ForceDeleteArticle(ctx, req.ArticleID); err != nil {
		return err
	}
	return nil
}
