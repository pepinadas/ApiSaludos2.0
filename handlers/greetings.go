package handlers

import (
	"RestGolang/models"
	"RestGolang/repository"
	"RestGolang/server"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/segmentio/ksuid"
)

type UpsertGreetRequest struct {
	GreetingContent string `json:"greet_content"`
	Lenguage        string `json:"lenguage"`
}

type GreetResponse struct {
	Id           string `json:"id"`
	GreetContent string `json:"greet_content"`
	Lenguage     string `json:"lenguage"`
}

type PostUpdateResponse struct {
	Message string `json:"message"`
}

func UpdateGreetHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
		token, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(s.Config().JWTSecret), nil
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
			var greetRequest = UpsertGreetRequest{}
			if err := json.NewDecoder(r.Body).Decode(&greetRequest); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			greet := models.Greet{
				Id:              params["id"],
				GreetingContent: greetRequest.GreetingContent,
				Lenguage:        greetRequest.Lenguage,
				UserId:          claims.UserId,
			}
			err = repository.UpdateGreet(r.Context(), &greet)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(PostUpdateResponse{
				Message: "Greet updated",
			})
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func InsertGreetHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
		token, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(s.Config().JWTSecret), nil
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
			var greetRequest = UpsertGreetRequest{}
			if err := json.NewDecoder(r.Body).Decode(&greetRequest); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			id, err := ksuid.NewRandom()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			greet := models.Greet{
				Id:              id.String(),
				GreetingContent: greetRequest.GreetingContent,
				Lenguage:        greetRequest.Lenguage,
				UserId:          claims.UserId,
			}
			err = repository.InsertGreet(r.Context(), &greet)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(GreetResponse{
				Id:           greet.Id,
				GreetContent: greet.GreetingContent,
				Lenguage:     greet.Lenguage,
			})
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func DeleteGreetHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
		token, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(s.Config().JWTSecret), nil
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {

			err = repository.DeleteGreet(r.Context(), params["id"], claims.UserId)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(PostUpdateResponse{
				Message: "Greet deleted",
			})
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func GetGreetByIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		post, err := repository.GetGreetById(r.Context(), params["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(post)
	}
}

func ListGreetHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		pageStr := r.URL.Query().Get("page")
		var page = uint64(0)
		if pageStr != "" {
			page, err = strconv.ParseUint(pageStr, 10, 64)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}
		greets, err := repository.ListGreet(r.Context(), page)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(greets)
	}
}
