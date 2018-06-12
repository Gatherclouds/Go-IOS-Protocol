package log


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
