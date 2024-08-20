package folders

import (
	"errors"
	"github.com/gofrs/uuid"
)

/*
	Improvements made:
	- There is error handling.
	- There are no unnecessary variables.
	- GetAllFolders is optimised
	- Variable names have been improved
	- The code now works
*/

func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	// retrieves the slice directly and check for any errors
	selectedFolders, err := FetchAllFoldersByOrgID(req.OrgID)
	if err != nil {
		return nil, err
	}

	// adds it to response which is returned 
	response := &FetchFolderResponse{Folders: selectedFolders}

	// returns ffr and nil as error 
	return response, nil
}

// this functions grabs the folders' pointers given the orgId that is sent
// edit: i have changed it to var so that i can use unit tests instead
var FetchAllFoldersByOrgID = func(orgID uuid.UUID) ([]*Folder, error) {
	
	if(orgID == uuid.Nil){
		return nil, errors.New("invalid OrgID")
	}
	
	folders := GetSampleData()

	resFolders := []*Folder{}
	// errors are ignored
	for _, folder := range folders {
		// this dereferences the folders pointers from GetSampleData to grab the orgId
		if folder.OrgId == orgID {
			// adds the pointer to the resFolder 
			resFolders = append(resFolders, folder)
		}
	}
	// if no folders match, an error is sent instead of empty slice
	if len(resFolders) == 0 {
		return nil, errors.New("no folders match the given OrgID")
	}
	// return resFolder and nil as error
	return resFolders, nil
}
