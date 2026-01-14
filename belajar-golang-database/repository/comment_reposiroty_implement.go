package repository

import (
	"belajar-golang-database/entity"
	"context"
	"database/sql"
	"errors"
	"strconv"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepositoryImpl(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (repository *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	query := "INSERT INTO comments(email, comment) VALUES($1, $2) RETURNING id"
	var resultId int32
	err := repository.DB.QueryRowContext(ctx, query, comment.Email, comment.Comment).Scan(&resultId)
	if err != nil {
		return comment, err
	}
	comment.Id = resultId
	return comment, nil
}

func (repository *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	query := "SELECT id, email, comment FROM comments WHERE id = $1 LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, query, id)
	comment := entity.Comment{}
	if err != nil {
		return comment, err
	}
	defer rows.Close()
	if rows.Next() {
		//ada
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		//tidak ada
		return comment, errors.New("id " + strconv.Itoa(int(id)) + " does not exist")
	}
}

func (repository *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	query := "SELECT * FROM comments"
	rows, err := repository.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []entity.Comment
	for rows.Next() {
		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}
	return comments, nil
}
