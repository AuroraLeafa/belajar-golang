package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (a UserSlice) Len() int           { return len(a) }
func (a UserSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a UserSlice) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	users := []User{
		{"Eko", 30},
		{"Eko1", 29},
		{"Eko2", 13},
		{"Eko3", 11},
		{"Eko4", 40},
		{"Eko5", 60},
	}
	sort.Sort(UserSlice(users))
	fmt.Println(users)
}
