package main

import (
	"bufio"
	"fmt"
	"os"

	"currydeveloper.com/ldap-client/ldapservice"
)

func main() {
	user := "uid=admin,ou=system"
	host := "localhost"
	port := "10389"
	fmt.Println("Entering ldap client")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the password")
	pwd, _ := reader.ReadString('\n')
	fmt.Printf("Entered %v", pwd)
	ldapservice.NewConnection(user, pwd, host, port)
}
