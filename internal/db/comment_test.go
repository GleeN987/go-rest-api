//go:build integration
// +build integration

package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/GleeN987/go-rest-api/internal/comment"
	"github.com/stretchr/testify/assert"
)

func TestCommentDatabase(t *testing.T) {
	t.Run("test post comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)

		cmt, err := db.PostComment(context.Background(), comment.Comment{
			Slug:   "slug",
			Author: "author",
			Body:   "body",
		})
		assert.NoError(t, err)

		createdCmt, err := db.GetComment(context.Background(), cmt.ID)
		assert.NoError(t, err)
		assert.Equal(t, "slug", createdCmt.Slug, "comment field should be equal to slug")
		fmt.Println("testing creating the comment")
	})

	t.Run("test get comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)

		cmt, err := db.PostComment(context.Background(), comment.Comment{
			Slug:   "slugget",
			Body:   "bodyget",
			Author: "authorget",
		})
		assert.NoError(t, err)

		getCmt, err := db.GetComment(context.Background(), cmt.ID)
		assert.NoError(t, err)

		assert.Equal(t, cmt, getCmt, "comments should be equal")
	})

	t.Run("test delete comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)

		cmt, err := db.PostComment(context.Background(), comment.Comment{
			Slug:   "slugdelete",
			Author: "authordelete",
			Body:   "bodydelete",
		})
		assert.NoError(t, err)

		err = db.DeleteComment(context.Background(), cmt.ID)
		assert.NoError(t, err)

		deletedCmt, err := db.GetComment(context.Background(), cmt.ID)
		assert.Error(t, err)
		assert.Equal(t, deletedCmt, comment.Comment{}, "comment should be empty")
	})

	t.Run("test update comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)

		cmt, err := db.PostComment(context.Background(), comment.Comment{
			Slug:   "slug",
			Author: "author",
			Body:   "body",
		})
		assert.NoError(t, err)

		postCmt, err := db.UpdateComment(context.Background(), cmt.ID, comment.Comment{
			Slug:   "slugput",
			Author: "authorput",
			Body:   "bodyput",
		})
		assert.NoError(t, err)
		assert.Equal(t, postCmt.Slug, "slugput", "comment field should be equal to slugput")
	})
}
