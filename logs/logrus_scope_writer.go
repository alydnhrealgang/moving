package logs

import (
	"io"
)

func NewWriter(scope *LogrusScope) io.Writer {
	return &logrusScopeWriter{
		scope:      scope,
		lineBuffer: make([]byte, 0, 1000),
	}
}

type logrusScopeWriter struct {
	scope      *LogrusScope
	lineBuffer []byte
}

func (l *logrusScopeWriter) Write(p []byte) (n int, err error) {
	for _, c := range p {
		if c == '\n' {
			l.scope.Info(string(l.lineBuffer))
			l.lineBuffer = make([]byte, 0, 1000)
		} else {
			l.lineBuffer = append(l.lineBuffer, c)
		}
	}

	return len(p), nil
}
