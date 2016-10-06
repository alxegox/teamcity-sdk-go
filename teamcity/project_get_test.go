package teamcity

import (
  "github.com/stretchr/testify/assert"
  "github.com/stretchr/testify/require"
  "github.com/umweltdk/teamcity/types"
  "testing"
)

func TestClientGetProject(t *testing.T) {
  client, err := NewRealTestClient(t)
  require.NoError(t, err, "Expected no error")

  config, err := client.GetProject("Empty")
  require.NoError(t, err, "Expected no error")
  require.NotNil(t, config, "Create to return config")

  assert.Equal(t, types.Properties{
    "env.MUH": types.Property{
      Value: client.VersionParameterValue(t, "env.MUH"),
      Spec: &types.PropertySpec{
        Label: "Muh value",
        Description: "The Muh value that does all the Muhing",
        Type: types.PasswordType{},
      },
    },
  }, config.Parameters, "Parameters")
}

func TestClientGetProjectMissing(t *testing.T) {
  client, err := NewRealTestClient(t)
  require.NoError(t, err, "Expected no error")

  config, err := client.GetProject("Empt")
  require.NoError(t, err, "Expected no error")
  require.Nil(t, config, "Expected no config")
}