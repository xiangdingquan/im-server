package model

func WrapperGoFunc(f interface{}, goF func()) interface{} {
	if goF != nil {
		go goF()
	}
	return f
}
