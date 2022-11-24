package input

type FindByNoteIdInput struct {
	Id string
}

func NewFindByNoteIdInput(id string) FindByNoteIdInput {
	return FindByNoteIdInput{
		Id: id,
	}
}
