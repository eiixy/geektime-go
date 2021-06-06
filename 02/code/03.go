package code

import (
	"bufio"
	"fmt"
	"io"

	"github.com/pkg/errors"
)

var ErrTest = errors.New("test error")

func Test03() {
	err := withMessage()
	if errors.Is(err, ErrTest) {
		// errors.Cause(err) 获取根因
		fmt.Printf("original error: %T %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trace:\n%+v\n", err)
	}
}

func CountLinesA(r io.Reader) (int, error) {
	var (
		br    = bufio.NewReader(r)
		lines int
		err   error
	)
	for {
		_, err = br.ReadString('\n')
		lines++
		if err != nil {
			break
		}
	}
	if err != nil {
		return 0, err
	}

	return lines, nil
}

// 先暂存err，然后再返回
func CountLinesB(r io.Reader) (int, error) {
	sc := bufio.NewScanner(r)
	lines := 0
	for sc.Scan() {
		lines++
	}
	return lines, sc.Err()
}

type Header struct {
	Key, Value string
}

type Status struct {
	Code   int
	Reason string
}

func WriteResponseA(w io.Writer, st Status, headers []Header, body io.Reader) error {
	_, err := fmt.Fprintf(w, "HTTP/1.1 %d %s\r\n", st.Code, st.Reason)
	if err != nil {
		return err
	}

	for _, h := range headers {
		_, err := fmt.Fprintf(w, "%s: %s\r\n", h.Key, h.Value)
		if err != nil {
			return err
		}
	}

	if _, err := fmt.Fprint(w, "\r\n"); err != nil {
		return err
	}

	_, err = io.Copy(w, body)

	return err
}

type errWriter struct {
	io.Writer
	err error
}

func (e *errWriter) Write(buf []byte) (int, error) {
	if e.err != nil {
		return 0, e.err
	}

	var n int
	n, e.err = e.Writer.Write(buf)
	return n, nil
}

func WriteResponseB(w io.Writer, st Status, headers []Header, body io.Reader) error {
	ew := &errWriter{Writer: w}
	fmt.Fprintf(ew, "HTTP/1.1 %d %s\r\n", st.Code, st.Reason)

	for _, h := range headers {
		fmt.Fprintf(ew, "%s: %s\r\n", h.Key, h.Value)
	}

	fmt.Fprint(ew, "\r\n")
	io.Copy(ew, body)

	return ew.err
}

func wrapError() error {
	// 包装 error
	return errors.Wrap(ErrTest, "wrap message")
}

func withMessage() error {
	err := wrapError()
	return errors.WithMessage(err, "with message")
}
