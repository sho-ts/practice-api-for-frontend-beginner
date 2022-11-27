package output

type UpdateNoteOutput struct {
	Id string `json:"id"`
}

func NewUpdateNoteOutput(
	id string,
) UpdateNoteOutput {
	return UpdateNoteOutput{
		Id: id,
	}
}
