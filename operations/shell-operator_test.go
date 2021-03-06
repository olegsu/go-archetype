package operations

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewShellOperator(t *testing.T) {
	assert := assert.New(t)

	o := newShellOperator(OperationSpec{Sh: []string{"hello"}})
	assert.NotNil(o)
	assert.IsType(&shellOperation{}, o)
}

func TestShellOperatorTemplate(t *testing.T) {
	assert := assert.New(t)

	o := newShellOperator(OperationSpec{Sh: []string{"hello {{.source}}"}})
	require.NotNil(t, o)
	vars := map[string]string{
		"source": "world",
	}
	err := o.Template(vars)
	require.NoError(t, err)

	assert.Equal("hello world", o.sh[0])
}

func TestShellOperatorOperate(t *testing.T) {
	o := newShellOperator(OperationSpec{Sh: []string{"echo hello"}})
	require.NotNil(t, o)
	err := o.Operate()
	require.NoError(t, err)
}
