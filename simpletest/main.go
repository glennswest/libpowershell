package main

import (
	"github.com/masterzen/winrm"
        "log"
        "time"
	"os"
)

func main(){
host := "winnode01"
endpoint := winrm.NewEndpoint(host, 5985, false, false, nil, nil, nil, 0)
client, err := winrm.NewClient(endpoint, "Administrator", "Secret2018")
if err != nil {
        log.Printf("Panic\n")
	panic(err)
}
log.Printf("Prepare to ipconfig\n")
client.Run("ipconfig /all > x.out", os.Stdout, os.Stderr)
time.Sleep(2 * time.Second)
log.Printf("Finish\n");
}

