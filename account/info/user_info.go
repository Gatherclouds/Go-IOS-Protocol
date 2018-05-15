package info

type UserInfo struct {
	txt []byte
}

func (info *UserInfo) ToString() string { return string(info.txt) }

