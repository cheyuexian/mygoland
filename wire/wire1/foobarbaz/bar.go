package foobarbaz

type Bar struct {
	X int
}

func ProviderBar(foo Foo) Bar{
	return Bar{X:-foo.X}
}

