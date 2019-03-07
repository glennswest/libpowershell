package main

import (
	"github.com/masterzen/winrm"
        "log"
	"os"
)

func main(){
host := "winnode01.red.k.e2e.bos.redhat.com"
endpoint := winrm.NewEndpoint(host, 5985, false, false, nil, nil, nil, 0)
client, err := winrm.NewClient(endpoint, "Administrator", "Secret2018")
if err != nil {
        log.Printf("Panic\n")
	panic(err)
}
log.Printf("Prepare to ipconfig\n")
client.Run("ipconfig /all", os.Stdout, os.Stderr)
log.Printf("Finish\n");
}

