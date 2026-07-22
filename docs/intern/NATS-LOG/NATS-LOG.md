# Stajyer Görevi: NATS ile Gerçek Zamanlı Log Boyutu İzleme Uygulaması

## Amaç

Docker Compose kullanılarak çalışan, *NATS tabanlı gerçek zamanlı veri akışı* oluşturan örnek bir uygulama geliştirilmesi.

## Gereksinimler

### 1. Docker Compose Ortamı

* Docker Compose dosyası oluşturulacak.
* Aşağıdaki servisler compose içerisinde yer alacak:

  * NATS Server
  * Go API Server
  * Frontend uygulaması
* Servisler aynı Docker ağı üzerinde çalışacak.
* Gerekli volume ve environment tanımları yapılacak.

### 2. NATS Server

* Custom configuration file kullanılacak.
* TLS kullanılmayacak.
* Authentication (username/password) aktif olacak.
* Gerekli subject isimleri anlamlı şekilde belirlenecek.

### 3. Go API Server

* Go dili ile geliştirilecek.
* Belirli aralıklarla (örneğin her 1 saniyede bir) belirlenen log dosyalarının boyutlarını okuyacak.
* Okunan bilgiler JSON formatında NATS üzerinden publish edilecek.
* NATS bağlantısı koparsa otomatik reconnect mekanizması bulunacak.
* Hata durumlarında uygulama kapanmak yerine log üretmeye devam edecek.

### 4. Frontend

* NATS'tan gelen verileri gerçek zamanlı olarak okuyacak.
* Log dosyalarının isimleri ve boyutlarını ekranda gösterecek.
* Arayüz sürekli büyümeyecek veya bellek tüketimi artmayacak şekilde tasarlanacak.

## Mimari Beklentiler

Arayüzde zamanla veri birikmesi nedeniyle performans problemi yaşanmaması için uygun önlemler alınmalıdır. Örneğin;

* Sürekli yeni satır eklemek yerine mevcut satırlar güncellenebilir.
* Aynı log dosyası için yalnızca son durum tutulabilir.
* Gereksiz DOM güncellemelerinden kaçınılmalıdır.
* Bellekte sınırsız veri tutulmamalıdır.
* Eğer geçmiş veri gösterilecekse maksimum kayıt sayısı sınırlandırılmalıdır (örneğin son 100 kayıt).

## Bonus Çalışmalar

* Docker healthcheck tanımlanması.
* Graceful shutdown desteği.
* Yapılandırmanın environment variable ile yönetilmesi.
* Frontend ve backend için README hazırlanması.
* Log boyutu belirli bir eşiği geçtiğinde farklı renkte gösterilmesi.
* Timestamp bilgisinin de yayınlanması.
* Docker image'larının mümkün olduğunca küçük tutulması (multi-stage build kullanılması).
* NATS subject yapısının ileride farklı metrikleri destekleyecek şekilde tasarlanması.
* Basit unit testler eklenmesi.

## Teslim Kriterleri

* Docker Compose ile tek komutla ayağa kalkmalıdır.
* Tüm servisler container içerisinde çalışmalıdır.
* README dosyasında kurulum ve çalıştırma adımları yer almalıdır.
* Kod okunabilir, modüler ve anlaşılır olmalıdır.
* Gereksiz bağımlılıklardan kaçınılmalıdır.
* Hatalı durumlar (NATS bağlantısı kesilmesi vb.) uygun şekilde ele alınmalıdır.
* Proje Git üzerinden düzenli commit geçmişi ile teslim edilmelidir.