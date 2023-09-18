package models

type SID struct {
	SID string `json:"sid"`
}

type GetAllNotesRequest struct {
	SID
}

type GetAllNotesResponse struct {
	Notes []Note `json:"notes"`
}

type AddNoteRequest struct {
	SID
	Note string `json:"note"`
}

type AddNoteResponse struct {
	ID uint32 `json:"id"`
}

type DeleteNoteRequest struct {
	SID
	ID uint32 `json:"id"`
}
