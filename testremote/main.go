package main;

import(
	    "fmt"
            "github.com/glennswest/libpowershell/pshell"
	)


func main(){
  host := "winnode01"
  pshell.SetRemoteMode(host,"Administrator","Secret2018")
  
  thever := pshell.GetWinVersion();
  fmt.Printf("Windows Version = %s\n",thever)
}


