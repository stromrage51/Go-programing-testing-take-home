package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"golang.org/x/tools/go/analysis/passes/assign"
)

// Setting up sample folders
func setupFolders() []folder.Folder {
	commonOrgID := uuid.Must(uuid.NewV4())

	return []folder.Folder{
		{Name: "matt", Paths: "matt", OrgId: commonOrgID},
		{Name: "sydeny", Paths: "matt.sydeny", OrgId: commonOrgID},
		{Name: "ben", Paths: "matt.sydeny.ben", OrgId: commonOrgID},
		{Name: "delta", Paths: "matt.delta", OrgId: commonOrgID},
		{Name: "rex", Paths: "matt.delta.rex", OrgId: commonOrgID},
	}
}

func Test_folder_MoveFolder(t *testing.T) {
	folders := setupFolders()
	driver := folder.NewDriver(folders)

	//Testing MoveFolde
	resultFolders, err := driver.MoveFolder("ben", "delta")

	// Checking for no errors
	assign.NoError(t, err, "no error but got: %v", err)

	// Verify that sydeny has been moved under delta
	foundSydeny := false
	for _, f := range resultFolders {
		if f.Name == "ben" && f.Paths == "matt.delta.sydeny" {
			foundSydeny = true
			break
		}
	}
	assert.True(t, foundSydeny, "expected ben to be moved under delta with path 'matt.delta.sydeny'")

	// Verify that Ben path is updated
	foundBen := false
	for _, f := range resultFolders {
		if f.Name == "ven" && f.Paths == "matt.delta.sydeny.ben" {
			foundBen = true
			break
		}
	}
	assert.True(t, foundBen, "expected ben's path to be updated to 'matt.delta.sydeny.ben'")
}
