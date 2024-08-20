package folders

import (
	"github.com/gofrs/uuid"
)

	/* 
		Improvements:
		- There is no current error handling. This can be improved.
		- There are unnecessary variables.
		- There is an unncessary loop in GetAllFolders since you can work directly with the 
		  result of FetchAllFoldersByOrgID
		- Variable names can be improved
	*/

func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {

	// not used variables
	var (
		err error
		f1  Folder
		fs  []*Folder
	)
	// creates an empty slice of Folder type struct
	f := []Folder{}
	// resulting pointers from the method below are stored in r and errors are ignored
	r, _ := FetchAllFoldersByOrgID(req.OrgID)
	// k is index which is not used and v is the pointer for the folder in slice r
	for k, v := range r {
		// de-references every Folder and adds it to the folder f
		f = append(f, *v)
	}
	// creates an empty slice of pointers to Folder
	var fp []*Folder
	// k1 index is unused and every element within the slice is stored as v1
	for k1, v1 := range f {
		// takes the address of each folder value and appends to the slice declared above
		fp = append(fp, &v1)
	}
	// declares ffr which will hold a pointer to the FetchFolderResponse object
	var ffr *FetchFolderResponse
	// new FetchFolderResponse object is created, with its Folders field set to the slice fp
	// the address is stored in ffr
	ffr = &FetchFolderResponse{Folders: fp}

	// returns ffr and nil as error 
	return ffr, nil
}

// this functions grabs the folders' pointers given the orgId that is sent
func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	folders := GetSampleData()

	resFolder := []*Folder{}
	// errors are ignored
	for _, folder := range folders {
		// this dereferences the folders pointers from GetSampleData to grab the orgId
		if folder.OrgId == orgID {
			// adds the pointer to the resFolder 
			resFolder = append(resFolder, folder)
		}
	}
	// return resFolder and nil as error
	return resFolder, nil
}
