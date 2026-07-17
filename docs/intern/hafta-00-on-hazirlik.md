# Hafta 0 — Ön Hazırlık Teslim Formu

> **Stajyer:** _____________________  **Tarih:** __________  **Onaylayan:** __________

## Amaç
Laptop'ta pratik ortamının kurulması ve doğrulanması.

## Kurulum Checklist

| Araç | Kuruldu mu? | Versiyon | Notlar |
|------|:-----------:|----------|--------|
| Docker Desktop / Engine | ☑ | 29.1.3 | |
| minikube | ☑ | | 1.38.1 |
| kubectl | ☑ | | Client/Kustomize version? |
| git | ☑ | 2.53.0 | |
| VS Code + eklentiler | ☑ | 1.129.0 | |
| Terminal (WSL2 / native) | ☑ | 5.3.9 | |

## Doğrulama Komutları (çıktıları yapıştır)

```bash
docker --version
docker run hello-world
kubectl version --client
git --version
```

<!-- Çıktıları buraya yapıştır -->
```
Docker version 29.1.3, build 29.1.3-0ubuntu4.1

Hello from Docker!
This message shows that your installation appears to be working correctly.

To generate this message, Docker took the following steps:
 1. The Docker client contacted the Docker daemon.
 2. The Docker daemon pulled the "hello-world" image from the Docker Hub.
    (arm64v8)
 3. The Docker daemon created a new container from that image which runs the
    executable that produces the output you are currently reading.
 4. The Docker daemon streamed that output to the Docker client, which sent it
    to your terminal.

To try something more ambitious, you can run an Ubuntu container with:
 $ docker run -it ubuntu bash

Share images, automate workflows, and more with a free Docker ID:
 https://hub.docker.com/

For more examples and ideas, visit:
 https://docs.docker.com/get-started/

Client Version: v1.36.2
Kustomize Version: v5.8.1
git version 2.53.0
```

## Karşılaşılan Sorunlar ve Çözümler
-

## Sorular / Takıldığım Noktalar
- kubectl ve minikube ne için kullanılır?
- git desktop gerekli mi? ubuntu için git desktop'un resmi bir app'i yok sanırım.

## Öz Değerlendirme (1-5)
- Ortam kurulumuna hakimiyet: ☐1 ☐2 ☐3 ☑4 ☐5
