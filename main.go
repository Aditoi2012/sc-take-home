package main

import (
	"fmt"
	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
)

func main() {
	req := &folders.FetchFolderRequest{
		OrgID: uuid.FromStringOrNil(folders.DefaultOrgID),

	}

	res, err := folders.GetAllFolders(req)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	data := res
	pageSize := 3

	// Paginate the data and generate tokens
	err = folders.PaginateData(data.Folders, pageSize)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Start with the initial token (empty string for the first page)
	token := ""

	for {
		// Retrieve the current page data using the token
		result, exists := folders.Request(token)
		if !exists {
			break
		}

		// Process the current page data
		for i:= 0; i < pageSize; i++{
			fmt.Print(result.Data[i].Name + " , ")
		}

		fmt.Println(result.Token)
		// Check if there's a next token
		token = result.Token
		if token == "" {
			break
		}
	}
}
