package SpellNote

type Note struct {
	ID          int    `json:"id" db:"id"`
	Description string `json:"description" db:"description"`
	UserId      int    `json:"userId" db:"user_id"`
}

type NoteInput struct {
	Description string `json:"description" db:"description"`
}
