package interfacetype

import (
	"fmt"
	"strings"

	"github.com/cosmos/gogoproto/options"
	"github.com/cosmos/gogoproto/proto"
	"github.com/cosmos/gogoproto/protoc-gen-gogo/descriptor"
	"github.com/cosmos/gogoproto/protoc-gen-gogo/generator"
)

type interfacetype struct {
	*generator.Generator
	generator.PluginImports
}

func NewInterfaceType() *interfacetype {
	return &interfacetype{}
}

func (p *interfacetype) Name() string {
	return "interfacetype"
}

func (p *interfacetype) Init(g *generator.Generator) {
	p.Generator = g
}

func GetInterfaceType(message *descriptor.DescriptorProto) string {
	if message == nil {
		return ""
	}
	if message.Options != nil {
		v, err := proto.GetExtension(message.Options, options.E_InterfaceType)
		if err == nil && v.(*string) != nil {
			return *(v.(*string))
		}
	}
	return ""
}

func (p *interfacetype) Generate(file *generator.FileDescriptor) {
	p.PluginImports = generator.NewPluginImports(p.Generator)

	for _, message := range file.Messages() {
		iface := GetInterfaceType(message.DescriptorProto)
		if len(iface) == 0 {
			continue
		}
		handleNonPointer := true
		if iface[0] == '*' {
			iface = iface[1:]
			handleNonPointer = false
		}
		if len(message.OneofDecl) != 1 {
			panic("interfacetype only supports messages with exactly one oneof declaration")
		}
		for _, field := range message.Field {
			if idx := field.OneofIndex; idx == nil || *idx != 0 {
				panic("all fields in interfacetype message must belong to the oneof")
			}
		}

		ifacePackage, ifaceName := splitCPackageType(iface)
		ifaceRef := ifaceName
		if len(ifacePackage) != 0 {
			imp := p.PluginImports.NewImport(ifacePackage).Use()
			ifaceRef = fmt.Sprintf("%s.%s", imp, ifaceName)
		}

		ccTypeName := generator.CamelCaseSlice(message.TypeName())
		p.P(`func (this *`, ccTypeName, `) Get`, ifaceName, `() `, ifaceRef, ` {`)
		p.In()
		for _, field := range message.Field {
			fieldname := p.GetOneOfFieldName(message, field)
			if fieldname == "Value" {
				panic("cannot have a onlyone message " + ccTypeName + " with a field named Value")
			}
			p.P(`if x := this.Get`, fieldname, `(); x != nil {`)
			p.In()
			p.P(`return x`)
			p.Out()
			p.P(`}`)
		}
		p.P(`return nil`)
		p.Out()
		p.P(`}`)
		p.P(``)
		p.P(`func (this *`, ccTypeName, `) Set`, ifaceName, `(value `, ifaceRef, `) error {`)
		p.In()
		p.P(`if value == nil {`)
		p.In()
		p.P(`this.`, p.GetFieldName(message, message.Field[0]), ` = nil`)
		p.P(`return nil`)
		p.Out()
		p.P("}")
		p.P(`switch vt := value.(type) {`)
		p.In()
		for _, field := range message.Field {
			oneofName := p.GetFieldName(message, field)
			structName := p.OneOfTypeName(message, field)
			goTyp, _ := p.GoType(message, field)
			p.P(`case `, goTyp, `:`)
			p.In()
			p.P(`this.`, oneofName, ` = &`, structName, `{vt}`)
			p.P("return nil")
			p.Out()
			if handleNonPointer {
				// Handle non-pointer case
				if goTyp[0] == '*' {
					p.P(`case `, goTyp[1:], `:`)
					p.In()
					p.P(`this.`, oneofName, ` = &`, structName, `{&vt}`)
					p.P("return nil")
					p.Out()
				}
			}
		}
		p.P(`}`)
		p.P(`return fmt.Errorf("can't encode value of type %T as message `, ccTypeName, `", value)`)
		p.Out()
		p.P(`}`)
		p.P(``)
	}
}

func splitCPackageType(ctype string) (packageName string, typ string) {
	ss := strings.Split(ctype, ".")
	if len(ss) == 1 {
		return "", ctype
	}
	packageName = strings.Join(ss[0:len(ss)-1], ".")
	typeName := ss[len(ss)-1]
	return packageName, typeName
}

func init() {
	generator.RegisterPlugin(NewInterfaceType())
}
