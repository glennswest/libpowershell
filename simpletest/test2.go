package main

import (
	"github.com/glennswest/libpowershell/pshell"
        "log"
)

func main(){
pcmd := `$Env:master="openshift.green.k.e2e.bos.redhat.com";$Env:wmmurl="winmachineman.apps.green.k.e2e.b
os.redhat.com";$Env:template="/templates/win10.0.17763.template";$Env:kubernetes_io_hostname="winnode01";$Env:beta_kuber
netes_io_arch="amd64";$Env:beta_kubernetes_io_os="windows";$Env:host_ip="10.19.114.61";$Env:ovn_host_subnet="10.128.55.0
/24";\bin\prereq1809.ps1`
result := pshell.Powershell(pcmd)
log.Printf("Finish %s\n",result);
}

