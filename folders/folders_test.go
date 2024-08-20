package folders

import (
	"errors"
	"testing"
	// "github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func MockFetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {

	mockFolder1 := &Folder{
		OrgId: orgID,
		Name:  "Folder1",
	}
	mockFolder2 := &Folder{
		OrgId: orgID,
		Name:  "Folder2",
	}

	if orgID == uuid.Nil {
		return nil, errors.New("invalid OrgID")
	}
	if orgID == uuid.FromStringOrNil("00000000-0000-0000-0000-000000000001") {
		return nil, errors.New("no folders match the given OrgID")
	}
	return []*Folder{mockFolder1, mockFolder2}, nil
}

// test case for no matching folders
func TestGetAllFolders_NoMatchingFolders(t *testing.T) {

	req := &FetchFolderRequest{
		OrgID: uuid.FromStringOrNil("00000000-0000-0000-0000-000000000001"),
	}

	FetchAllFoldersByOrgID = MockFetchAllFoldersByOrgID

	res, err := GetAllFolders(req)

	assert.NotNil(t, err)
	assert.Nil(t, res)
	assert.Equal(t, "no folders match the given OrgID", err.Error())
}

// test case for successful retrieval
func TestGetAllFolders_Success(t *testing.T) {

	orgID := uuid.Must(uuid.NewV4())
	req := &FetchFolderRequest{
		OrgID: orgID,
	}

	FetchAllFoldersByOrgID = MockFetchAllFoldersByOrgID

	res, err := GetAllFolders(req)

	assert.Nil(t, err)

	assert.NotNil(t, res)
	assert.Equal(t, 2, len(res.Folders))
	assert.Equal(t, "Folder1", res.Folders[0].Name)
	assert.Equal(t, "Folder2", res.Folders[1].Name)
}

// Test case for error when OrgID is nil
func TestGetAllFolders_InvalidOrgID(t *testing.T) {
	req := &FetchFolderRequest{
		OrgID: uuid.Nil,
	}

	FetchAllFoldersByOrgID = MockFetchAllFoldersByOrgID
	res, err := GetAllFolders(req)

	assert.NotNil(t, err)
	assert.Nil(t, res)
	assert.Equal(t, "invalid OrgID", err.Error())
}


func Test_GetAllFolders(t *testing.T) {
	t.Run("TestGetAllFolders_Success", TestGetAllFolders_Success)
	t.Run("TestGetAllFolders_InvalidOrgID", TestGetAllFolders_InvalidOrgID)
	t.Run("TestGetAllFolders_NoMatchingFolders", TestGetAllFolders_NoMatchingFolders)
}
