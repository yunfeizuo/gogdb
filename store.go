package main

type Store struct {
}

func (*Store) Get(key string) interface{} {
	var id int64
	id = 2
	if key == "1.root" {
		return id
	}
	if key == "2.Content" {
		return "XYZ"
	}
	if key == "2.Length" {
		return 123
	}
	return nil
}
