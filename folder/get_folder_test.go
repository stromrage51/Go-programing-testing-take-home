package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	// "github.com/stretchr/testify/assert"
)

// feel free to change how the unit test is structured
func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()

	// Sample organization IDs
	orgID1, _ := uuid.NewV4()
	orgID2, _ := uuid.NewV4()

	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{

		{
			name:  "Get folders for orgID1",
			orgID: orgID1,
			folders: []folder.Folder{
				{Name: "matt", Paths: "matt", OrgId: orgID1},
				{Name: "sydeny", Paths: "matt.sydeny", OrgId: orgID1},
				{Name: "ben", Paths: "matt.sydeny.ben", OrgId: orgID1},
				{Name: "delta", Paths: "matt.delta", OrgId: orgID1},
				{Name: "rex", Paths: "rex", OrgId: orgID2}, //Different orgID to test
			},
			want: []folder.Folder{
				{Name: "matt", Paths: "matt", OrgId: orgID1},
				{Name: "sydeny", Paths: "matt.sydeny", OrgId: orgID1},
				{Name: "ben", Paths: "matt.sydeny.ben", OrgId: orgID1},
				{Name: "delta", Paths: "matt.delta", OrgId: orgID1},
			},
		},
		{
			name:  "Get folders for orgID2 (no folders found)",
			orgID: orgID2,
			folders: []folder.Folder{
				{Name: "matt", Paths: "matt", OrgId: orgID1},
				{Name: "sydeny", Paths: "matt.sydeny", OrgId: orgID1},
			},
			want: []folder.Folder{}, // Expecting no folders
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// f := folder.NewDriver(tt.folders)
			// get := f.GetFoldersByOrgID(tt.orgID)

		})
	}
}
