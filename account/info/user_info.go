package info

type UserInfo struct {
	txt []byte
}

func (info *UserInfo) ToString() string { return string(info.txt) }

func GenerateUserInfo(txt []byte) *UserInfo {
	return &UserInfo{
		txt: txt,
	}
}

func (info *UserInfo) Update(new_txt []byte) {
	info.txt = new_txt
}
