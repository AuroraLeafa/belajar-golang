package main

import (
	"flag"
	"fmt"
)

func main() {
	var username *string = flag.String("username", "root", "database_username")
	var password *string = flag.String("password", "root", "database_password")
	var host *string = flag.String("host", "root", "database_host")
	var port *int = flag.Int("port", 0, "database_port")

	flag.Parse()
	fmt.Println("Username : ", *username)
	fmt.Println("Password : ", *password)
	fmt.Println("host : ", *host)
	fmt.Println("port : ", *port)

}
