package mapper

import (
	"backend/api"
	"backend/domain/note/model"
	"github.com/oapi-codegen/nullable"
)

func MapToNoteApi(note *model.Notes) api.Note {
	return api.Note{
		Id:     note.ID,
		UserId: note.UserID,
		Title:  note.Title,
		Content: api.Content{
			Nullable: toNullable(note.Content),
		},
		IsArchived: note.IsArchived,
		ParentId:   note.ParentID,
		CoverImage: api.CoverImage{
			Nullable: toNullable(note.CoverImage),
		},
		Icon: api.Icon{Nullable: toNullable(note.Icon)},
	}
}

func toNullable(s *string) nullable.Nullable[string] {
	if s == nil {
		return nullable.NewNullNullable[string]()
	}
	return nullable.NewNullableWithValue(*s)
}
