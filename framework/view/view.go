package view

import (
	_ "embed"
	"fmt"

	"github.com/livebud/bud/framework"
	"github.com/livebud/bud/framework/transform/transformrt"
	"github.com/livebud/bud/internal/gotemplate"
	"github.com/livebud/bud/package/budfs"
	"github.com/livebud/bud/package/gomod"
)

//go:embed view.gotext
var template string

var generator = gotemplate.MustParse("framework/view/view.gotext", template)

// Generate the view from state
func Generate(state *State) ([]byte, error) {
	return generator.Generate(state)
}

func New(module *gomod.Module, transform *transformrt.Map, flag *framework.Flag) *Generator {
	return &Generator{
		flag:      flag,
		module:    module,
		transform: transform,
	}
}

type Generator struct {
	flag      *framework.Flag
	module    *gomod.Module
	transform *transformrt.Map
}

func (c *Generator) GenerateFile(fsys budfs.FS, file *budfs.File) error {
	state, err := Load(fsys.Context(), fsys, c.module, c.transform, c.flag)
	if err != nil {
		return err
	}
	code, err := Generate(state)
	if err != nil {
		return err
	}
	file.Data = code
	return nil
}

func (g *Generator) ServeFile(fsys budfs.FS, file *budfs.File) error {
	fmt.Println("serving", file.Target())
	// return g.GenerateFile(fsys, file)
	return nil
}
