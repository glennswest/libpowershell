package main;

import(
	    "fmt"
            "github.com/glennswest/libpowershell/libpowershell"
	)


func main(){
  host := "winnode01"
  SetRemoteMode(host,"Administrator","Secret2018")
  
  fmt.Printf("Windows Version = %s\n",thever)
}


