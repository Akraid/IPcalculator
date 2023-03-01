package main

import(
	"fmt"
	"flag"
	"strings"
	"strconv"
	"math"
	"encoding/json"
)

type(

	SubNet struct {
		SubnetNumber int `json:"subnet_number"`
		SubnetAddress [4]int `json:"subnet2_number"`
		Broadcast [4]int `json:"subnet3_number"`
		StartAddress [4]int `json:"subnet4_number"`
		EndAddress [4]int `json:"subnet6_number"`
	}

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

func (blabla *NetworkAddress) maxAddress() {
	for i:=0; i <= 3; i++ {
		if blabla.IPAddress[i] == 0 {
			blabla.IPAddress[i] = blabla.IPAddress[i] + 255
		}
	}
}

func (blabla *NetworkAddress) hostsAddress() (subNets []SubNet) {
	blabla.minAddress()
	minad := blabla.IPAddress
	minad2 := blabla.IPAddress
	for i :=1; i <= blabla.subNet(); i++ {
		var subnet SubNet
		switch blabla.countOctets() {
		case 1:
			minad[3] += int(blabla.hosts())
			blabla.sub_Net(&minad, &minad2, i, &subnet)
		case 2:
			koren := int(blabla.hosts()) / int(blabla.subNet())
			result := math.Sqrt(float64(koren))
			minad[2] = minad[2] + int(result)
			blabla.sub_Net(&minad, &minad2, i, &subnet)
		case 3:
			koren := int(blabla.hosts()) / int(blabla.subNet())
			result := math.Sqrt(float64(koren)) / math.Sqrt(float64(256))
			minad[1] = minad[1] + int(result)
			blabla.sub_Net(&minad, &minad2, i, &subnet)
		case 4:
			koren := int(blabla.hosts()) / int(blabla.subNet())
			result := math.Sqrt(float64(koren)) / math.Sqrt(float64(256))
			result = result / math.Sqrt(float64(256))
			minad[0] = minad[0] + int(result)
			blabla.sub_Net(&minad, &minad2, i, &subnet)
		}

	subNets = append(subNets, subnet)

	}

	return
}

func (blabla *NetworkAddress) sub_Net(minad, minad2 *[4]int, i int, subnet *SubNet) {
	minad3 := *minad
	minad3[4 - blabla.countOctets()] -= 1

	for k := 1; k < blabla.countOctets() ; k++ {

		minad3[4 - k] = 255
	}

	fmt.Println("")
	fmt.Println("Подсеть номер:", i)
	subnet.SubnetNumber = i
	fmt.Println("Подсеть -", minad2)
	subnet.SubnetAddress = *minad2
	fmt.Println("Broadcast -", minad3)
	subnet.Broadcast = minad3
	minad2[3] += 1
	minad3[3] -= 1
	fmt.Println(minad2, "-", minad3)
	subnet.StartAddress = *minad2
	subnet.EndAddress = minad3
	*minad2 = *minad	
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
	netAddr.minAddress()
	fmt.Println("Минимальный адрес хоста", netAddr.IPAddress)
	netAddr.maxAddress()
	fmt.Println("Максимальный адрес хоста", netAddr.IPAddress)
	subnets := netAddr.hostsAddress()
	b, err := json.Marshal(subnets)
	if err != nil {
		panic("Json Error")
	}
	fmt.Println(string(b))
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
