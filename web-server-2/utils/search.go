package utils

type SearchableList interface {
	Len() int
	Get(i int) Searchable
}

type Searchable interface {
	IsLookup(key interface{}) bool
}

func LinearSearch(slice SearchableList, key interface{}) (int, interface{}) {
	for i := 0; i < slice.Len(); i++ {
		if slice.Get(i).IsLookup(key) {
			return i, slice.Get(i)
		}
	}

	return -1, nil
}
