package main

import(
	"fmt"
	"flag"
	"strings"
	"strconv"
	"math"
)

type(
	NetworkAddress struct {
		IPAddress [4]int
		Address string
		Netmask int
	}

)

func (networkAddress *NetworkAddress) netMaskBinary() (output string){
	for i := 0; i < networkAddress.Netmask; i++ {
		output += "1"
	}

	for j := networkAddress.Netmask; j < 32; j++ {
		output += "0"
	}
	return
} 

func (networkAddress *NetworkAddress) netMaskDecimal(netMaskBinary string) {
	var cop string
	for i:=0; i < 4; i++ {
		v, _ := ConvertInt(netMaskBinary[i*8:i*8+8], 2, 10)
		cop += v + "."
	} 
	fmt.Println("Маска сети в десяти", cop[:len(cop) - 1])
}

func (networkAddress *NetworkAddress) hosts() (hosts float64){
	zero := 32 - networkAddress.Netmask
	hosts = math.Pow(2.0, float64(zero))
	return hosts
 
}

func (networkAddress *NetworkAddress) subNet() (subNet int) {
	subNet = int(math.Pow(2.0, float64(networkAddress.Netmask % 8)))
	return subNet
}

func parseAddress(address string) (output *NetworkAddress) {
	s := strings.Split(address, "/")
	netMask, err := strconv.Atoi(s[1])
	if err != nil {
		panic("Wrong netMask")
	}

	output = &NetworkAddress {
		Address: s[0],
		Netmask: netMask,
	}

	b := strings.Split(s[0], ".")
	for i, v := range b {
		a, _:= strconv.Atoi(v)
		output.IPAddress[i] = a 
	}

	return
}

func  (networkAddress *NetworkAddress) countOctets() (output int) {
	output = 4 - networkAddress.Netmask/8
	return
}

func (blabla *NetworkAddress) minAddress() {
	for i := 0; i < blabla.countOctets(); i++ {
		blabla.IPAddress[3 - i] = 0
	} 
}

func main() {
	inputAddress := flag.String("addr", "59.124.163.151/27", "Network address")
	flag.Parse()
	netAddr := parseAddress(*inputAddress)
	fmt.Println(netAddr.Address, netAddr.Netmask)
	netMaskBinary := netAddr.netMaskBinary()
	fmt.Println("Маска в формате двоичных чисел :", netMaskBinary)
	netAddr.netMaskDecimal(netMaskBinary)
	fmt.Println("Кол-во подсетей", netAddr.subNet())
	fmt.Println("Кол-во хостов в подсети", netAddr.hosts())
	fmt.Println("Кол-во последних изменяемых октетов:", netAddr.countOctets())

	result := 0
	for i:= 0; i < netAddr.subNet(); i++ {
		result = result + int(netAddr.hosts())
	fmt.Println(result)
	}
	netAddr.minAddress()
	fmt.Println(netAddr.IPAddress)

	/*maxAddress := int64(networkAddress.hosts()) * int64(networkAddress.subNet()) -1
	v := strconv.FormatInt(maxAddress, 2)
	fmt.Println(v)*/
}

func ConvertInt(val string, base, toBase int) (string, error) {
	i, err := strconv.ParseInt(val, base, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, toBase), nil
}

func ConvertString(val int,) (str string){
	str = strconv.FormatInt(int64(val), 10)
	return str
}
