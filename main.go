package main

import (
	"fmt"
	"os"
)

func showMenu() {
	fmt.Printf("选择你要进行的操作,键入数字以进行选择\n")
	fmt.Printf("1.增加学员信息\n")
	fmt.Printf("2.编辑学员信息\n")
	fmt.Printf("3.展示所有学员信息\n")
	fmt.Printf("4.退出\n")
}
func input() int {

	var in int
	fmt.Scanf("%d\n", &in)
	res := in
	return res
}

func getInput() *student {
	fmt.Println("请输入学员信息")
	var (
		id    int
		name  string
		class string
	)
	fmt.Println("请输入学号")
	fmt.Scanf("%d\n", &id)
	fmt.Println("请输入名字")
	fmt.Scanf("%s\n", &name)
	fmt.Println("请输入班级")
	fmt.Scanf("%s\n", &class)
	stu := newStudent(id, name, class)
	return stu
}

func main() {
	sm := newStudents()
	for {
		showMenu()
		res := input()
		switch res {
		case 1:
			stu := getInput()
			sm.addStudent(stu)
			os.Clearenv()
		case 2:
			stu := getInput()
			sm.modifyStudent(stu)
		case 3:
			sm.showStudent()
		case 4:
			//退出
			os.Exit(0)
		}
	}

}
