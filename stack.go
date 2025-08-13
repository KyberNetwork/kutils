package kutils

import (
	"bytes"
	"runtime"
)

// ShortStack returns a short stack trace, including only until the first few traces from the same package organization
// prefix as the caller.
func ShortStack(count ...int) []byte {
	cnt := 2
	if len(count) > 0 {
		cnt = count[0]
	}
	buf := make([]byte, 2048)
	for {
		n := runtime.Stack(buf, false)

		var prevIdxNewLine, idxNewLine int
		for range 4 {
			prevIdxNewLine = idxNewLine
			if idxNewLine = bytes.IndexByte(buf[prevIdxNewLine+1:], '\n'); idxNewLine < 0 {
				break
			}
			idxNewLine += prevIdxNewLine + 1
		}
		var org []byte
		var start int
		if idxNewLine >= 0 {
			caller := buf[prevIdxNewLine+1 : idxNewLine]
			idxSlash := bytes.IndexByte(caller, '/')
			if idxSlash >= 0 {
				idxSecondSlash := bytes.IndexByte(caller[idxSlash+1:], '/')
				if idxSecondSlash >= 0 {
					org = caller[:idxSlash+idxSecondSlash+1]
				}
			}
			start = bytes.IndexByte(buf[idxNewLine+1:], '\n')
			if start >= 0 {
				start += idxNewLine + 2
				idxSameOrg := start
				var prevIdxSameOrg int
				for range cnt {
					prevIdxSameOrg = idxSameOrg
					if idxSameOrg = bytes.Index(buf[prevIdxSameOrg+1:], org); idxSameOrg < 0 {
						break
					}
					idxSameOrg += prevIdxSameOrg + 1
				}
				if idxSameOrg >= 0 {
					idxNewLine = bytes.IndexByte(buf[idxSameOrg+1:], '\n')
					if idxNewLine >= 0 {
						return buf[start : idxSameOrg+1+idxNewLine]
					}
				}
			}
		}

		if n < len(buf) {
			return buf[start:n]
		}
		buf = make([]byte, 2*len(buf))
	}
}
