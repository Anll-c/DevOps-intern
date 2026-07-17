# DevOps Stajyer Yol Haritası

## Ödev Felsefesi

Her ödev üç adımlı ilerler:

1. **Oku** — Bu repo'daki gerçek bir konfigürasyonu incele.
2. **Açıkla** — "Ne işe yarıyor?" sorusunu yazılı olarak yanıtla (README formatında).
3. **Kur** — Aynı şeyin laptop'ta küçük bir replikasını oluştur.

Bu yaklaşım teoriyi doğrudan şirketin gerçek stack'ine bağlar.

## Ön Hazırlık (Hafta 0 — 1 gün)

Stajyerin laptop'una kurulacaklar:

- [ ] Docker Desktop (veya Linux'ta Docker Engine)
- [ ] `minikube` + `kubectl`
- [ ] `git`
- [ ] VS Code + eklentiler (Docker, YAML)
- [ ] Bir terminal (Linux/macOS native, Windows'ta WSL2 önerilir)

---

## Hafta 1-2 — Linux & Ağ Temelleri (ÖNCELİKLİ)

En kritik ve genişletilmiş faz. Sıfırdan başlayan biri için sağlam temel burada atılır.

### Öğrenilecek Kavramlar

**Linux**
- Dosya sistemi ve gezinme: `cd`, `ls`, `pwd`, `cp`, `mv`, `rm`, `mkdir`, `cat`, `less`
- İzinler: `chmod`, `chown`, kullanıcı/grup, `sudo`
- Paket yönetimi: `apt` (kurulum, güncelleme, arama)
- Süreçler: `ps`, `top`, `kill`, `&`, `jobs`
- Servisler: `systemd`, `systemctl`, `journalctl`
- Bash: değişkenler, `if`, `for`, basit script yazma, pipe `|`, redirect `>`
- SSH: bağlantı, `ssh-keygen`, public/private key auth

**Ağ**
- IP adresi, subnet, gateway kavramları
- Araçlar: `ping`, `curl`, `ss`, `ip`, `dig`/`nslookup`
- DNS nedir, DHCP nedir, NTP nedir
- Port ve firewall mantığı (`ufw`)
- Reverse proxy nedir (nginx örneği)

### Ödevler

1. Docker'da bir Ubuntu container başlat, içinde yeni bir kullanıcı oluştur ve SSH key auth kur.
2. **Dosya izinleri ve sahiplikleri:** `chmod` (sembolik + octal) ve `chown`/`chown -R` ile
   izin/sahiplik değiştir; `ls -l` çıktısındaki izin bloğunu (`-rwxr-xr--`) ve `umask` mantığını
   açıkla. Mini senaryo: bir `deploy.sh`'ı yalnızca sahibi çalıştırabilsin, grup okuyabilsin,
   diğerleri erişemesin.
3. Bir dizindeki `.log` dosyalarını sayan basit bir bash script yaz.
4. (VM veya WSL'de) basit bir `systemd` servisi tanımla ve `journalctl` ile loglarını izle.
5. **Oku & Açıkla:** Aşağıdaki gerçek config'leri incele ve her biri için 3-5 cümlelik özet yaz:
   - `ifeserver/bind9/` — DNS sunucusu
   - `ifeserver/chrony/chrony.conf` — zaman senkronizasyonu
   - `ifeserver/dhcp/kea-dhcp4.conf` — IP dağıtımı
   - `ifeserver/netplan/` — ağ arayüz yapılandırması
   - `ifeserver/nginx/` — reverse proxy

### Kaynaklar
- "Linux Journey" (linuxjourney.com), OverTheWire Bandit (terminal pratiği)

---

## Hafta 3 — Git & Docker Temeli

### Öğrenilecek Kavramlar

**Git**
- `clone`, `add`, `commit`, `push`, `pull`
- Branch, merge, Pull Request (PR) akışı
- `.gitignore`, çakışma (conflict) çözme

**Docker**
- Image vs container farkı
- `Dockerfile`: `FROM`, `RUN`, `COPY`, `CMD`, `ENTRYPOINT`
- Layer ve cache mantığı
- Volume, port mapping, network temelleri

### Ödevler

1. Örnek bir repo'da feature branch aç → değişiklik yap → PR oluştur.
2. Basit bir uygulamayı (örn. küçük bir Python/Node web servisi) `Dockerfile` ile paketle,
   build et ve çalıştır.
3. **Oku & Açıkla:** `nats/docker-compose.yml` ve `sonarqube/docker-compose.yml` dosyalarını
   incele, her servisin ne yaptığını yaz.

---

## Hafta 4 — Docker Compose & Ansible Girişi

### Öğrenilecek Kavramlar

**Docker Compose**
- `docker-compose.yml` yapısı: services, volumes, networks, depends_on
- Multi-container uygulama çalıştırma

**Ansible**
- Inventory, playbook, task, role kavramları
- Idempotency (aynı komutu tekrar çalıştırınca sonuç değişmez)
- `group_vars` mantığı

### Ödevler

1. 2-3 servisli kendi compose stack'ini kur (örn. bir web app + PostgreSQL).
2. `localhost`'a bir paket kuran (örn. nginx) küçük bir Ansible playbook yaz.
3. **Oku & Açıkla:** `harbor-iac/` yapısını incele (`roles/common`, `roles/docker`,
   `roles/harbor`) ve akışı bir diyagram/şema ile anlat.

---

## Hafta 5 — CI/CD (Jenkins)

### Öğrenilecek Kavramlar
- CI/CD nedir, neden gerekli
- Pipeline, stage, step, agent kavramları
- Declarative `Jenkinsfile` yapısı
- Credentials yönetimi

### Ödevler

1. Jenkins'i Docker'da çalıştır (`jenkins/jenkins:lts` image).
2. Basit bir "hello world" pipeline yaz (checkout → echo → basit bir komut).
3. **Oku & Açıkla:** `jenkins/` altındaki bir Jenkinsfile'ı (örn. `crew-control-panel`)
   satır satır açıkla.
4. Standart pipeline akışını şemalandır: `checkout → scan → build → deploy`.

---

## Hafta 6 — Kubernetes (kind/minikube)

### Öğrenilecek Kavramlar
- Neden Kubernetes: orkestrasyon, ölçeklenme, self-healing
- Temel nesneler: Pod, Deployment, Service, Ingress
- ConfigMap, Secret, Namespace
- PV / PVC (kalıcı depolama) temeli
- `kubectl` temel komutları

### Ödevler

1. kind veya minikube ile yerel cluster kur.
2. Bir uygulamayı Deployment + Service + Ingress ile deploy et.
3. `kubectl get/describe/logs/exec` komutlarını pratik et.
4. **Oku & Açıkla:** `k8s/CLUSTER.md` ve `k8s/jenkins-k8s/` manifestlerini incele,
   cluster mimarisini ve deploy akışını özetle.

---

## Hafta 7 — Monitoring + Güvenlik Temeli

### Öğrenilecek Kavramlar

**Monitoring / Observability**
- Metrics vs logs farkı
- Prometheus scrape mantığı, temel PromQL
- Grafana dashboard
- Loki + Promtail ile log toplama

**Güvenlik (DevSecOps girişi)**
- Image güvenlik taraması (Trivy)
- Secret sızıntısı taraması (Gitleaks)

### Ödevler

1. `ifeserver/monitor_v2` stack'ini laptop'ta ayağa kaldır.
2. Basit bir Grafana dashboard + 1 adet alert kuralı oluştur.
3. Bir Docker image'ı Trivy ile, bir repo'yu Gitleaks ile tara ve sonuçları raporla.

---

## Hafta 8 — Bitirme Projesi (Capstone)

Tüm öğrenilenleri birleştiren, uçtan uca ve tamamı laptop'ta yapılan proje.

### Hedef Akış

1. Küçük bir uygulamayı `Dockerfile` ile paketle.
2. Jenkins pipeline yaz: güvenlik taraması (Trivy/Gitleaks) + build.
3. Image'ı yerel bir registry'ye push et.
4. kind/minikube cluster'a deploy et (Deployment + Service + Ingress).
5. Prometheus + Grafana ile izle.
6. Tüm süreci dokümante et (README) ve ekibe **sun**.

### Değerlendirme Kriterleri
- [ ] Uygulama container olarak çalışıyor
- [ ] Pipeline başarıyla build + scan yapıyor
- [ ] K8s'te erişilebilir şekilde deploy edilmiş
- [ ] En az 1 dashboard ile izleniyor
- [ ] Anlaşılır bir README ile dokümante edilmiş

---

## Bonus / İleri Konular (süre kalırsa veya staj uzarsa)

Bunlar 1-2 aylık sıfırdan programa sığmaz; ileri seviye için ayrılmıştır:

- **Terraform (IaC):** `drm-cloud/terraform` inceleme + küçük bir local modül yazma
- **ArgoCD / GitOps:** `k8s/argocd`
- **Harbor registry & Cosign** ile image imzalama
- **Semgrep** ile SAST
- **HA Kubernetes** cluster mimarisi: `k8s/CLUSTER.md` derin okuma (HAProxy, Keepalived, Cilium, etcd)

---

## İzleme ve Değerlendirme

- **Her hafta sonu:** kısa demo + README formatında yazılı özet
- **Haftalık 1:1** review görüşmesi
- Ödevler her zaman "gerçek config oku → açıkla → laptop'ta küçük replika kur" mantığında
