package libpowershell;

import(
	    "fmt"
	    "os/exec"
            "strings"
	)


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

func Powershell(thecmd string) string {
	theargs := strings.Split(thecmd," ");
	c,err := exec.Command("powershell", theargs...).CombinedOutput();
	cmd := string(c);

	if  err != nil {
	    return("");
    } else {
	    return(cmd);
    }
}
