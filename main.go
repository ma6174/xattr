package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pkg/xattr"
)

func main() {
	var err error
	var b []byte
	op, path := os.Args[1], os.Args[2]
	var name string
	var data string
	if len(os.Args) > 3 {
		name = os.Args[3]
	}
	if len(os.Args) > 4 {
		data = os.Args[4]
	}
	var names []string
	switch op {
	case "get", "g":
		b, err = xattr.Get(path, name)
	case "lget", "lg":
		b, err = xattr.LGet(path, name)
	case "set", "s":
		err = xattr.Set(path, name, []byte(data))
	case "lset", "ls":
		err = xattr.LSet(path, name, []byte(data))
	case "remove", "r":
		err = xattr.Remove(path, name)
	case "lremove", "lr":
		err = xattr.LRemove(path, name)
	case "list", "l":
		names, err = xattr.List(path)
	case "llist", "ll":
		names, err = xattr.LList(path)
	}
	if err != nil {
		log.Fatalln(err, os.Args[1:])
	}
	fmt.Println(os.Args[1:], "==>", string(b), strings.Join(names, ","))
}
