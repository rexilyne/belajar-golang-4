package main

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"
)

type Motor interface {
	Berjalan()
	PengisianDaya()
	Maintenance()
}

type Tesla struct {
	batteryType         string
	batteryLevel        float64
	maintenanceLocation string
}

// bagaimana caranya agar tesla
// dapat tergolongkan sebagai implementator dari motor?

// caranya adalah dengan implementasi
// function yang dimiliki motor, oleh pointer dari tesla
func (t *Tesla) Berjalan() {
	// tesla akan bisa berjalan jika battery level > 30%
	if t.batteryLevel > 30.0 {
		fmt.Println("bisa berjalan")
		return
	}
	fmt.Println("harus mengisi daya")
}

func (t *Tesla) PengisianDaya() {

}

func (t *Tesla) Maintenance() {

}

type Honda struct {
	maintenanceLocation string
	fuelType            string
	fuelLevel           float64
}

func (h *Honda) Berjalan() {
	// honda akan bisa berjalan jika battery level > 10%
	if h.fuelLevel > 10.0 {
		fmt.Println("bisa berjalan")
		return
	}
	fmt.Println("harus mengisi daya")
}

func (h *Honda) PengisianDaya() {

}

func (h *Honda) Maintenance() {

}

func main() {
	// defer statement
	// statement yang membantu kita
	// untuk ensure function akan dijalankan
	// di akhir process
	// meskipun deklarasinya di awal process
	defer fmt.Println("process done (defer)")

	// #4 Concurrency
	// cara implementasinya -> go
	go func() {
		// ini akan berjalan di belakang layar
		// di belakang process main function
		fmt.Println("inside go routine")
	}()
	time.Sleep(time.Second * 1)

	go AsyncFunction()
	time.Sleep(time.Second * 1)
	// main function tidak tau apakah
	// async function sudah dijalankan dan selesai
	// atau belum

	// apakahj main function bisa tau
	// function ini sudah selesai atau belum
	// bisa:
	// - wait group: akan menunggu hingga process selesai
	// - channel: akan menunggu hingga mendapatkan value dari goroutine

	// goroutine:
	// process yang terjadi di dalam process besar lainnya
	// pocess besar ini (Thread)
	// go routine (Thread yang ringan dan kecil)
	// Thread dan Go Routine sama2 tempat menjalankan process

	// concurenncy : 1 thread n goroutine, bisa membaca variabel dari goroutine lainnya
	// paralel : multithreading, tidak bisa membaca variabel dari thread lainnya

	// usersName := make(map[int]string, 0)

	// go func() {
	// 	for key, val := range usersName {
	// 		fmt.Println(key, val)
	// 	}
	// }()

	// for i := 0; i < 10; i++ {
	// 	usersName[i] = fmt.Sprintf("user %v", i)
	// }

	// dalam function di bawah ini
	// main function hanya melempar value ke go routine
	// go rotine mana yang execute duluan
	// main function tidak peduli dengan itu
	for i := 0; i < 20; i++ {
		go PrintName(fmt.Sprintf("name%v", i))
	}
	time.Sleep(time.Second * 1)

	// #5 Wait Group -> termasuk ke dalam sync package
	// variabel yang memaksa main function
	// baru boleh menjalankan process selanjutnya
	// setelah semua go routine selesai dijalankan
	var wg sync.WaitGroup
	// 1. setiap go routine yang ingin di sync
	// wg harus di add1 di dalam variabelnya
	// 2. setiap selesai menjalanakan go routine
	// wg harus memanggil done
	// 3. wg harus memanggil wait untuk menunggu semua
	// go routine selesai
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(input int) {
			fmt.Println("user", input)
			wg.Done()
		}(i)
	}
	wg.Wait() // wait untul all go routine done

	var wg2 sync.WaitGroup
	WaitGroupFunc(&wg2)
	wg.Wait()

	WaitGroupFuncInside()
	fmt.Println("process done")

	// common problem in concurrency
	fmt.Println(DataRace())

	// Datarace:
	// memory hanya boleh diakses (read / write)
	// sekali dalam satu waktu
	// Data Race terjadi ketika ->
	// memory diakses oleh program / go routine yang berbeda
	// dan melakukan operasi read / write
}

func DataRace() string {
	t := "Hi"
	go func() {
		t = "Hello"
	}()
	return t
}

func WaitGroupFunc(wg *sync.WaitGroup) {
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(input int) {
			defer wg.Done()
			fmt.Println("user", input)
		}(i)
	}
}

func WaitGroupFuncInside() {
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(input int) {
			defer wg.Done()
			fmt.Println("user", input)
		}(i)
	}
	wg.Wait()
}

func PrintName(s string) {
	fmt.Println(s)
}

func AsyncFunction() {
	// ini akan berjalan di belakang layar
	// di belakang process main function
	fmt.Println("inside go routine (AsyncFucntion)")
}

func InterfaceImpl() {
	// #1 Interface
	// -> abstrak yang merepresentasikan suatu kumpulan methods
	// yang dapat diimplementasikan dengan menggunakan struct
	var motor1 Motor

	// jika dipanggil di sini akan error karena masih nil
	// motor1.Berjalan()

	motor1 = &Tesla{
		maintenanceLocation: "US",
		batteryType:         "Lithium",
		batteryLevel:        100.0,
	}
	fmt.Println(motor1)

	motor1.Berjalan()

	var motor2 Motor
	motor2 = &Honda{
		maintenanceLocation: "ID",
		fuelType:            "Pertamax",
		fuelLevel:           5.0,
	}
	motor2.Berjalan()

	fmt.Println(motor1 == motor2)

	var motor3 Motor
	motor3, _ = motor3.(*Honda)
	fmt.Println("motor 3 type: ", reflect.TypeOf(motor3))

	motor3 = &Tesla{}
	fmt.Println("motor 3 type: ", reflect.TypeOf(motor3))

	motor3 = &Honda{}
	fmt.Println("motor 3 type: ", reflect.TypeOf(motor3))

	// #2 Empty Interface
	// variable bebas yang dapat di assign ke siapapun
	// interface yang tidak memiliki syarat variable apa yang bisa diassign ke interface itu
	var variable1 interface{}
	variable1 = 1 // int
	fmt.Println(variable1)
	variable1 = 10.0 // float 64
	fmt.Println(variable1)
	variable1 = "Hello World" // string
	fmt.Println(variable1)
	variable1 = Honda{} // struct honda
	fmt.Println(variable1)

	// empty interface:
	// tidak memiliki method
	// sehingga type apapun bisa di assign ke interface tsb

	// abstraction interface:
	// memiliki method/function
	// yang menjadi syarat minimum agar
	// suatu type bisa diassign ke interface tsb

	// bukan suatu best practice ketika
	// kita menggunakan variable type menjadi
	// empty interface semua, karena GO bukan diciptakan untuk itu

	// asertion / type casting
	// any adalah alias untuk interface
	var intf1 any
	intf1 = 10
	num := intf1.(int)
	fmt.Println(num)

	// var intf2 any
	// intf2 = 10
	// // gagal typecasting karena mencoba mengubah int menjadi float
	// num2 := intf2.(float64)
	// fmt.Println(num2)
}

func ReflectionImpl() {
	// #3 Reflection
	// package built in golang
	// yang digunakan untuk manipulasi
	// data type di golang

	// menentukan type data
	int1 := 10
	var int2 any
	int2 = 11
	fmt.Println(reflect.TypeOf(int1), reflect.TypeOf(int2))
	if reflect.TypeOf(int2).Kind() != reflect.Int {
		fmt.Println("bukan int nih bos")
	}

	// deep equal
	motor1 := Honda{
		maintenanceLocation: "ID",
		fuelType:            "Pertamax",
		fuelLevel:           100.0,
	}

	motor2 := Honda{
		maintenanceLocation: "ID",
		fuelType:            "Pertamax",
		fuelLevel:           100.0,
	}

	fmt.Println(motor1 == motor2, reflect.DeepEqual(motor1, motor2))

	// ?type=sekolah
	// ?type=Sekolah
	str1 := "sekolah"
	str2 := "Sekolah"
	fmt.Println(str1 == str2,
		reflect.DeepEqual(str1, str2),
		strings.EqualFold(str1, str2),
	)

	// deep equal -> function return boolean
	// bolean -> data type

}
