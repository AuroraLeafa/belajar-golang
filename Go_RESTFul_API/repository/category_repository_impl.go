package repository

import (
	"Go_RESTFul_API/helper"
	"Go_RESTFul_API/model/domain"
	"context"
	"database/sql"
	"errors"
)

type CategoryRepositoryImpl struct {
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "insert into customer(name) values (?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(id)
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "update category set name=? where id=?"
	result, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.PanicIfError(err)
	rowCount, err := result.RowsAffected()
	helper.PanicIfError(err)
	if rowCount == 0 {
		return domain.Category{}
	}
	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	SQL := "delete from category where id=?"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	SQL := "select * from category where id=?"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)

	category := domain.Category{}
	if rows.Next() {
		return category, nil
	} else {
		return category, errors.New("Category is Not Found")
	}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "select id, name from category"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories
}
