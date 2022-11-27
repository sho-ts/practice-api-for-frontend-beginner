package output

type DeleteNoteOutput struct {
	Id string `json:"id"`
}

func NewDeleteNoteOutput(
	id string,
) DeleteNoteOutput {
	return DeleteNoteOutput{
		Id: id,
	}
}
