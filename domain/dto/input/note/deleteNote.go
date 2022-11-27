package input

type DeleteNoteInput struct {
	Id string
}

func NewDeleteNoteInput(
	id string,
) DeleteNoteInput {
	return DeleteNoteInput{
		Id: id,
	}
}
