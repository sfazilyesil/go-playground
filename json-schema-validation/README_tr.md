### JSON Biçimindeki Mesajın Geçerlilik Kontrolü

## Problem Tanımı: 
Gönderdiğimiz JSON formatındaki/biçimindeki bir mesajın içeriğinin belirli bir kontrata/şemaya
uygunluğunu nasıl denetleyebiliriz? Eğer mesajı alacak olan karşı tarafın denetleme
yeteneği varsa mesajı işlemeden önce alıp denetimden geçirebilir. Mesaj hatalıysa uygun bir dille
bize mesajın hatalı olduğunu söyleyebilir.  Ancak ya karşı taraf/mesajın alıcısı mesajı geçici 
(veya kalıcı) bir süreliğine tutan ve yazma esnasında JSON şema denetimi yapmayan bir veri deposu ise? Doğru mesaj/veri ilettiğimizden nasıl
emin olabiliriz? Bozuk olan yahut okuyacak olanın beklediği biçimde olmayan bir mesaj göndermekten
nasıl korunabilriz?

## Çözüm 1:
JSON dökümana/mesaja çevirmeden önce, çalışma anında JSON'a dönüştürülecek olan 
nesneleri denetlemek çözümlerden biri. ``ValidateModel(model)`` gibi bir metod 
işimizi görebilir. Metodun içinde denetimleri statik olarak ("hard-coded") 
olarak kodlayabiliriz. 

Kullandığımız dil Java ise anotasyon tabanlı olan 
"JSR-303 Bean Validation" spesifikasyonunu gerçekleştiren herhangi bir kütüphane 
deklaratif olarak kontrolleri ifade edip denetimleri kolayca yapmamızı sağlayabilir.


Go dilini kullanıyorsak, struct etiketleri (struct tags) ve go-validator kütüphanesi
benzer bir geliştirici deneyimi sunabilir bize.

## Çözüm 2:
İkinci bir çözüm ise JSON şema kullanmak. Adımlar gayet basit. 

1- Gönderilecek mesajın şemasını, yani dökümanın içindeki alanların tiplerini, 
uyacağı kuralları belirle. 

2- Çalışma anında bellekteki nesneleri JSON mesaja/dökümana çevirdikten sonra (serialization/encoding/marshalling) JSON şema denetimcisine
hem mesajı hem de uyulması gereken şemayı ver, uyumluluğu/geçerliliği denetlesin.

3- Eğer döküman/mesaj şemaya uygunsa yolla, eğer bir hata/uyumsuzluk varsa hata ver.
 
## Çözüm 1 Kod Örneği:

## Çözüm 2 Kod Örneği:


## Kısa Bir Değerlendirme:

## Sonuç:

## Kaynaklar:
- https://github.com/onsi/gomega
- http://onsi.github.io/ginkgo
- https://github.com/xeipuuv/gojsonschema
- http://json-schema.org
