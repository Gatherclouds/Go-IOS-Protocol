package log

import (
	"net/http"
	"io/ioutil"
	"errors"
	"strconv"
	"time"
	"net/url"
)

var Server = "127.0.0.1:1001"
var LocalID = "default"

func Report(msg Msg) error {
	resp, err := http.PostForm(Server+"/report",
		msg.Form())

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(strconv.Itoa(resp.StatusCode) + " " + string(body))
	}
	return nil
}

type Timestamp struct {
	second int64
	nano   int
}

func (t Timestamp) String() string {
	return strconv.FormatInt(t.second, 10) + "+" + strconv.Itoa(t.nano)
}

func Now() Timestamp {
	now := time.Now()
	return Timestamp{
		second: now.Unix(),
		nano:   now.Nanosecond(),
	}
}

type Msg interface {
	Form() url.Values
}
