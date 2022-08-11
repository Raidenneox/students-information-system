package main

import "fmt"

type student struct {
	id    int
	name  string
	class string
}

//student 的构造函数
func newStudent(id int, name, class string) *student {
	return &student{
		id:    id,
		name:  name,
		class: class,
	}
}

type students struct {
	allStudents []*student
}

func newStudents() *students {
	return &students{
		allStudents: make([]*student, 0, 100),
	}
}

func (s *students) addStudent(newStu *student) {
	s.allStudents = append(s.allStudents, newStu)
}

func (s *students) modifyStudent(newStu *student) {
	for k, v := range s.allStudents {
		if newStu.id == v.id { //当学号相同时就是找到了要修改的学生
			s.allStudents[k] = newStu
			return
		}
	}
	fmt.Printf("没有学号是%d的信息\n", newStu.id)
}

func (s *students) showStudent() {
	for _, v := range s.allStudents {
		fmt.Printf("学号:%d\n", v.id)
		fmt.Printf("姓名:%s\n", v.name)
		fmt.Printf("班级:%s\n", v.class)
	}
}
