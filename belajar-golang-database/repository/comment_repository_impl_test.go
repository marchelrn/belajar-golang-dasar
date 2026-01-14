package repository

import (
	"belajar-golang-database"
	"belajar-golang-database/entity"
	"context"
	"fmt"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepositoryImpl(belajar_golang_database.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "repository@test.com",
		Comment: "Test Repository",
	}
	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(result)
}

func TestCommentRepositoryImpl_FindById(t *testing.T) {
	commentRepository := NewCommentRepositoryImpl(belajar_golang_database.GetConnection())
	ctx := context.Background()
	comment, err := commentRepository.FindById(ctx, 68)
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(comment)
}

func TestCommentRepositoryImpl_FindAll(t *testing.T) {
	commentRepository := NewCommentRepositoryImpl(belajar_golang_database.GetConnection())
	ctx := context.Background()
	comments, err := commentRepository.FindAll(ctx)
	if err != nil {
		t.Fatal(err.Error())
	}
	for _, comment := range comments {
		fmt.Println(comment)
	}
}
