package main;

import(
	    "fmt"
            "strings"
            "github.com/glennswest/libpowershell/pshell"
	)


func main(){
  host := "winnode01"
  pshell.SetRemoteMode(host,"Administrator","Secret2018")
  
  thever := pshell.GetWinVersion();
  fmt.Printf("Windows Version = %s\n",thever)
 
  result := pshell.Powershell("Test-Path -Path \"/Program` Files/WindowsNodeManager/winnodeman.exe\"")
  fmt.Printf("Result = %s\n",result)
  fmt.Printf("Length = %d\n",len(result))
  if (strings.Compare(result[0:4],"True") == 0){
     fmt.Printf("It Really True\n")
     }
}


