package libpowershell;

import(
	    "fmt"
            "strings"
            "github.com/libpowershell/libremotepowershell"
	)


func main(){
  host := "winnode01"
  username := "Adminisrator"
  password := "Secret2018"
  thever := WmiGetWinVersion(host,username,password);
  fmt.Printf("Windows Version = %s\n",thever);
}


