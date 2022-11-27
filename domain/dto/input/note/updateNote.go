package input

type UpdateNoteInput struct {
	Id      string
	Title   string
	Content string
}

func NewUpdateNoteInput(
	id string,
	title string,
	content string,
) UpdateNoteInput {
	return UpdateNoteInput{
		Id:      id,
		Title:   title,
		Content: content,
	}
}
