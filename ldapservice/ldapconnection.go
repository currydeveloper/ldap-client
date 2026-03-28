package ldapservice

import (
	"fmt"

	ldap "github.com/go-ldap/ldap/v3"
)

func NewConnection(user string, password string, host string, port string) {
	fmt.Println("Called from the new func")
	l, err := ldap.DialURL("ldap://" + host + ":" + port)
	if err != nil {
		fmt.Println(fmt.Errorf("dial failed: %w", err))
		return
	}
	defer l.Close()
	fmt.Println("Entered the connection")
	bind := ldap.SimpleBindRequest{Username: user, Password: password}
	_, err = l.SimpleBind(&bind)
	fmt.Println("Bind completed")

}
