# Ubuntu container kurulumu

## Minimum SSH Akışı

```bash
docker compose up -d --build
```

Bu komut image'ı build eder ve SSH container'ı ayağa kaldırır.

Host tarafında public key üret:

```bash
ssh-keygen -t ed25519 -f ~/.ssh/usr_test -N ""
cat ~/.ssh/usr_test.pub > ssh/authorized_keys
chmod 600 ssh/authorized_keys
```

Container'a bağlan:

```bash
ssh -i ~/.ssh/usr_test -p 2222 usr@localhost
```

Container içinden host'a geri bağlanmak için host tarafında `sshd` açık olmalı ve host'a ait bir kullanıcı için public key yetkisi verilmiş olmalı. Container içinden örnek bağlantı:

```bash
ssh <host-user>@host.docker.internal
```

Bu kurulumda otomasyon script'i zorunlu değil; temel akış sadece iki SSH uç noktasıdır: host -> container ve container -> host.