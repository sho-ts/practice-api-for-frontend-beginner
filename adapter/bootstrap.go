package adapter

import (
	c "note-app/adapter/controller"
	ni "note-app/adapter/interactor/note"
	r "note-app/repository"
)

var nr = r.NewNoteRepository()

var NoteController = c.NewNoteController(
	ni.NewCreateNoteInteractor(nr),
	ni.NewFindByNoteIdInteractor(nr),
	ni.NewFindAllNoteInteractor(nr),
	ni.NewUpdateNoteInteractor(nr),
	ni.NewDeleteNoteInteractor(nr),
)
