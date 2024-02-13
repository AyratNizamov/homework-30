package controller

import (
	"30New/usecase"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Service struct {
	Store usecase.Methods
}

func (s *Service) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		content, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		var id int
		if err := json.Unmarshal(content, &id); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		result := usecase.Delete(id)

		jsonResult, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResult)
		return
	}

	w.WriteHeader(http.StatusBadRequest)
}

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (s *Service) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		content, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()
		var u user

		if err := json.Unmarshal(content, &u); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("1")
			w.Write([]byte(err.Error()))
			return
		}

		result := usecase.Create(u.Name, u.Age)

		jsonResult, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResult)

		return
	}

	w.WriteHeader(http.StatusBadRequest)
}

func (s *Service) UserFriends(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		content, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()
		var idUser int
		if err := json.Unmarshal(content, &idUser); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		result := usecase.GetFriends(idUser)

		jsonResult, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResult)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

type FriendsID struct {
	Source_id int `json:"source_id"`
	Target_id int `json:"target_id"`
}

func (s *Service) MakeFriends(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		content, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		var friends FriendsID

		if err := json.Unmarshal(content, &friends); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		result, ok := usecase.AddFriend(friends.Source_id, friends.Target_id)

		jsonResult, _ := json.Marshal(result)

		if ok == false {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonResult)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(jsonResult)
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
}

type ElementUser struct {
	IDUser int `json:"iduser"`
	NewAge int `json:"newage"`
}

func (s *Service) UpdateAge(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		content, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()
		var elements ElementUser

		if err := json.Unmarshal(content, &elements); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		result, ok := usecase.UpdateAge(elements.IDUser, elements.NewAge)

		jsonResult, _ := json.Marshal(result)

		if ok == false {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonResult)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(jsonResult)
			return
		}

	}
	w.WriteHeader(http.StatusBadRequest)
}
