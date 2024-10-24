package main

type NameStore map[string]int

func (ns *NameStore) GetName(name string) int {
	num, ok := (*ns)[name]
	if ok {
		return num
	}
	return 0
}

func (ns *NameStore) PutName(name string) int {
	num := (*ns)[name]
	(*ns)[name] = num + 1
	return (*ns)[name]
}

func (ns *NameStore) DeleteName(name string) int {
	num := (*ns)[name]
	if num > 1 {
		(*ns)[name] = num - 1
	} else if num == 1 {
		delete((*ns), name)
	}

	return (*ns)[name]
}