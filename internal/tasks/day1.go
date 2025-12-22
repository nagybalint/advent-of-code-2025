package tasks

type Day1Task1 struct { }

func (*Day1Task1) GetName() string {
	return "1/1"
}

func (*Day1Task1) Run() int32 {
	return 1
}
