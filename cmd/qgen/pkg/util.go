package pkg

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func init() {
	fmt.Println("init pkg")
}

const (
	NO_TEMPLATE = iota
	TEMPLATE
)

func Generate(target, template, name string) error {
	newGenerator(target, template, name).generate()
	return nil
}

// q-gen factory
func newGenerator(target, template, name string) generator {
	var mode int
	if template == "" {
		mode = NO_TEMPLATE
	} else {
		mode = TEMPLATE
	}
	switch mode {
	case NO_TEMPLATE:
		return newNoTemplateGenerator(target, name)
	case TEMPLATE:
		return newTemplateGenerator(target, template, name)
	default:
		panic("invalid mode")
	}
}

type generator interface {
	generate() error
}

type noTemplateGenerator struct {
	target string
	name   string
}

func newNoTemplateGenerator(target, name string) *noTemplateGenerator {
	return &noTemplateGenerator{
		target: target,
		name:   name,
	}
}

func (g *noTemplateGenerator) generate() error {
	return generate(ANS_TEMPLATE, ANS_TEST_TEMPLATE, g.target, g.name)
}

type templateGenerator struct {
	target   string
	template string
	name     string
}

func newTemplateGenerator(target, template, name string) *templateGenerator {
	return &templateGenerator{
		target:   target,
		template: template,
		name:     name,
	}
}

func (g *templateGenerator) generate() error {
	ans, test := path.Join(g.template, "ans.go"), path.Join(g.template, "ans_test.go")
	ans_template, err := os.ReadFile(ans)
	if err != nil {
		return err
	}
	test_template, err := os.ReadFile(test)
	if err != nil {
		return err
	}
	return generate(string(ans_template), string(test_template), g.target, g.name)
}

func generate(ans_t, test_t, target, name string) error {
	// make dir "q{{name}}" at target
	target_path := path.Join(target, "q"+name)
	err := os.Mkdir(target_path, 0755)
	if err != nil {
		return err
	}
	// make file "ans.go" and "ans_test.go" at target_path
	ans_path := path.Join(target_path, "ans.go")
	test_path := path.Join(target_path, "ans_test.go")

	ans, err := os.Create(ans_path)
	if err != nil {
		return err
	}
	defer ans.Close()
	test, err := os.Create(test_path)
	if err != nil {
		return err
	}
	defer test.Close()

	// write template to ans.go
	err = write(ans, strings.ReplaceAll(ans_t, "{{.Name}}", name))
	if err != nil {
		return err
	}
	// write template to ans_test.go
	err = write(test, strings.ReplaceAll(test_t, "{{.Name}}", name))
	if err != nil {
		return err
	}
	fmt.Printf("Generated %s\n", target_path)
	fmt.Printf("2 Files Generated:\n- %s,\n- %s\n", ans_path, test_path)
	return nil
}

func write(f *os.File, s string) error {
	fmt.Printf("Writing in to file %s ... \n", f.Name())
	_, err := f.WriteString(s)
	return err
}
