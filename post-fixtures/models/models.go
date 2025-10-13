package models

type User struct {
	ID        int
	FirstName string
	LastName  string
}

type Post struct {
	ID    int
	User  *User
	Title string
	Body  string
}

func (p *Post) PostName() string {
	return "A post titled '" + p.Title + "' " +
		"by " + p.User.FirstName + " " + p.User.LastName
}

func main()
