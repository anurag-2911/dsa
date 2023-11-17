package recursion

func A()string{
	return " hi "+ B()
}
func B()string{
return " hello " + C()
}
func C()string{
return " halo"
}
