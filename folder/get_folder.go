package folder

import (
	"strings"

	"github.com/gofrs/uuid"
)

// You will need to implement the following:

// 1. A method to get all child folders of a given folder.
// 2. The method should return a list of all child folders.
// 3. Implement any necessary error handling (e.g. invalid orgID, invalid paths, etc).

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res

}

// 1. A method to get all child folders of a given folder.
func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder {

	// Gather all folders with that has orgID
	folders := f.GetFoldersByOrgID(orgID)

	// Locating folder with the extract given name
	parentFolder := f.findFolderByName(folders, name)

	// 3. the error checking to check if folder exists
	if parentFolder == nil {
		return []Folder{}
	}

	// Returning all child folders
	return f.collectChildFolders(folders, parentFolder.Paths)
}

// 2. The method should return a list of all child folders.
func (f *driver) findFolderByName(folders []Folder, name string) *Folder {
	for _, folder := range folders {
		if folder.Name == name {
			return &folder
		}
	}
	return nil
}

// Gathering the child folders on path prefix
func (f *driver) collectChildFolders(folders []Folder, parentPath string) []Folder {
	var childFolders []Folder
	for _, folder := range folders {
		if folder.Paths != parentPath && strings.HasPrefix(folder.Paths, parentPath+".") {
			childFolders = append(childFolders, folder)
		}
	}
	return childFolders
}
