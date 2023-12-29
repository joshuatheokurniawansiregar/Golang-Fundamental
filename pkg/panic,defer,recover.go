package pkg

import (
	"fmt"
	_ "fmt"
	"os"
)

func CallsPanic() {
	_, err := os.Create("/tmp/fil")
	if err != nil {
		panic(err)
	}
}

func endApp() {
	fmt.Println("end app")
}
func endApp1() {
	fmt.Println("end app 1")
}

func logging() {
	fmt.Print("selesai memanggil function")
}

func RunApplication(x int) {
	defer logging()
	y := 10 / x
	fmt.Println("Result: ", y)
}

func CallsDefer1(error bool) {
	defer endApp()
	defer endApp1()
	if error {
		panic("boo")
	}
}

func CallsDefer(boo bool) {
	f := createFile("tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func endapp() {
	r := recover()
	fmt.Println("error dengan message", r)
	fmt.Println("aplikasi berhenti")
}

func CallsRecover() {
	defer endapp()
	isTrue := false
	if isTrue {
		panic("Aplikasi Error")
	}
	fmt.Println("Aplikasi berjalan")
}

func createFile(pt string) *os.File {
	fmt.Println("create file")
	f, err_file := os.Create(pt)
	if err_file != nil {
		panic(err_file)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("write file")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	err := f.Close()
	if err != nil {
		fmt.Fprint(os.Stderr, "error = %v\n", err)
	}
	fmt.Println("Close file")
}
