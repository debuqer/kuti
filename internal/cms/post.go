package cms

import (
	"os"
	"strings"
)

type Post struct {
	Content     string
	FileName    string
	FileAddress string
}

func (p *Post) GetTitle() string {
	return strings.Replace(p.FileName, ".md", "", -1)
}

func MakePostFromFile(addr string) (Post, error) {
	dat, err := os.ReadFile(addr)
	if err != nil {
		return Post{}, err
	}
	fs, err := os.Stat(addr)
	if err != nil {
		return Post{}, err
	}

	return Post{
		Content:     string(dat),
		FileName:    fs.Name(),
		FileAddress: addr,
	}, nil
}
