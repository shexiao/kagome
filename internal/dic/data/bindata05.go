package data

import(
	"os"
	"time"
)

var _dicUniChardefDic = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\xdc\xed\x52\xda\x4c\x14\x00\xe0\x6c\xe4\x23\xaf\x32\xf3\x5e\xd3\x96\x52\xa5\x20\x75\x50\x7f\x38\xbd\x9a\xf6\x4f\x7b\xd3\x1d\x1a\x28\xda\x16\xa5\x82\x49\x5c\x22\xcf\xc3\x47\xc8\x26\x24\x67\x73\x92\x03\x33\x3b\x93\xc1\xe2\x4b\x1e\xf2\xc5\xd7\x2c\x0c\xb2\xec\x73\x39\x3d\xeb\xbf\x1f\x7d\x88\xb7\xd3\x9b\xee\xf5\x55\x1c\x8e\xba\x93\x38\xfb\x38\xee\x5d\xdf\x5d\xbe\xfb\x34\xed\xcf\x6e\x2f\x47\xf3\xf1\xb0\x1b\xa7\x57\x17\xb1\xb8\x18\xcf\xe3\x79\x9c\xc5\x62\x12\x6f\x62\xb9\x5e\x1c\xac\x56\xbe\x5f\xe9\x7c\x3e\x1a\x4d\x8a\xe1\xdd\x7c\x3c\x9d\x8e\x87\x3f\x42\xd6\x3b\xcd\xca\xf7\xdf\x42\x08\xd9\x9f\xf3\xf7\xcd\x27\x7f\xeb\x3c\x58\x37\x74\xb7\xda\x65\xf9\xe3\x1d\x6e\x3a\xf9\xb7\xed\x9b\xa7\x41\xcf\xe7\x8d\xad\xfe\x3b\x10\xcb\x58\x4e\x8f\xd6\xaf\x5c\x3c\xdd\x0a\x00\x00\x00\x00\x00\xd0\x5e\xa9\x07\x12\xd3\x49\x7d\xe4\x0f\xc3\x33\xc3\xeb\xb5\xeb\xec\xe9\xb5\xe3\xab\x6a\xdf\xfe\xb5\xbd\xbf\xb4\xdb\x5b\xab\x2f\xfb\x3a\xf8\x00\x69\x92\xf4\x93\x54\xd3\xf5\xb7\xed\xf5\x99\x66\xa5\x3e\xff\x38\x6e\x79\x02\xa9\xf7\xdf\xbc\xdd\x8f\x7f\xe5\xfa\xd1\x6b\x56\xd6\x6f\x56\xdd\xe7\xf3\xa1\xdb\xec\x7f\x1d\xbf\x21\x2d\x76\xf4\xff\x7f\x52\x57\x2a\x00\x00\x00\x00\x00\x00\x96\xea\x1b\x0f\x2c\xf2\xbc\x7c\x96\xaf\xa2\x81\x30\xf7\xdc\x66\xf1\xb2\xaf\x41\x25\xd5\xcf\xb7\x9d\xb6\xd0\xcc\x35\x46\x5a\xb2\x0a\x87\xc8\x75\x09\x00\x00\x00\x00\x00\x00\x00\xaf\xc2\x10\x3d\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0\x56\xd5\x77\xbf\x47\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x80\xa6\xa5\xbe\xb3\x73\xbb\x54\x3a\x5e\xa9\x53\x0d\x54\x74\x52\x41\xea\xd8\x9f\xb4\x11\x63\xe7\xc1\xba\xa1\xbb\xd5\x6e\xcb\x57\xfa\x2f\x56\x73\xf7\x1e\xe5\x60\xb0\xf8\x96\x87\x7c\xf1\x3d\x0b\x65\x7d\xfe\xbf\x9c\x9e\x95\x8d\x21\x84\xac\x7c\x84\xd5\xfc\xf2\xf3\xda\xcf\x00\x00\x00\xff\xff\x54\x4a\x95\xbb\xa0\x00\x01\x00"

func dicUniChardefDicBytes() ([]byte, error) {
	return bindataRead(
		_dicUniChardefDic,
		"dic/uni/chardef.dic",
	)
}

func dicUniChardefDic() (*asset, error) {
	bytes, err := dicUniChardefDicBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dic/uni/chardef.dic", size: 65696, mode: os.FileMode(420), modTime: time.Unix(1451916092, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

