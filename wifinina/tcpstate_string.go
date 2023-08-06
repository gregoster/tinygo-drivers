//go:build wifidebug

// Code generated by "stringer -type=TCPState -trimprefix=TCPState"; DO NOT EDIT.

package wifinina

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TCPStateClosed-0]
	_ = x[TCPStateListen-1]
	_ = x[TCPStateSynSent-2]
	_ = x[TCPStateSynRcvd-3]
	_ = x[TCPStateEstablished-4]
	_ = x[TCPStateFinWait1-5]
	_ = x[TCPStateFinWait2-6]
	_ = x[TCPStateCloseWait-7]
	_ = x[TCPStateClosing-8]
	_ = x[TCPStateLastACK-9]
	_ = x[TCPStateTimeWait-10]
}

const _TCPState_name = "ClosedListenSynSentSynRcvdEstablishedFinWait1FinWait2CloseWaitClosingLastACKTimeWait"

var _TCPState_index = [...]uint8{0, 6, 12, 19, 26, 37, 45, 53, 62, 69, 76, 84}

func (i TCPState) String() string {
	if i >= TCPState(len(_TCPState_index)-1) {
		return "TCPState(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TCPState_name[_TCPState_index[i]:_TCPState_index[i+1]]
}
