package main

import (
	"t/libpowershell/libpowershell"
        "log"
)

func main(){
host := "winnode01.red.k.e2e.bos.redhat.com"
SetRemoteMode(host,"Administrator","Secret2018")
log.Printf("Prepare to ipconfig\n")
result := Powershell("ipconfig /all")
log.Printf("Finish %s\n",result);
}

