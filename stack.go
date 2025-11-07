package kutils

import (
	"bytes"
	"runtime"
)

// ShortStack returns a short stack trace, including only until the first few traces from the same package prefix as the
// caller.
func ShortStack(count ...int) []byte {
	cnt := 3
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
		var pkg []byte
		var start int
		if idxNewLine >= 0 {
			idxSlash := prevIdxNewLine
			var prevIdxSlash int
			for range 2 {
				prevIdxSlash = idxSlash
				if idxSlash = bytes.IndexByte(buf[prevIdxSlash+1:idxNewLine], '/'); idxSlash < 0 {
					break
				} else if idxSlash += prevIdxSlash + 1; idxSlash >= idxNewLine-1 {
					break
				}
			}
			if idxSlash >= 0 {
				prevIdxSlash = idxSlash
				if idxSlash = bytes.IndexAny(buf[prevIdxSlash+1:idxNewLine], "/."); idxSlash >= 0 {
					pkg = buf[prevIdxNewLine+1 : prevIdxSlash+idxSlash]
					start = bytes.IndexByte(buf[idxNewLine+1:], '\n')
				}
			}
			if start >= 0 && pkg != nil {
				start += idxNewLine + 2
				idxSamePkg := start
				var prevIdxSamePkg int
				for range cnt {
					prevIdxSamePkg = idxSamePkg
					if idxSamePkg = bytes.Index(buf[prevIdxSamePkg+1:], pkg); idxSamePkg < 0 {
						break
					}
					idxSamePkg += prevIdxSamePkg + 1
				}
				if idxSamePkg >= 0 {
					idxNewLine = bytes.IndexByte(buf[idxSamePkg+1:], '\n')
					if idxNewLine >= 0 {
						return buf[start : idxSamePkg+1+idxNewLine]
					}
				}
			} else {
				start = idxNewLine + 1
			}
		}

		if n < len(buf) {
			return buf[start:n]
		}
		buf = make([]byte, 2*len(buf))
	}
}
