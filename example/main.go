package main //包声明
import "unicode/utf8"

func main() {
	println(len("你好"))                    //输出6
	println(utf8.RuneCountInString("你好")) //输出2
}

func Fun1(a string, b string) (int, string) {

	_ := Fun2()
}
