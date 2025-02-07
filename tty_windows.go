// Copyright 2019 Graham Clark. All rights reserved.  Use of this source
// code is governed by the MIT license that can be found in the LICENSE
// file.

package termshark

//======================================================================

type TerminalSignals struct {
	set bool
}

func (t *TerminalSignals) IsSet() bool {
	return t.set
}

func (t *TerminalSignals) Restore() {
	t.set = false
}

func (t *TerminalSignals) Set() error {
	t.set = true
	return nil
}

//======================================================================
// Local Variables:
// mode: Go
// fill-column: 78
// End:
