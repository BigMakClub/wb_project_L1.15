package main

//var justString string
//
//func someFunc() {
//	v := createHugeString(1 &lt;&lt; 10)
//	justString = v[:100]
//}

// Проблема реализации функции, которая закоменчена выше, в том, что когда мы делаем срез,
// то указатель внутри на массив будет указывать все равно на указаетель из слайса v,
// в таком случае у нас будет зависимость justString от указателя на v,который занимает много памяти
// а нам нужно, что было justString был свой указатель на массив байт длинной в 100 елементов,
// чтобы сборщик мусора очистил v

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = string([]rune(v)[:100])
	// Руны использую, чтобы срез не порезал буквы, если там используется кириллица и т.д
	//rune(v)[:100] - копирует символы и работает с новым массивом
	// в данной реализации v может быть удалена сборщиком мусора
}
func main() {
	someFunc()
}
