package user

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/MrTomSawyer/chat/internal/app/model"
	"github.com/gorilla/mux"
)

type UserServiceManager interface {
	Create(ctx context.Context, user *model.User) error
}

type Controller struct {
	r      *mux.Router
	s      UserServiceManager
	prefix string
}

func NewUserController(r *mux.Router, s UserServiceManager) *Controller {
	return &Controller{
		r:      r,
		s:      s,
		prefix: "user",
	}
}

func (c *Controller) Init() {
	c.r.HandleFunc(c.prefix+"/sign-up", c.Create).Methods("GET")
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := r.Body.Close()
		if err != nil {
			fmt.Println("failed to close request body")
		}
	}()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read request body: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var user model.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, "failed to unmarshal request body: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = c.s.Create(r.Context(), &user)
	if err != nil {
		http.Error(w, "failed to create user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
