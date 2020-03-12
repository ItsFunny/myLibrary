/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-07-04 13:35 
# @File : command.go
# @Description : command 插件
*/
package plugin

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"log"
	"os"
	"os/exec"
	"syscall"
)

// CapturingPassThroughWriter is a writer that remembers
// data written to it and passes it to w
type CapturingPassThroughWriter struct {
	buf bytes.Buffer
	w   io.Writer
}

func APP() {

}

// NewCapturingPassThroughWriter creates new CapturingPassThroughWriter
func NewCapturingPassThroughWriter(w io.Writer) *CapturingPassThroughWriter {
	return &CapturingPassThroughWriter{
		w: w,
	}
}
func (w *CapturingPassThroughWriter) Write(d []byte) (int, error) {
	w.buf.Write(d)
	return w.w.Write(d)
}

// Bytes returns bytes written to the writer
func (w *CapturingPassThroughWriter) Bytes() []byte {
	return w.buf.Bytes()
}

func ExecCmdWithLog(cmds string) error {
	var errStdout, errStderr error
	cmd := exec.Command("/bin/sh", "-c", cmds)
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	stdout := NewCapturingPassThroughWriter(os.Stdout)
	stderr := NewCapturingPassThroughWriter(os.Stderr)
	err := cmd.Start()
	if err != nil {
		return errors.Errorf("cmd.Start() failed with '%s'\n", err)
	}
	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
	}()
	go func() {
		_, errStderr = io.Copy(stderr, stderrIn)
	}()
	err = cmd.Wait()
	if err != nil {
		return errors.Errorf("cmd.Run() failed with %s\n", err)
	}
	if errStdout != nil || errStderr != nil {
		return errors.Errorf("failed to capture stdout or stderr\n")
	}
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)

	return nil
}

// ExecCmdWithLog 的升级版,提供杀死进程的feature
func ExecCmdWithLogUnderControl(cmds string, killSigChan chan struct{}) error {
	var errStdout, errStderr error
	cmd := exec.Command("/bin/sh", "-c", cmds)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	go func() {
		for {
			if _, ok := <-killSigChan; ok {
				err := syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
				if nil != err {
					log.Println("杀死所有进程出现错误:", err.Error())
				}
				break
			} else {
				break
			}
		}
	}()
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	stdout := NewCapturingPassThroughWriter(os.Stdout)
	stderr := NewCapturingPassThroughWriter(os.Stderr)
	err := cmd.Start()
	if err != nil {
		return errors.New(fmt.Sprintf("cmd.Start() failed with '%s'\n", err))
	}
	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
	}()
	go func() {
		_, errStderr = io.Copy(stderr, stderrIn)
	}()
	err = cmd.Wait()
	if err != nil {
		return errors.New(fmt.Sprintf("cmd.Run() failed with %s\n", err))
	}
	if errStdout != nil || errStderr != nil {
		log.Println("failed to capture stdout or stderr")
	}
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)

	return nil
}
