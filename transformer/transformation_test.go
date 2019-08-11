package transformer

import (
	"testing"

	"github.com/rantav/go-archetype/inputs"
	"github.com/rantav/go-archetype/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransformationsTransform(t *testing.T) {
	assert := assert.New(t)

	// empty transformers list
	ts := Transformations{}
	file, err := ts.Transform(types.File{
		Contents: "x",
	})
	require.NoError(t, err)
	assert.Equal("x", file.Contents)

	// Single replacer
	ts = Transformations{
		transformers: []Transformer{newTextReplacer(
			transformationSpec{
				Pattern:     "x",
				Replacement: "y",
				Files:       []types.FilePattern{"*.go"},
			})},
	}
	file, err = ts.Transform(types.File{
		Contents: "x",
		Path:     "hello.go",
	})
	require.NoError(t, err)
	assert.Equal("y", file.Contents)

	// A file that doesn't match
	ts = Transformations{
		transformers: []Transformer{newTextReplacer(
			transformationSpec{
				Pattern:     "x",
				Replacement: "y",
				Files:       []types.FilePattern{"hello.go"},
			})},
	}
	file, err = ts.Transform(types.File{
		Path:     "go.away",
		Contents: "x",
	})
	require.NoError(t, err)
	assert.Equal("x", file.Contents)
}

func TestTransformationsTemplate(t *testing.T) {
	// empty transformers list
	ts := Transformations{
		transformers: []Transformer{
			newTextReplacer(transformationSpec{}),
		},
		prompters: []inputs.Prompter{
			inputs.NewPrompt(inputs.InputSpec{Type: "text"}),
		},
	}
	err := ts.Template()
	require.NoError(t, err)
}

func TestTransformationsMatched(t *testing.T) {
	assert := assert.New(t)
	assert.True(matched("hello.go", []types.FilePattern{"hello.go"}, false))
	assert.True(matched("all/hello.go", []types.FilePattern{"all/"}, true))
}
