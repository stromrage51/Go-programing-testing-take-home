package folder

import (
	"errors"
	"strings"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {

	// obtain all the necassry folders
	folders := f.folders
	// fmt.Printf("!folders is a %T\n", folders)

	sourceFolder, destinationFolder, err := f.findFolders(folders, name, dst)
	if err != nil {
		return nil, err
	}

	//////////////////////////////////////////////

	//Error checking

	// checking if destination folder is a child folder

	if strings.HasPrefix(destinationFolder.Paths, sourceFolder.Paths+".") {
		return nil, errors.New("cannot move a folder to a child of itself")
	}

	////////////////////////////////////////////
	// // Create a new path for folder and moving it.
	newPath := destinationFolder.Paths + "." + sourceFolder.Name
	f.updateFolderPaths(folders, sourceFolder, newPath)

	return []Folder{}, nil
}

func (f *driver) findFolders(folders []Folder, name, dst string) (*Folder, *Folder, error) {
	var sourceFolder *Folder

	var destinationFolder *Folder

	// Finding the sourceFolder
	// Finding the destinationFolder
	for _, folder := range folders {
		switch folder.Name {
		case name:
			sourceFolder = &folder
		case dst:
			destinationFolder = &folder
		}

		// For when both folders are found
		if sourceFolder != nil && destinationFolder != nil {
			break
		}
	}

	// Checking if the destination folder exists and source folder exits
	if sourceFolder == nil {
		return nil, nil, errors.New("destination folder does not exist")
	}

	if destinationFolder == nil {
		return nil, nil, errors.New("source folder does not exist")
	}

	return sourceFolder, destinationFolder, nil
}

// Create a new path for folder and moving it.
func (f *driver) updateFolderPaths(folders []Folder, sourceFolder *Folder, newPath string) {
	sourcePath := sourceFolder.Paths
	sourceFolder.Paths = newPath

	// call function to update childern paths
	f.updateChildPaths(folders, sourcePath, newPath)
}

func (f *driver) updateChildPaths(folders []Folder, oldPath, newPath string) {
	for index := range folders {
		if strings.HasPrefix(folders[index].Paths, oldPath+".") {
			folders[index].Paths = newPath + folders[index].Paths[len(oldPath):]
			f.updateChildPaths(folders, folders[index].Paths, newPath+"."+folders[index].Name)
		}
	}
}
