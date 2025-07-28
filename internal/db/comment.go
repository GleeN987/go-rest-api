package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/GleeN987/go-rest-api/internal/comment"
	"github.com/google/uuid"
)

type CommentRow struct {
	ID     string
	Slug   sql.NullString
	Body   sql.NullString
	Author sql.NullString
}

func commentRowToComment(c CommentRow) comment.Comment {
	return comment.Comment{
		ID:     c.ID,
		Slug:   c.Slug.String,
		Body:   c.Body.String,
		Author: c.Author.String,
	}
}

func (d *Database) GetComment(ctx context.Context, id string) (comment.Comment, error) {
	var commentRow CommentRow
	row := d.Client.DB.QueryRowContext(
		ctx,
		`SELECT id, slug, body, author FROM comments WHERE id=$1`,
		id,
	)
	err := row.Scan(&commentRow.ID, &commentRow.Slug, &commentRow.Body, &commentRow.Author)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("error when finding comment by id")
	}
	return commentRowToComment(commentRow), nil
}

func (d *Database) PostComment(ctx context.Context, cmt comment.Comment) (comment.Comment, error) {
	cmt.ID = uuid.New().String()
	postRow := CommentRow{
		ID:     cmt.ID,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Body:   sql.NullString{String: cmt.Body, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
	}

	_, err := d.Client.NamedExecContext(
		ctx,
		`INSERT INTO comments (id, slug, body, author)
		VALUES (:id, :slug, :body, :author)`,
		postRow,
	)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("error inserting comment into db: %w", err)
	}

	return cmt, nil
}

func (d *Database) DeleteComment(ctx context.Context, id string) error {
	_, err := d.Client.ExecContext(
		ctx,
		`DELETE FROM comments WHERE id = $1`,
		id,
	)
	if err != nil {
		return fmt.Errorf("error deleting comment: %w", err)
	}
	return nil
}

func (d *Database) UpdateComment(ctx context.Context, id string, cmt comment.Comment) (comment.Comment, error) {
	updateRow := CommentRow{
		ID:     id,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Body:   sql.NullString{String: cmt.Body, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
	}

	_, err := d.Client.NamedExecContext(
		ctx,
		`UPDATE comments SET slug = :slug, body = :body, author = :author WHERE id = :id`,
		updateRow,
	)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("error updating the comment: %w", err)
	}
	return cmt, nil
}
