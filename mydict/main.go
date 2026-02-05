package mydict

//
//func main() {
//	dictionary := mydict.Dictionary{}
//	fmt.Println(dictionary)
//	dictionary["hello"] = "world"
//	fmt.Println(dictionary)
//	diResult := dictionary["hello"]
//	fmt.Println(diResult)
//
//	dictionary = mydict.Dictionary{"first": "First word"}
//	fmt.Println(dictionary["first"])
//	search, err := dictionary.Search("first")
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(search)
//
//}
//
//func main() {
//	dictionary := mydict.Dictionary{}
//	word := "hello"
//	definition := "greeting"
//	err := dictionary.Add(word, definition)
//	if err != nil {
//		fmt.Println(err)
//	}
//	hello, _ := dictionary.Search(word)
//	fmt.Println("found", word, "definition", hello)
//	err2 := dictionary.Add(word, definition)
//	if err2 != nil {
//		fmt.Println(err2)
//	}
//	fmt.Println(definition)
//}

//func main() {
//	dictionary := mydict.Dictionary{}
//	baseWord := "hello"
//	dictionary.Add(baseWord, "First")
//	err := dictionary.Update(baseWord, "Second")
//	if err != nil {
//		fmt.Println(err)
//	}
//	word, err := dictionary.Search(baseWord)
//	fmt.Println(word)
//}
//
//func main() {
//	dictionary := mydict.Dictionary{}
//	baseWord := "hello"
//	baseDefinition := "First"
//	dictionary.Add(baseWord, baseDefinition)
//	err := dictionary.Delete(baseWord)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(dictionary)
//}
