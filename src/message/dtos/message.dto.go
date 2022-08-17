package dtos

type MessageDto struct {
	// required: true
	Body string `form:"body" json:"body" xml:"body"  binding:"required,min=1,max=300"`
}
