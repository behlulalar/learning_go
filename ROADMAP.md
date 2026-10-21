# Go Yol Haritası (21 Ekim → 15 Kasım 2026)

~3.5 hafta. Tempo: günde ~1 saat ideal; yoğun günlerde 30 dk yeter. Hafta sonları telafi/pekiştirme. Plan esnektir; geride kalırsak son hafta daraltılır (★ = atlanabilir).

## Hafta 1 (21–27 Ekim): Temeller
- **Ders 1:** Kurulum, `go version`, `go run`, ilk program (`package main`, `func main`, `fmt.Println`)
- **Ders 2:** Değişkenler, sabitler, temel tipler (`int`, `string`, `bool`, `float64`), `:=` ve `var`
- **Ders 3:** Operatörler, `if / else`, `switch`
- **Ders 4:** Döngüler (`for` — Go'da tek döngü), `break/continue`
- **Ders 5:** Fonksiyonlar, çoklu dönüş değeri, paketler ve `go mod init`
- **Mini görev:** Hesap makinesi + FizzBuzz

## Hafta 2 (28 Ekim–3 Kasım): Veri yapıları
- **Ders 6:** Array ve slice (`append`, `len`, `cap`, dilimleme)
- **Ders 7:** Map
- **Ders 8:** Struct ve metotlar
- **Ders 9:** Pointer (işaretçi) — sade anlatımla
- **Ders 10:** Hata yönetimi (`error`, `if err != nil`, `errors.New`, `fmt.Errorf`)
- **Mini görev:** Terminalde yapılacaklar (todo) listesi — bellekte

## Hafta 3 (4–10 Kasım): Go'ya özgü güç
- **Ders 11:** Interface'ler
- **Ders 12:** Dosya okuma/yazma, JSON (`encoding/json`)
- **Ders 13:** Test yazma (`go test`)
- **Ders 14:** Goroutine ve channel'a giriş (Go'nun ayırt edici özelliği)
- **Ders 15 ★:** `sync.WaitGroup`, `select`, `context` temelleri
- **Mini görev:** Todo uygulamasını JSON dosyasına kaydeden hale getir

## Hafta 4 (11–15 Kasım): Proje ve toparlama
- **Ders 16 ★:** `net/http` ile basit REST API
- **Final projesi:** Küçük bir CLI veya HTTP API (kullanıcı seçer; todo uygulamasının genişletilmişi önerilir)
- **Tekrar:** Zayıf kalan konular, kod okuma, sonraki öğrenme önerileri

## Çıkışta ulaşılacak seviye
Go syntax'ını rahat okuyup yazmak, struct/interface/hata yönetimiyle küçük programlar kurmak, temel test ve goroutine kullanmak, basit bir API/CLI çıkarabilmek.
