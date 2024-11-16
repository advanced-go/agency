package common

type Activity interface {
	Add(agent string, content any)
}

var activity []any

type testActivity struct{}

func (t testActivity) Add(agent string, content any) {
	activity = append(activity, content)
}

func runActivity[T Activity](s string) {
	var t T

	t.Add(s, nil)
	//fmt.Printf("test:\n",t.Add(s,nil))
}
