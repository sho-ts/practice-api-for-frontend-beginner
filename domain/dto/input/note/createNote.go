package input

type CreateNoteInput struct {
	Title   string
	Content string
}

func NewCreateNoteInput(title string, content string) CreateNoteInput {
	return CreateNoteInput{
		Title:   title,
		Content: content,
	}
}
