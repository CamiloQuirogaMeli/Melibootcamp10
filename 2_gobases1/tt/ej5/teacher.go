package ej5

type Students struct{
	NamesList []string
}

func (s *Students) AddStudent(newStudent string){
	s.NamesList = append(s.NamesList, newStudent)
}

func (s *Students) GetList() []string{
	return s.NamesList
}

func (s *Students) SetList(list []string) {
	s.NamesList = list
}