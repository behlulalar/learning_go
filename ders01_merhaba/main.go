package main //Her Go dosyası bir pakete aittir. main özel bir pakettir, "bu bir çalıştırılabilir program" demektir.

import "fmt" //Standart kütüphaneden fmt (format) paketini getirir. Python'daki import ile aynı fikir.

func main() {

	fmt.Println("Hello go!")
	fmt.Println("Selamlarsssss!")
}

//func main() { ... }: func fonksiyon tanımlar. Program çalışınca Go her zaman main fonksiyonundan başlar. Python'daki if __name__ == "__main__" bloğuna benzer.
//fmt.Println(...): fmt paketinin Println fonksiyonunu çağırır, yani Python'daki print. Println büyük P ile başlar. Bu tesadüf değil, ilerde nedenini göreceğiz.
