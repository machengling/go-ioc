package common

type TestConponent struct {
	Name string
}

// ConponentName 自定义组件名称
func (this *TestConponent) ConponentName() string {
	return "TestConponent1"
}
