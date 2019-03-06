package libpowershell;

import(
	    "fmt"
            "strings"
	)



func WmiGetWinVersion() string {
// Major  Minor  Build  Revision
// -----  -----  -----  --------
// 10     0      17134  0
	result := WmiPowershell("[System.Environment]::OSVersion.Version");
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

        r, w, _ := os.Pipe()
        outC := make(chan string)
        go func(){
             var buf bytes.Buffer
             io.Copy(&buf,r)
             outC <- buf.String()
             }()

        endpoint := winrm.NewEndpoint(host, 5986, false, false, nil, nil, nil, 0)
	client, err := winrm.NewClient(endpoint, "Administrator", "secret")
	if err != nil {
            log.Printf("WmiPowerShell: %s) %s - Failed - %s\n",host,thecmd,err)
            return(err)
	    }

	client.Run(thecmd, w, w)
        w.Close()
       result := <-outC
       log.Printf("%s\n",result)
       return result
}
