package forms

import (
	"log"
	"os/exec"
)

var shellScriptsFolder = ".\\shellScripts\\"

func callScriptBat(scriptPath string, arg ...string) string {
	cmd := exec.Command(scriptPath, arg...)
	out, err := cmd.CombinedOutput()
	sOut := string(out)
	//sOut = strings.Replace(sOut, "\r\n", "", -1)
	if err != nil {
		log.Println(sOut)
		log.Println(err)
		return ""
	}
	return sOut
}
