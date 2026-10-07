# Uçtan Uca DevOps CI/CD ve Kubernetes Mimarisi Projesi

Bu depo, modern DevOps kültürünü, CI/CD (Sürekli Entegrasyon ve Sürekli Dağıtım) pratiklerini ve altyapı otomasyon süreçlerini uçtan uca uygulamak amacıyla yürütülen 8 haftalık kapsamlı bir mühendislik bitirme projesinin (Capstone) çıktılarını içermektedir. Proje kapsamında, NATS ve Go tabanlı bir mikroservis mimarisinin tüm yazılım yaşam döngüsü tasarlanmış ve hayata geçirilmiştir.

## Projenin Amacı

Geliştirilen uygulamanın (NATS-LOG) kaynak koddan alınıp; güvenli bir şekilde derlenmesi, test edilmesi, konteynerize edilmesi ve Kubernetes ortamında kesintisiz (Zero-Downtime) olarak canlıya alınması hedeflenmiştir. 

Süreç boyunca endüstri standardı kabul edilen "İncele, Açıkla ve Uygula" metodolojisi benimsenmiş; teorik yapılandırmalar detaylıca analiz edildikten sonra, yerel ortamda (localhost/minikube) çalışan replikaları inşa edilmiştir.

---

## Mimari Şema ve İş Akışı

Aşağıdaki şema, bu projede kurulan CI/CD dağıtım hattının (pipeline) ve altyapının mimari işleyişini göstermektedir:

```mermaid
graph TD
    A[Geliştirici] -->|Git Push| B(GitHub Repository)
    B -->|Webhook / Poll| C{Jenkins CI/CD}
    
    subgraph Sürekli Entegrasyon & Güvenlik
    C -->|1. Checkout & Test| D[Go Vet & Birim Testler]
    D -->|2. Güvenlik| E[Gitleaks Secret Taraması]
    E -->|3. Build| F[Docker Buildx - Multi-Stage]
    F -->|4. Güvenlik| G[Trivy Image Taraması]
    end

    G -->|5. Push| H[(Yerel / Docker Registry)]

    subgraph Sürekli Dağıtım & Orkestrasyon
    H -->|6. Deploy| I[Kubernetes Cluster - Minikube]
    I --> J[K8s Deployment & Pods]
    I --> K[K8s Service & Ingress]
    end

    subgraph İzleme ve Loglama
    J -.->|Metrics| L((Prometheus))
    L -.-> M[Grafana Dashboard]
    J -.->|Logs| N((Loki & Promtail))
    end
