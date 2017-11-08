package golang

import (
	"fmt"
	"go/token"
	"io"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/reflect/errawr-gen/doc"
	"github.com/serenize/snaker"
)

const (
	public  = "github.com/reflect/errawr-go"
	private = "github.com/reflect/errawr-go/impl"
)

func argumentGoName(name string) string {
	candidate := snaker.SnakeToCamelLower(name)
	if token.Lookup(candidate).IsKeyword() {
		candidate += "_"
	}

	return candidate
}

func argumentType(a *doc.DocumentErrorArgument) jen.Code {
	if a == nil {
		return jen.String()
	}

	switch a.Type {
	case "string":
		return jen.String()
	case "number":
		return jen.Float64()
	case "integer":
		return jen.Int64()
	case "boolean":
		return jen.Bool()
	case "list<string>":
		return jen.Index().String()
	case "list<number>":
		return jen.Index().Float64()
	case "list<integer>":
		return jen.Index().Int64()
	case "list<boolean>":
		return jen.Index().Bool()
	default:
		return jen.String()
	}
}

type Error struct {
	Section    Section
	Name       string
	GoName     string
	Definition doc.DocumentError
}

type Section struct {
	Name       string
	GoName     string
	Definition doc.DocumentSection
}

type FileGenerator struct {
	pkg string
	doc *doc.Document
	f   *jen.File
}

func (fg *FileGenerator) init() {
	fg.f.PackageComment(fmt.Sprintf("Package %s contains errors for the domain %q.", fg.pkg, fg.doc.Domain.Key))
	fg.f.PackageComment("")
	fg.f.PackageComment("This file is automatically generated by errawr-gen. Do not modify it.")

	// Create an interface type inheriting the public Errawr error interface.
	fg.f.Comment("Error is the type of all errors generated by this package.")
	fg.f.Type().Id("Error").Interface(jen.Qual(public, "Error"))

	// Create a type for the domain.
	fg.f.Comment("Domain is the general domain in which all errors in this package belong.")
	fg.f.Var().Id("Domain").Op("=").Op("&").Qual(private, "ErrorDomain").Values(jen.Dict{
		jen.Id("Key"):   jen.Lit(fg.doc.Domain.Key),
		jen.Id("Title"): jen.Lit(fg.doc.Domain.Title),
	})
	fg.f.Line()
}

func (fg *FileGenerator) def(def Error) {
	// Create the builder type.
	builderTypeName := fmt.Sprintf(`%s%sBuilder`, def.Section.GoName, def.GoName)

	fg.f.Commentf(`%s is a builder for %q errors.`, builderTypeName, def.Name)
	fg.f.Type().Id(builderTypeName).Struct(jen.Id("arguments").Qual(private, "ErrorArguments"))
	fg.f.Line()

	// Create builder methods for all optional arguments.
	for _, a := range def.Definition.OrderedArguments {
		if !a.Argument.IsOptional() {
			continue
		}

		withFuncName := fmt.Sprintf(`With%s`, snaker.SnakeToCamel(a.Name))

		description := a.Argument.Description
		if description == "" {
			description = fmt.Sprintf("the %q argument", a.Name)
		}

		fg.f.Commentf(`%s sets %s in this builder.`, withFuncName, description)
		fg.f.Func().Params(
			jen.Id("b").Op("*").Id(builderTypeName),
		).Id(withFuncName).Params(
			jen.Id("value").Add(argumentType(a.Argument)),
		).Op("*").Id(builderTypeName).Block(
			jen.Id("b").Dot("arguments").Index(jen.Lit(a.Name)).Dot("Set").Call(jen.Id("value")),
			jen.Return(jen.Id("b")),
		)
		fg.f.Line()
	}

	// Create final Build() method.
	fg.f.Commentf(`Build creates the error for the code %q from this builder.`, def.Name)
	fg.f.Func().Params(
		jen.Id("b").Op("*").Id(builderTypeName),
	).Id("Build").Params().Id("Error").BlockFunc(func(g *jen.Group) {
		// Add runtime argument validation.
		for _, a := range def.Definition.OrderedArguments {
			if a.Argument == nil || len(a.Argument.Validators) == 0 {
				continue
			}

			for _, validator := range a.Argument.Validators {
				g.Id("b").Dot("arguments").Index(jen.Lit(a.Name)).Dot("Validate").Call(jen.Lit(validator))
			}
			g.Line()
		}

		g.Id("description").Op(":=").Op("&").Qual(private, "ErrorDescription").Values(jen.Dict{
			jen.Id("Friendly"):  jen.Lit(strings.TrimSpace(def.Definition.Description.Friendly)),
			jen.Id("Technical"): jen.Lit(strings.TrimSpace(def.Definition.Description.Technical)),
		})
		g.Line()

		g.Return(jen.Op("&").Qual(private, "Error").Values(jen.Dict{
			jen.Id("ErrorDomain"):      jen.Id("Domain"),
			jen.Id("ErrorSection"):     jen.Id(fmt.Sprintf(`%sSection`, def.Section.GoName)),
			jen.Id("ErrorCode"):        jen.Lit(def.Name),
			jen.Id("ErrorTitle"):       jen.Lit(def.Definition.Title),
			jen.Id("ErrorDescription"): jen.Id("description"),
			jen.Id("ErrorArguments"):   jen.Id("b").Dot("arguments"),
		}))
	})
	fg.f.Line()

	// Create constructors.
	constructorFuncName := fmt.Sprintf(`New%s`, builderTypeName)
	params := func(g *jen.Group) {
		for _, a := range def.Definition.OrderedArguments {
			if a.Argument.IsOptional() {
				continue
			}

			g.Id(argumentGoName(a.Name)).Add(argumentType(a.Argument))
		}
	}

	fg.f.Commentf(`%s creates a new error builder for the code %q.`, constructorFuncName, def.Name)
	fg.f.Func().Id(constructorFuncName).ParamsFunc(params).Op("*").Id(builderTypeName).Block(
		jen.Return(jen.Op("&").Id(builderTypeName).Values(jen.Dict{
			jen.Id("arguments"): jen.Qual(private, "ErrorArguments").Values(jen.DictFunc(func(d jen.Dict) {
				for _, a := range def.Definition.OrderedArguments {
					d[jen.Lit(a.Name)] = jen.Qual(private, "NewErrorArgument").Call(jen.Do(func(s *jen.Statement) {
						if a.Argument.IsOptional() {
							s.Lit(a.Argument.Default)
						} else {
							s.Id(argumentGoName(a.Name))
						}
					}), jen.LitFunc(func() interface{} {
						if a.Argument == nil {
							return ""
						}

						return a.Argument.Description
					}))
				}
			})),
		})),
	)
	fg.f.Line()

	defaultFuncName := fmt.Sprintf(`New%s%s`, def.Section.GoName, def.GoName)

	fg.f.Commentf(`%s creates a new error with the code %q.`, defaultFuncName, def.Name)
	fg.f.Func().Id(defaultFuncName).ParamsFunc(params).Id("Error").Block(
		jen.Return(
			jen.Id(constructorFuncName).CallFunc(func(g *jen.Group) {
				for _, a := range def.Definition.OrderedArguments {
					if a.Argument.IsOptional() {
						continue
					}

					g.Id(argumentGoName(a.Name))
				}
			}).Dot("Build").Call(),
		),
	)
	fg.f.Line()
}

func (fg *FileGenerator) section(section Section) {
	// Create a type for the section.
	sectionTypeName := fmt.Sprintf(`%sSection`, section.GoName)

	fg.f.Commentf(`%s defines a section of errors with the following scope:`, sectionTypeName)
	fg.f.Comment(section.Definition.Title)
	fg.f.Var().Id(sectionTypeName).Op("=").Op("&").Qual(private, "ErrorSection").Values(jen.Dict{
		jen.Id("Key"):   jen.Lit(section.Name),
		jen.Id("Title"): jen.Lit(section.Definition.Title),
	})
	fg.f.Line()

	for name, def := range section.Definition.Errors {
		fg.def(Error{
			Section:    section,
			Name:       name,
			GoName:     snaker.SnakeToCamel(name),
			Definition: def,
		})
	}
}

func (fg *FileGenerator) all() {
	for name, section := range fg.doc.Sections {
		fg.section(Section{
			Name:       name,
			GoName:     snaker.SnakeToCamel(name),
			Definition: section,
		})
	}
}

func (fg *FileGenerator) Render(w io.Writer) error {
	return fg.f.Render(w)
}

func NewFileGenerator(pkg string, document *doc.Document) *FileGenerator {
	fg := &FileGenerator{
		pkg: pkg,
		doc: document,
		f:   jen.NewFile(pkg),
	}
	fg.init()
	fg.all()

	return fg
}
