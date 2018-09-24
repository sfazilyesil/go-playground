package main

import "fmt"

func main() {
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	const sample2 = "\x31\xb2\x32\x33\x34\x65\x66\x67"
	fmt.Println(sample)
	fmt.Println(sample2)

	fmt.Println("-------")

	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}

	fmt.Println("\n-------")

	fmt.Printf("% x\n", sample)
	fmt.Printf("%q\n", sample)
	fmt.Printf("%+q\n", sample)

	fmt.Println("len(sample): ", len(sample))
	fmt.Printf("\n")

	fmt.Println("#####################")

	const poi = "⌘"
	fmt.Printf("plain string: ")
	fmt.Printf("%s", poi)
	fmt.Printf("\n")

	fmt.Printf("quoted string: ")
	fmt.Printf("%+q", poi)
	fmt.Printf("\n")

	fmt.Println("len(poi): ", len(poi))

	fmt.Printf("hex bytes: ")
	for i := 0; i < len(poi); i++ {
		fmt.Printf("%x ", poi[i])
	}
	fmt.Printf("\n")
	fmt.Println("#####################")

	const poi2 = "\u2318"
	fmt.Printf("plain string: ")
	fmt.Printf("%s", poi2)
	fmt.Printf("\n")

	fmt.Printf("quoted string: ")
	fmt.Printf("%+q", poi2)
	fmt.Printf("\n")

	fmt.Println("len(poi): ", len(poi2))

	fmt.Printf("hex bytes: ")
	for i := 0; i < len(poi2); i++ {
		fmt.Printf("%x ", poi2[i])
	}
	fmt.Printf("\n")
	fmt.Println("#####################")

	bytes := []byte{255, 255, 100}
	fmt.Printf("\n%x\n", bytes)

	bytes2 := []byte("\n")
	fmt.Printf("\n%b\n", bytes2)
	fmt.Printf("\n%x\n", bytes2)
	fmt.Printf("\n%d\n", bytes2)

	fmt.Printf("%s\n", "\"string\"")
	fmt.Printf("%q\n", "\"string\"")
	fmt.Printf("%x\n", "heŞ:) thisIIÖÇöçüÜ")
	fmt.Printf("%b\n", []byte("hex this"))
	fmt.Printf("%b\n", []byte("u"))
	fmt.Printf("%s\n", "hex this")
	fmt.Printf("%q\n", "hex this`'\"")

}
