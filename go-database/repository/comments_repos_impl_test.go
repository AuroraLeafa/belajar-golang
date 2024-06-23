package repository

import (
	"context"
	"fmt"
	godatabase "go-database"
	"go-database/entity"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	CommentRepository := NewCommentRepository(godatabase.GetConnection())
	ctx := context.Background()
	comment := entity.Comment{
		Email:   "repos@test.com",
		Comment: "TESTING 123",
	}
	res, err := CommentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println("Result:", res)
}

func TestCommentFindById(t *testing.T) {
	CommentRepository := NewCommentRepository(godatabase.GetConnection())
	ctx := context.Background()
	res, err := CommentRepository.FindById(ctx, 11)
	if err != nil {
		panic(err)
	}

	fmt.Println("Result:", res)
}

func TestCommentFindAll(t *testing.T) {
	CommentRepository := NewCommentRepository(godatabase.GetConnection())
	ctx := context.Background()
	res, err := CommentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("Result:", res)
}
