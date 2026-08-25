package split

import (
	"testing"

	splitsdk "github.com/harness/harness-go-sdk/harness/split"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFMEEnvironmentChangePermissionsEditors(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceFMEEnvironment().Schema, map[string]interface{}{
		"change_permissions": []interface{}{
			map[string]interface{}{
				"are_editors_restricted": true,
				"editors": []interface{}{
					map[string]interface{}{
						"id":   "group-1",
						"name": "Editors",
						"type": "group",
					},
				},
			},
		},
	})

	permissions := expandChangePermissions(d)
	require.NotNil(t, permissions)
	require.NotNil(t, permissions.AreEditorsRestricted)
	assert.True(t, *permissions.AreEditorsRestricted)
	assert.Equal(t, []splitsdk.PermissionEntity{
		{ID: "group-1", Name: "Editors", Type: "group"},
	}, permissions.Editors)

	flattened := flattenChangePermissions(permissions)
	require.Len(t, flattened, 1)
	permissionMap := flattened[0].(map[string]interface{})
	assert.Equal(t, true, permissionMap["are_editors_restricted"])
	assert.Equal(t, []interface{}{
		map[string]interface{}{
			"id":   "group-1",
			"name": "Editors",
			"type": "group",
		},
	}, permissionMap["editors"])
}
