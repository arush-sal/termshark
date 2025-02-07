// Copyright 2019 Graham Clark. All rights reserved.  Use of this source
// code is governed by the MIT license that can be found in the LICENSE
// file.
//
// +build !windows

package termshark

import (
	"os"
	"os/signal"
	"syscall"
)

//======================================================================

func RegisterForSignals(ch chan<- os.Signal) {
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGTSTP, syscall.SIGCONT)
}

func IsSigTSTP(sig os.Signal) bool {
	if ssig, ok := sig.(syscall.Signal); ok && ssig == syscall.SIGTSTP {
		return true
	}
	return false
}

func IsSigCont(sig os.Signal) bool {
	if ssig, ok := sig.(syscall.Signal); ok && ssig == syscall.SIGCONT {
		return true
	}
	return false
}

func StopMyself() error {
	return syscall.Kill(syscall.Getpid(), syscall.SIGSTOP)
}

//======================================================================
// Local Variables:
// mode: Go
// fill-column: 78
// End:
