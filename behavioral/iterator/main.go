package main

import "fmt"

// iterator提供一个简单遍历接口，他人无需了解你内部数据结构即可遍历你的数据。
func main() {
	user1 := &user{
		name: "a",
		age:  30,
	}
	user2 := &user{
		name: "b",
		age:  20,
	}

	userCollection := &userCollection{
		users: []*user{user1, user2},
	}

	iterator := userCollection.createIterator()

	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Printf("User is %+v\n", user)
	}
}

// iterator interface
type iterator interface {
	getNext() *user
	hasNext() bool
}

// collection interface
type collection interface {
	createIterator() iterator
}

type user struct {
	name string
	age  int
}

// userCollection : concrete interface
type userCollection struct {
	users []*user
}

func (u *userCollection) createIterator() iterator {
	return &userIterator{users: u.users}
}

type userIterator struct {
	index int
	users []*user
}

func (u *userIterator) hasNext() bool {
	if u.index < len(u.users) {
		return true
	}
	return false

}
func (u *userIterator) getNext() *user {
	if u.hasNext() {
		user := u.users[u.index]
		u.index++
		return user
	}
	return nil
}
