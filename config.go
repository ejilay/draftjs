package draftjs

type Config struct {
	entityDecorators map[string]*Descriptor
	blockMap         map[string]*Descriptor
	styleMap         map[string]*Descriptor
}

type Descriptor struct {
	Type      string
	Element   string
	Wrapper   string
	Decorator Decorator
}

func (config *Config) SetEntityDecorator(descriptor *Descriptor) {
	if descriptor == nil || descriptor.Type == "" {
		return
	}
	if config.entityDecorators == nil {
		config.entityDecorators = make(map[string]*Descriptor)
	}
	config.entityDecorators[descriptor.Type] = descriptor
}

func (config *Config) GetEntityDecorator(descriptorType string) *Descriptor {
	return GetDescriptorFromMap(descriptorType, config.entityDecorators)
}

func (config *Config) SetBlockMapElement(descriptor *Descriptor) {
	if descriptor == nil || descriptor.Type == "" {
		return
	}
	if config.blockMap == nil {
		config.blockMap = make(map[string]*Descriptor)
	}
	config.blockMap[descriptor.Type] = descriptor
}

func (config *Config) GetBlockMapElement(descriptorType string) *Descriptor {
	return GetDescriptorFromMap(descriptorType, config.blockMap)
}

func (config *Config) SetStyleMapElement(descriptor *Descriptor) {
	if descriptor == nil || descriptor.Type == "" {
		return
	}
	if config.styleMap == nil {
		config.styleMap = make(map[string]*Descriptor)
	}
	config.styleMap[descriptor.Type] = descriptor
}

func (config *Config) GetStyleMapElement(descriptorType string) *Descriptor {
	return GetDescriptorFromMap(descriptorType, config.styleMap)
}

func SetDefaultBlocks(config *Config) {
	descriptor := new(Descriptor)
	descriptor.Type = "header-one"
	descriptor.Element = "h1"
	config.SetBlockMapElement(descriptor)

	descriptor = new(Descriptor)
	descriptor.Type = "unordered-list-item"
	descriptor.Element = "li"
	descriptor.Wrapper = "ul"
	config.SetBlockMapElement(descriptor)

	descriptor = new(Descriptor)
	descriptor.Type = "ordered-list-item"
	descriptor.Element = "li"
	descriptor.Wrapper = "ol"
	config.SetBlockMapElement(descriptor)

	descriptor = new(Descriptor)
	descriptor.Type = "unstyled"
	descriptor.Element = "p"
	config.SetBlockMapElement(descriptor)
}

func SetDefaultStyles(config *Config) {
	descriptor := new(Descriptor)
	descriptor.Type = "ITALIC"
	descriptor.Element = "em"
	config.SetStyleMapElement(descriptor)

	descriptor = new(Descriptor)
	descriptor.Type = "CODE"
	descriptor.Element = "code"
	config.SetStyleMapElement(descriptor)

	descriptor = new(Descriptor)
	descriptor.Type = "BOLD"
	descriptor.Element = "strong"
	config.SetStyleMapElement(descriptor)

	descriptor = new(Descriptor)
	descriptor.Type = "STRIKETHROUGH"
	descriptor.Element = "del"
	config.SetStyleMapElement(descriptor)

	descriptor = new(Descriptor)
	descriptor.Type = "UNDERLINE"
	descriptor.Element = "ins"
	config.SetStyleMapElement(descriptor)
}

func SetDefaultDecorators(config *Config) {
	descriptor := new(Descriptor)
	descriptor.Type = "LINK"
	descriptor.Decorator = new(LinkDecorator)
	config.SetEntityDecorator(descriptor)

}

// NewDefaultConfig Makes new config and fills it with
// default HTML elements
func NewDefaultConfig() *Config {
	config := new(Config)
	SetDefaultBlocks(config)
	SetDefaultStyles(config)
	SetDefaultDecorators(config)
	return config
}
