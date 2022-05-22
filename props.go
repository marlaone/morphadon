package engine

type Properties map[string]interface{}

func PropStr(propName string, props Properties) string {
	return PropStrWithDefault(propName, props, "")
}

func PropStrWithDefault(propName string, props Properties, defaultValue string) string {
	str, ok := props[propName].(string)
	if !ok {
		return defaultValue
	}
	return str
}

func PropBool(propName string, props Properties) bool {
	return PropBoolWithDefault(propName, props, false)
}

func PropBoolWithDefault(propName string, props Properties, defaultValue bool) bool {
	b, ok := props[propName].(bool)
	if !ok {
		return defaultValue
	}
	return b
}
