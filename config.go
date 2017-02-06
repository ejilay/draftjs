package draftjs

type HTMLConfig struct {
	entityDecorators map[string]*HTMLDescriptor
	blockMap         map[string]*HTMLDescriptor
	styleMap         map[string]*HTMLDescriptor
}

type HTMLDescriptor struct {
	Type      string
	Element   string
	Wrapper   string
	Decorator Decorator
}

func (config *HTMLConfig) SetEntityDecorator(descriptor *HTMLDescriptor) {
	if descriptor == nil || descriptor.Type == "" {
		return
	}
	if config.entityDecorators == nil {
		config.entityDecorators = make(map[string]*HTMLDescriptor)
	}
	config.entityDecorators[descriptor.Type] = descriptor
}

func (config *HTMLConfig) GetEntityDecorator(descriptorType string) *HTMLDescriptor {
	return GetDescriptorFromMap(descriptorType, config.entityDecorators)
}

func (config *HTMLConfig) SetBlockMapElement(descriptor *HTMLDescriptor) {
	if descriptor == nil || descriptor.Type == "" {
		return
	}
	if config.blockMap == nil {
		config.blockMap = make(map[string]*HTMLDescriptor)
	}
	config.blockMap[descriptor.Type] = descriptor
}

func (config *HTMLConfig) GetBlockMapElement(descriptorType string) *HTMLDescriptor {
	return GetDescriptorFromMap(descriptorType, config.blockMap)
}

func (config *HTMLConfig) SetStyleMapElement(descriptor *HTMLDescriptor) {
	if descriptor == nil || descriptor.Type == "" {
		return
	}
	if config.styleMap == nil {
		config.styleMap = make(map[string]*HTMLDescriptor)
	}
	config.styleMap[descriptor.Type] = descriptor
}

func (config *HTMLConfig) GetStyleMapElement(descriptorType string) *HTMLDescriptor {
	return GetDescriptorFromMap(descriptorType, config.styleMap)
}

func SetDefaultBlocks(config *HTMLConfig) {
	descriptor := new(HTMLDescriptor)
	descriptor.Type = "header-one"
	descriptor.Element = "h1"
	config.SetBlockMapElement(descriptor)

	descriptor = new(HTMLDescriptor)
	descriptor.Type = "unordered-list-item"
	descriptor.Element = "li"
	descriptor.Wrapper = "ul"
	config.SetBlockMapElement(descriptor)

	descriptor = new(HTMLDescriptor)
	descriptor.Type = "ordered-list-item"
	descriptor.Element = "li"
	descriptor.Wrapper = "ol"
	config.SetBlockMapElement(descriptor)

	descriptor = new(HTMLDescriptor)
	descriptor.Type = "unstyled"
	descriptor.Element = "p"
	config.SetBlockMapElement(descriptor)
}

func SetDefaultStyles(config *HTMLConfig) {
	descriptor := new(HTMLDescriptor)
	descriptor.Type = "ITALIC"
	descriptor.Element = "em"
	config.SetStyleMapElement(descriptor)

	descriptor = new(HTMLDescriptor)
	descriptor.Type = "CODE"
	descriptor.Element = "code"
	config.SetStyleMapElement(descriptor)

	descriptor = new(HTMLDescriptor)
	descriptor.Type = "BOLD"
	descriptor.Element = "strong"
	config.SetStyleMapElement(descriptor)

	descriptor = new(HTMLDescriptor)
	descriptor.Type = "STRIKETHROUGH"
	descriptor.Element = "del"
	config.SetStyleMapElement(descriptor)

	descriptor = new(HTMLDescriptor)
	descriptor.Type = "UNDERLINE"
	descriptor.Element = "ins"
	config.SetStyleMapElement(descriptor)
}

func SetDefaultDecorators(config *HTMLConfig) {
	descriptor := new(HTMLDescriptor)
	descriptor.Type = "LINK"
	descriptor.Decorator = new(LinkDecorator)
	config.SetEntityDecorator(descriptor)

}

func DefaultConfig() *HTMLConfig {
	config := new(HTMLConfig)
	SetDefaultBlocks(config)
	SetDefaultStyles(config)
	SetDefaultDecorators(config)
	return config
}
