package repository

import (
	db "Golang-Fundamental/database_golang"
	"Golang-Fundamental/entity"
	"context"
	"database/sql"
)

type CommentRepositoryImpl struct {
	ctx context.Context
}

func ConstructorCommentInterfaceImpl(ctx context.Context) *CommentRepositoryImpl {
	return &CommentRepositoryImpl{ctx: ctx}
}

type CommentInterface interface {
	Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error)
	FindById(ctx context.Context, id uint64) (entity.Comment, error)
	FindAll(ctx context.Context, comment entity.Comment) (entity.Comment, error)
}

func transaction(ctx context.Context) (*sql.Tx, error) {
	connection := db.GetConnection()
	defer connection.Close()

	tx, err := connection.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func prepareContext(ctx context.Context, querytransaction string) (*sql.Stmt, error) {
	tx, err := transaction(ctx)
	if err != nil {
		panic(err)
	}
	preparecontext, err := tx.PrepareContext(ctx, querytransaction)
	if err != nil {
		return nil, err
	}

	return preparecontext, nil
}

func (repository *CommentRepositoryImpl) Insert(ctx context.Context, comment *entity.Comment) (*entity.Comment, error) {
	querytransaction := "INSERT INTO `comments`(`email`, `comment`) VALUES(?, ?);"
	preparecontext, err := prepareContext(ctx, querytransaction)
	if err != nil {
		return nil, err
	}
	defer preparecontext.Close()
	exec, err := preparecontext.ExecContext(ctx, comment.Email, comment.Comment)
	if err != nil {
		return nil, err
	}
	id, err := exec.LastInsertId()
	if err != nil {
		return nil, err
	}
	comment.ID = uint64(id)
	return comment, nil
}

func (CommentRepositoryImpl *CommentRepositoryImpl) FindById(ctx context.Context, id uint64) (*entity.Comment, error) {
	query_find_by_id := "SELECT * FROM `comments` WHERE `id` = ? LIMIT 1;"
	comment := &entity.Comment{}
	preparecontext, err := prepareContext(ctx, query_find_by_id)
	if err != nil {
		return nil, err
	}
	defer preparecontext.Close()
	rows, err := preparecontext.QueryContext(ctx, id)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		rows.Scan(&comment.ID, &comment.Email, &comment.Comment)
	}
	defer rows.Close()
	return comment, nil
}
func (CommentRepositoryImpl *CommentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	stmt, err := prepareContext(ctx, "SELECT * FROM `comments`;")
	if err != nil {
		return nil, err
	}
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	var all_comments []entity.Comment
	for rows.Next() {
		comment := entity.Comment{}
		rows.Scan(&comment.ID, &comment.Email, &comment.Comment)
		all_comments = append(all_comments, comment)
	}
	defer rows.Close()
	defer stmt.Close()
	return all_comments, nil
}
