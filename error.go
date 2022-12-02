package download

import (
	"fmt"
	"io"
	"io/ioutil"
)

var (
	_ error = (*InvalidResponseCode)(nil)
	_ error = (*DeadlineExceeded)(nil)
	_ error = (*Canceled)(nil)
)

func NewInvalidResponseCode(got, expected int, body io.ReadCloser) *InvalidResponseCode {
	resp := &InvalidResponseCode{
		expected: expected,
		got:      got,
	}
	if body != nil {
		bs, err := ioutil.ReadAll(body)
		if err != nil {
			resp.body = fmt.Sprintf("read body failed: %s", err)
			return resp
		}

		resp.body = string(bs)
	}
	return resp
}

// InvalidResponseCode is the error containing the invalid response code error information
type InvalidResponseCode struct {
	expected int
	got      int
	body     string
}

// Error returns the InvalidResponseCode error string
func (e *InvalidResponseCode) Error() string {
	if e.body == "" {
		return fmt.Sprintf("Invalid response code, received '%d' expected '%d'", e.got, e.expected)
	}
	return fmt.Sprintf("Invalid response code, received '%d' expected '%d', body:[%s]", e.got, e.expected, e.body)
}

// DeadlineExceeded is the error containing the deadline exceeded error information
type DeadlineExceeded struct {
	url string
}

// Error returns the DeadlineExceeded error string
func (e *DeadlineExceeded) Error() string {
	return fmt.Sprintf("Download timeout exceeded for '%s'", e.url)
}

// Canceled is the error containing the cancelled error information
type Canceled struct {
	url string
}

// Error returns the Canceled error string
func (e *Canceled) Error() string {
	return fmt.Sprintf("Download canceled for '%s'", e.url)
}
