package main;

import(
	    "fmt"
            "strings"
            "github.com/glennswest/libpowershell/pshell"
	)


func main(){
  host := "10.19.114.61"
  pshell.SetRemoteMode(host,"Administrator","Secret2018")
  
  thever := pshell.GetWinVersion();
  fmt.Printf("Windows Version = %s\n",thever)
 
  pcmd := `$Env:master="openshift.green.k.e2e.bos.redhat.com";$Env:wmmurl="winmachineman.apps.green.k.e2e.bos.redhat.com";$Env:template="/templates/win10.0.17763.template";$Env:beta_kubernetes_io_os="windows";$Env:kubernetes_io_hostname="winnode01";$Env:beta_kubernetes_io_arch="amd64";$Env:host_ip="10.19.114.61";$Env:ovn_host_subnet="10.128.64.0/24";C:/bin/prereq1809.ps1`
  //pcmd := `$Env:master=\"openshift.green.k.e2e.bos.redhat.com\";$Env:master`
  //pc := strings.Replace(pcmd,`\`,`\\`,-1)
  pc := strings.Replace(pcmd,`"`,`\"`,-1)
  result := pshell.Powershell(pc)
  fmt.Printf("Result = %s\n",result)
  fmt.Printf("Length = %d\n",len(result))
}


