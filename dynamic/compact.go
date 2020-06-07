package dynamic

func (name Name) Compact(args ...interface{}) (paramNames []string, paramAndValues map[string]interface{}) {
	return name.DepthCompact(1, args...)
}

func (name Name) DepthCompact(depth int, args ...interface{}) (paramNames []string, paramAndValues map[string]interface{}) {
	paramNames = name.VarNameDepth(depth+1, args...)

	// because of the variable depth
	// len(paramNames) would large than len(args)
	// so we put each args to paramNames by reversed order
	length := len(args)
	paramAndValues = make(map[string]interface{}, length)
	for i := 1; i <= length; i++ {
		paramAndValues[paramNames[len(paramNames)-i]] = args[len(args)-i]
	}

	return
}
