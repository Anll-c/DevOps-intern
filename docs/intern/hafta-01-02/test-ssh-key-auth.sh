#!/usr/bin/env bash
set -euo pipefail

# 1. Sabit dosya yollarını tanımla (Kayıt yolunu ekrana da basıyoruz)
KEY_PATH="${HOME}/.ssh/stajyer_key"
PUBKEY_PATH="${KEY_PATH}.pub"
AUTHORIZED_KEYS_PATH="./ssh/authorized_keys"

echo "Anahtarın aranacağı/kaydedileceği hedef yol: ${KEY_PATH}"

# 2. Proje içindeki ssh klasörünü oluştur
mkdir -p ./ssh

# 3. Eğer bilgisayarında bu anahtar yoksa otomatik üret
if [[ ! -f "${PUBKEY_PATH}" ]]; then
  echo "⚡ Dosya bulunamadı! Yeni SSH anahtarı üretiliyor..."
  ssh-keygen -t ed25519 -f "${KEY_PATH}" -C "stajyer@local" -N ""
else
  echo "Bu anahtar zaten sistemde var, yeniden üretilmeksizin mevcut olan kopyalanıyor."
fiexi

# 4. Açık anahtarı doğrudan konteynerin okuyacağı dosyaya kopyala ve izinleri sıkılaştır
cat "${PUBKEY_PATH}" > "${AUTHORIZED_KEYS_PATH}"
chmod 700 ./ssh
chmod 600 "${AUTHORIZED_KEYS_PATH}"

echo "SSH anahtarı hazırlandı: ${AUTHORIZED_KEYS_PATH}"
echo "Konteynere bağlanmak için şu komutu çalıştır:"
echo "ssh -i ${KEY_PATH} -p 2222 usr@localhost"