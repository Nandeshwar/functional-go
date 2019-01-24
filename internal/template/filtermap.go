package template

func FilterMap() string {
	return `
func FilterMap<CONDITIONAL_TYPE>(fFilter func(<TYPE>) bool, fMap func(<TYPE>) <TYPE>, list []<TYPE>) []<TYPE> {
	if fFilter == nil || fMap == nil {
		return []<TYPE>{}
	}
	var newList []<TYPE>
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}`
}
