package surprise

func aaa() (done func(), err error) {
	return func() {
		print("aaa: done")
	}, nil
}

func bbb() (done func(), _ error) {
	done, err := aaa()
	return func() {
		print("bbb: surprise!")
		done()
	}, err
}
