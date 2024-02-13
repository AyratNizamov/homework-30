package usecase

import (
	"30New/entity"
	"strconv"
)

type Methods interface {
	Create() string
	Delete() string
	GetFriends() string
	AddFriend() string
	UpdateAge() string
}

type UseCase struct {
	meth Methods
}

var r = make(map[int]*entity.User, 0)

func Create(name string, age int) string {
	newUser := entity.NewUser()
	newUser.Name = name
	newUser.Age = age
	r[len(r)] = newUser
	return "Пользователь id " + strconv.Itoa(len(r)-1) + " создан"
}

func Delete(id int) string {
	if id >= len(r) {
		return "Пользователя id " + strconv.Itoa(id) + " не существует"
	}
	name := r[id].Name
	delete(r, id)
	return "Пользователь " + name + " удален"
}

func GetFriends(idUser int) string {
	result := "Друзья пользователя " + strconv.Itoa(idUser) + ":\n"
	for _, id := range r[idUser].Friend {
		result += "Name is " + r[id].Name + " and age " + strconv.Itoa(r[id].Age) + "\n"
	}
	return result
}

func AddFriend(first, second int) (string, bool) {
	if first >= len(r) {
		return "Пользователя id " + strconv.Itoa(first) + " не существует", false
	}

	if second >= len(r) {
		return "Пользователя id " + strconv.Itoa(second) + " не существует", false
	}

	r[first].Friend = append(r[first].Friend, second)
	r[second].Friend = append(r[second].Friend, first)

	return r[first].Name + " и " + r[second].Name + " теперь друзья", true
}

func UpdateAge(id, age int) (string, bool) {
	if id >= len(r) {
		return "Пользователя id " + strconv.Itoa(id) + " не существует", false
	}

	r[id].Age = age
	return "Возраст пользователя id: " + strconv.Itoa(id) + " успешно обновлен", true
}
