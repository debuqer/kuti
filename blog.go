package main

type Post struct {
	Title         string
	Content       string
	Date          string
	EstimatedTime int
}

type Blog struct {
	Config Config
	Posts  []Post
}

func (b *Blog) fetch(_conf Config) error {
	post := Post{
		"Hello",
		"Hoy",
		"12/02/1402",
		1,
	}

	posts := make([]Post, 0)

	posts = append(posts, post)

	b.Config = _conf
	b.Posts = posts

	return nil
}
