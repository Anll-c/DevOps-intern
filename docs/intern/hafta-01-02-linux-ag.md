# Hafta 1-2 — Linux & Ağ Temelleri Teslim Formu

> **Stajyer:** _____________________  **Tarih:** __________  **Onaylayan:** __________

## Amaç
Linux terminal, dosya/izin/servis yönetimi ve temel ağ kavramlarının kazanılması.

## Ödevler

### Ödev 1 — Ubuntu container'da kullanıcı + SSH key auth
- [ ] Tamamlandı
- Kullanılan komutlar / adımlar:
```
(komutlar)
```
- Doğrulama (SSH ile key üzerinden bağlanma çıktısı):
```
(çıktı)
```

### Ödev 2 — Dosya izinleri ve sahiplikleri (chmod / chown)
- [ ] Tamamlandı
- Kavram özeti — `ls -l` çıktısındaki izin bloğunu açıkla (örn. `-rwxr-xr--`):
  - Dosya tipi:
  - Sahip (user) izinleri:
  - Grup izinleri:
  - Diğerleri (others) izinleri:
- Alıştırmalar (komut + çıktıyı yapıştır):
  - [ ] Bir dosya oluştur, `ls -l` ile başlangıç izinlerini göster
  - [ ] `chmod` ile **sembolik** yöntemle izin değiştir (örn. `chmod u+x`, `chmod go-w`)
  - [ ] `chmod` ile **sayısal (octal)** yöntemle izin değiştir (örn. `chmod 640`, `chmod 755`)
  - [ ] `chown` ile dosya sahibini değiştir (örn. `sudo chown user:group dosya`)
  - [ ] `chown -R` ile bir dizini özyinelemeli (recursive) sahiplendir
  - [ ] `umask` değerini göster ve yeni dosyanın varsayılan iznini açıkla
```bash
(komutlar)
```
```
(çıktılar)
```
- Octal izin tablosunu doldur:

| Octal | Sembolik | Anlamı |
|-------|----------|--------|
| 7 | rwx | oku + yaz + çalıştır |
| 6 | | |
| 5 | | |
| 4 | | |
| 0 | --- | izin yok |

- Mini senaryo: `deploy.sh` script'ini yalnızca sahibi çalıştırabilsin, grup okuyabilsin,
  diğerleri hiçbir şey yapamasın. Gereken tek `chmod` komutu nedir?
```bash
(cevap)
```

### Ödev 3 — .log dosyalarını sayan bash script
- [ ] Tamamlandı
- Script:
```bash
(script)
```
- Çalıştırma çıktısı:
```
(çıktı)
```

### Ödev 4 — systemd servisi + journalctl
- [ ] Tamamlandı
- Unit dosyası:
```ini
(unit içeriği)
```
- `journalctl` çıktısı:
```
(çıktı)
```

### Ödev 5 — Gerçek config okuma & açıklama
Her biri için 3-5 cümlelik özet yaz:

| Config | Ne işe yarıyor? (özet) |
|--------|------------------------|
| `ifeserver/bind9/` (DNS) | |
| `ifeserver/chrony/chrony.conf` (NTP) | |
| `ifeserver/dhcp/kea-dhcp4.conf` (DHCP) | |
| `ifeserver/netplan/` (ağ) | |
| `ifeserver/nginx/` (reverse proxy) | |

## Öğrendiğim Yeni Kavramlar
-

## Sorular / Takıldığım Noktalar
- docker-compose içersinde context tam olarak ne işe yarar? Eğer build sırasında lokasyon belirleyerek Dockerfile'ın hangi lokasyonları kullanabileceğini sınırlıyorsa birden fazla lokasyon kullanmamız gerektiğinde bunu belirtebiliyor muyuz?
- Dockerfile içerisinde "ENV DEBIAN_FRONTEND=noninteractive" kullandığımız durumda paket kurulumları sırasındaki interaktif kısımlar (y/n, timezone gibi) kapatılıyor ve default cevaplar varsayılarak paketler kuruluyor değil mi? Eğer kurmak istediğimiz bir pakette y/n olarak onay isterse ve default cevabı no ise nasıl izin verilebilir?
- "&& rm -rf /var/lib/apt/lists/*" Dockerfile içerisindeki bu komut standart ve genellikle kullanılan bir işlem midir? Sadece imaj boyutunu küçülmesi için mi tercih edilir?



## Öz Değerlendirme (1-5)
- Linux terminal: ☐1 ☐2 ☐3 ☐4 ☐5
- İzinler/sahiplik (chmod/chown): ☐1 ☐2 ☐3 ☐4 ☐5
- Ağ kavramları: ☐1 ☐2 ☐3 ☐4 ☐5
- systemd/servis: ☐1 ☐2 ☐3 ☐4 ☐5

## Demo
- [ ] Hafta sonu demo yapıldı — Tarih: __________
