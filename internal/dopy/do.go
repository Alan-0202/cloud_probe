package dopy

import (
	"bytes"
	"cloudprobe/internal/log"
	"cloudprobe/utils"
	"fmt"
	"os/exec"
)

func HandlerPyNoArgs(fPath, scriptName string) (res string, err error) {
	str := utils.CleanStr(fPath + scriptName)
	args := []string{str}

	cmd := exec.Command("python", args...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()

	if err != nil {
		log.Error(fmt.Sprintf("exec is err: %v", stderr.String()))
		return "Failed", err
	}
	res = out.String()
	fmt.Println(fmt.Sprintf("result: %v", res))

	r := utils.CleanStr(res)
	return r, nil
}

// with args
func HandlerPyWithArgs(fPath, scriptName, arg string) (res string, err error) {
	// clean
	str := utils.CleanStr(fPath + scriptName)
	str1 := utils.CleanStr(arg)

	args := []string{str, str1}

	cmd := exec.Command("python", args...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		log.Error(fmt.Sprintf("exec is err: %v", stderr.String()))
		return "Failed", err
	}
	res = out.String()
	r := utils.CleanStr(res)
	return r, nil
}