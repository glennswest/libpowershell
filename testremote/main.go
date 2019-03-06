package main;

import(
	    "fmt"
            "github.com/glennswest/libpowershell/libwmi"
	)


func main(){
  host := "winnode01"
  username := "Adminisrator"
  password := "Secret2018"
  thever := libwmi.WmiGetWinVersion(host,username,password)
  fmt.Printf("Windows Version = %s\n",thever)
}


