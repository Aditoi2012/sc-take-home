package folders

import "github.com/gofrs/uuid"

type FetchFolderRequest struct {
	OrgID uuid.UUID
}

type FetchFolderResponse struct {
	Folders []*Folder
}

type PaginatedResult struct {
    Data  []*Folder 
    Token string 
}