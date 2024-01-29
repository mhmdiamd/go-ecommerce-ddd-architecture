package utility

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestToken(t *testing.T) {
  t.Run("success", func(t *testing.T) {
    publicId := uuid.NewString()
    tokenString, err := GenerateToken(publicId, "user", "iniSecret")
    require.Nil(t, err)
    require.NotNil(t, tokenString)
    fmt.Println(tokenString)
  })
}


func TestValidateToken(t *testing.T) {
  t.Run("success validate token", func(t *testing.T) {
    publicId := uuid.NewString()
    role := "user"
    tokenString, err := GenerateToken(publicId, "user", "iniSecret")
    require.Nil(t, err)
    require.NotNil(t, tokenString)

    jwtId, jwtRole, err := ValidateToken(tokenString, "iniSecret")
    require.Nil(t, err)
    require.Equal(t, publicId, jwtId)
    require.Equal(t, role, jwtRole)

    fmt.Println(tokenString)
  })
}
