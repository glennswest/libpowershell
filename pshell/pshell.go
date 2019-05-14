package pshell

import(
	    "fmt"
	    "os/exec"
            "strings"
            "github.com/masterzen/winrm"
            "io/ioutil"
            "log"
            "os"
	)


var Mode string
var User string
var Password string
var Host string

func init(){
     SetLocalMode()
}

func SetLocalMode(){
     Mode = "local"
     User = ""
     Password = ""
     Host = ""
}

func SetRemoteMode(thehostname string,theusername string, thepassword string){
     Mode = "remote"
     Host = thehostname
     User = theusername
     Password = thepassword
}

func PsReset(){
	fmt.Printf("Reseting\n");
}

func GetLineArray(theresult string ,thelinenum int ) [] string {
	thelines := lines(theresult);
	output := standardizeSpaces(thelines[thelinenum]);
	varvals := strings.Split(output," ");
	return(varvals);
}

func GetPsHostName() string {
	cmd := "$env:COMPUTERNAME";
	result := Powershell(cmd);
	fmt.Printf("%s\n%s\n",cmd,result);
	return(result);
}

func GetPsInstalled(thepackage string) string {
	cmd := "Get-WindowsFeature -Name '" + thepackage + "'";
	result := Powershell(cmd);
	varvals := GetLineArray(result,3);
	return(varvals[4]);
}

func GetPsRegValue(thenamespace string,thevaluename string) string {
	cmd := "Get-ItemProperty -Path " + thenamespace + " -Name '" + thevaluename + "' -ErrorAction SilentlyContinue";
	result := Powershell(cmd);
	if (len(result) == 0){
		return "";
	  }
	varvals := GetLineArray(result,2);
	return(varvals[2]);
}

func SetPSRegValue(thenamespace string, thevaluename string, thevalue string, thetype string){
	thecmd := "New-ItemProperty -Path " + thenamespace + " -Name '" + thevaluename + "' -Value '" + thevalue + "' -PropertyType '" + thetype + "'";
  result := Powershell(thecmd);
	fmt.Printf("SetPSRegValue: %s\n", result);
}
func GetWinVersion() string {
// Major  Minor  Build  Revision
// -----  -----  -----  --------
// 10     0      17134  0
	result := Powershell("[System.Environment]::OSVersion.Version");
        thelines := lines(result);
        if (len(thelines) < 3){
           return ""
           }

	verstr := standardizeSpaces(thelines[3]);
	va := strings.Split(verstr," ");
	ver := va[0] + "." + va[1] + "." + va[2]
	return(ver);
}


func lines(theval string) [] string {
	 values := strings.Split(strings.Replace(theval, "\r\n", "\n", -1), "\n");
	 return(values);
}

func standardizeSpaces(s string) string {
    return strings.Join(strings.Fields(s), " ")
}

func Powershell(thecmd string) string {
      switch Mode {
          case "local":
               log.Printf("PS> %s\n", thecmd)
               lcmd := "powershell " + thecmd 
               result := LocalPowershell(lcmd)
               log.Printf("%s\n",result)
               return(result)
          case "remote":
               rcmd := "powershell " + thecmd 
               log.Printf("PS> %s\n", thecmd)
               result := WmiPowershell(Host,User,Password,rcmd)
               log.Printf("%s\n",result)
               return(result)
          }
      return("Error")
}


func LocalPowershell(thecmd string) string {
	theargs := strings.Split(thecmd," ");
	c,err := exec.Command("powershell", theargs...).CombinedOutput();
	cmd := string(c);

	if  err != nil {
	    return("");
    } else {
	    return(cmd);
    }
}

func WmiPowershell(host,user,password,thecmd string) string {
        r, w, _ := os.Pipe()

        endpoint := winrm.NewEndpoint(host, 5985, false, false, nil, nil, nil, 0)
        client, err := winrm.NewClient(endpoint, user, password)

        if err != nil {
            return(err.Error())
            }

        client.Run(thecmd, w, w)
        w.Close()
        out, _ := ioutil.ReadAll(r)
        result := string(out)
       return result
}
