package main
import ("fmt")

func main(){
	var mymap map[string]int //nill map no key gives error
	fmt.Println(mymap)
	//mymap["en"]=1 error

	// init

	cod := make(map[string]int)
	cod["ln"]=1
	fmt.Println(cod)
	codes := map[string]string{"en":"english","fr":"french"}
	fmt.Println(codes)
	fmt.Println(len(codes),len(cod),len(mymap))
	fmt.Println(codes["en"])

	//value found : mapname[key]
	value, found := codes["en"]
	fmt.Println(value,found)
	value, found = codes["esn"]
	fmt.Println(value,found)

	codes["it"]="Italian"
	fmt.Println(codes)
	codes["en"]="English ln"
	fmt.Println(codes)

	delete(codes,"en")
	fmt.Println(codes)

	//loop
	for key,value :=range codes{
		fmt.Println(key, "=>", value)
	}
}