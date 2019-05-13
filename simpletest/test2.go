package main

import (
	"github.com/glennswest/libpowershell/pshell"
        "log"
)

func main(){
pcmd := `certutil -addstore -f "TrustedPublisher" c:\k\data\cloudbase.sst;msiexec /i \bin\openvswitch-hyperv.msi ADDLOCAL="OpenvSwitchCLI,OpenvSwitchDriver,OVNHost" /qn /l*v log.txt`
result := pshell.Powershell(pcmd)
log.Printf("Finish %s\n",result);
}

