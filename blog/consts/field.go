package consts

const (
	// Mongodb field
	Database = "blogg"

	// Colection
	CollectionAdmin   = "admin"
	CollectionArticle = "article"
	CollectionTag     = "tag"

	// Article status
	Deleted    = -2
	Unapproved = -1
	Created    = 0
	Approverd  = 1
	Hot        = 2
)