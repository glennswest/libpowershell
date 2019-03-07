package libwmi;

import(
            "os"
            "strings"
            "github.com/masterzen/winrm"
            "io/ioutil"
            "log"
	)



func WmiGetWinVersion(host string,user string,password string) string {
// Major  Minor  Build  Revision
// -----  -----  -----  --------
// 10     0      17134  0
	result := WmiPowershell(host,user,password,"powershell [System.Environment]::OSVersion.Version");
        thelines := lines(result);

	verstr := standardizeSpaces(thelines[3]);
	va := strings.Split(verstr," ");
	ver := va[0] + "." + va[1] + "." + va[2] + "." + va[3];
	return(ver);
}


func lines(theval string) [] string {
	 values := strings.Split(strings.Replace(theval, "\r\n", "\n", -1), "\n");
	 return(values);
}

func standardizeSpaces(s string) string {
    return strings.Join(strings.Fields(s), " ")
}

func WmiPowershell(host,user,password,thecmd string) string {

        log.Printf("WmiPowershell: %s - %s - %s - %s\n",host,user,password,thecmd)
        r, w, _ := os.Pipe()

        log.Printf("Starting new endpoint\n")
        endpoint := winrm.NewEndpoint(host, 5985, false, false, nil, nil, nil, 0)
        log.Printf("WmiPowershell: Finished NewEndpoint\n")
	client, err := winrm.NewClient(endpoint, user, password)
        log.Printf("WmiPowershell: Finished Newclient\n")
       
	if err != nil {
            log.Printf("WmiPowerShell: %s) %s - Failed - %s\n",host,thecmd,err.Error())
            return(err.Error())
	    }

        log.Printf("WmiPowerShell: Starting client.Run\n")
	client.Run(thecmd, w, w)
        w.Close()
        out, _ := ioutil.ReadAll(r)
        result := string(out)
       log.Printf("%s\n",result)
       return result
}
