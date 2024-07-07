package domain

type Folder struct {
	ID    int
	Title string
}

func NewFolder(id int, title string) *Folder {
	return &Folder{
		ID:    id,
		Title: title,
	}
}

func (f *Folder) GetTitle() string {
	return f.Title
}
