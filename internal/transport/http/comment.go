package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/GleeN987/go-rest-api/internal/comment"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type PostCommentRequest struct {
	Slug   string `json:"slug" validate:"required"`
	Body   string `json:"body" validate:"required"`
	Author string `json:"author" validate:"required"`
}

func PostCommentRequestToComment(r PostCommentRequest) comment.Comment {
	return comment.Comment{
		Slug:   r.Slug,
		Body:   r.Body,
		Author: r.Author,
	}
}

type CommentService interface {
	GetComment(ctx context.Context, id string) (comment.Comment, error)
	PostComment(ctx context.Context, cmt comment.Comment) (comment.Comment, error)
	DeleteComment(ctx context.Context, id string) error
	UpdateComment(ctx context.Context, id string, cmt comment.Comment) (comment.Comment, error)
}

func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	var cmtRequest PostCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&cmtRequest); err != nil {
		log.Printf("Error while decoding request body: %v", err)
		return
	}

	validate := validator.New()
	err := validate.Struct(cmtRequest)
	if err != nil {
		http.Error(w, "invalid post request", http.StatusBadRequest)
	}

	cmt := PostCommentRequestToComment(cmtRequest)
	cmtPosted, err := h.Service.PostComment(r.Context(), cmt)
	if err != nil {
		log.Printf("Error while posting comment: %v", err)
		return
	}
	if err := json.NewEncoder(w).Encode(cmtPosted); err != nil {
		log.Printf("Error while encoding to response: %v", err)
		return
	}
}

func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cmt, err := h.Service.GetComment(r.Context(), id)
	if err != nil {
		log.Printf("Error while getting the comment: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}

}
func (h *Handler) PutComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var cmt comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
		log.Printf("Error while decoding request body: %v", err)
		return
	}
	cmt, err := h.Service.UpdateComment(r.Context(), id, cmt)
	if err != nil {
		log.Printf("Error while updating the comment: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}
}
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.Service.DeleteComment(r.Context(), id); err != nil {
		log.Printf("Error while deleting the comment: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode("Comment deleted"); err != nil {
		log.Printf("Error while encoding: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
