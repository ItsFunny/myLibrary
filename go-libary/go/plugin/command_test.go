/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-07-04 13:42 
# @File : command_test.go
# @Description : 
*/
package plugin

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"testing"
)

func TestAPP(t *testing.T) {
	var errStdout, errStderr error
	cmd := exec.Command("ls", "-lah")
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	stdout := NewCapturingPassThroughWriter(os.Stdout)
	stderr := NewCapturingPassThroughWriter(os.Stderr)
	err := cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}
	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
	}()
	go func() {
		_, errStderr = io.Copy(stderr, stderrIn)
	}()
	err = cmd.Wait()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	if errStdout != nil || errStderr != nil {
		log.Fatalf("failed to capture stdout or stderr\n")
	}
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)
}

// 命令执行过程中获取输出
func TestAPP2(t *testing.T) {

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd := exec.Command("ls", "-lah")
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	var errStdout, errStderr error
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)
	err := cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}
	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
	}()
	go func() {
		_, errStderr = io.Copy(stderr, stderrIn)
	}()
	err = cmd.Wait()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	if errStdout != nil || errStderr != nil {
		log.Fatal("failed to capture stdout or stderr\n")
	}
	outStr, errStr := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())
	fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)
}

func TestExecCmdWithLog(t *testing.T) {
	// withLog := ExecCmdWithLog("you-get -o ~/Desktop/ https://stallman.org/rms.jpg ")
	// if nil != withLog {
	// 	fmt.Println(withLog)
	// }
	getenv := os.Getenv("PWD")
	fmt.Println(getenv)
}
