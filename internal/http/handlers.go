package http

import (
	"API/internal/errors"
	"API/internal/model"
	"encoding/json"
	"fmt"
	"net/http"
)

type HTTPHandlers struct {
}

func (h *HTTPHandlers) CreatePost(w http.ResponseWriter, r *http.Request) {
	post := model.PostModel{}
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		errModel := model.NewErrorModel(err)
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(errModel); err != nil {
			fmt.Println(errors.ErrFromEncodeJson)
		}
	}

}
