// Bu dosya NATS websocket bağlantısını, DOM referanslarını ve
// start/stop butonlarının backend çağrılarını içerir.
// index.html sadece yapıyı (markup) barındırır, tüm mantık burada.

import { connect, StringCodec } from "https://cdn.jsdelivr.net/npm/nats.ws@1.30.3/esm/nats.js";

// nats.conf'daki websocket portu ile docker-compose.yml'daki
// NATS_WEBSOCKET_PORT (varsayılan 8880) burada da aynı olmalı.
const NATS_WS_PORT = 8880;
const NATS_USER = "usr";
const NATS_PASS = "usr1935";

// --- DOM referansları ---
const statusEl = document.getElementById("status");
const tbody = document.getElementById("log-table-body");
const summaryEl = document.getElementById("summary");
const startBtn = document.getElementById("start-btn");
const stopBtn = document.getElementById("stop-btn");
const actionMsgEl = document.getElementById("action-message");

// Dosya adı -> <tr> elementi eşlemesi.
// Aynı dosya için yeni satır EKLENMEZ, mevcut satır güncellenir.
const rows = new Map();

function formatBytes(bytes) {
  if (bytes >= 1024 * 1024) return (bytes / 1024 / 1024).toFixed(2) + " MB";
  if (bytes >= 1024) return (bytes / 1024).toFixed(2) + " KB";
  return bytes + " B";
}

function nowTime() {
  return new Date().toLocaleTimeString("tr-TR");
}

function updateFileRow(data) {
  let row = rows.get(data.filename);
  if (!row) {
    row = document.createElement("tr");
    row.innerHTML = "<td></td><td></td><td></td>";
    tbody.appendChild(row);
    rows.set(data.filename, row);
  }
  row.children[0].textContent = data.filename;
  row.children[1].textContent = formatBytes(data.size);
  row.children[2].textContent = nowTime();
}

function updateSummary(data) {
  summaryEl.textContent =
    `Toplam dosya: ${data.total_count} | Toplam boyut: ${formatBytes(data.total_size)} | Son güncelleme: ${nowTime()}`;
}

function setStatus(text, className) {
  statusEl.textContent = text;
  statusEl.className = className;
}

// --- Start / Stop butonları ---
// nginx.conf içindeki /api/ reverse proxy sayesinde tarayıcı
// go-api'ye değil kendi origin'ine (nginx) istek atıyor, CORS sorunu olmuyor.
async function callBackend(path) {
  actionMsgEl.textContent = "İşleniyor...";
  try {
    const res = await fetch(path, { method: "POST" });
    const data = await res.json();
    actionMsgEl.textContent = data.message || JSON.stringify(data);
  } catch (err) {
    actionMsgEl.textContent = "Hata: " + err.message;
  }
}

startBtn.addEventListener("click", () => callBackend("/api/start"));
stopBtn.addEventListener("click", () => callBackend("/api/stop"));

// --- NATS bağlantısı ---
const sc = new StringCodec();

async function main() {
  let nc;
  try {
    nc = await connect({
      servers: [`ws://${window.location.hostname}:${NATS_WS_PORT}`],
      user: NATS_USER,
      pass: NATS_PASS,
      reconnect: true,
      maxReconnectAttempts: -1,
    });
  } catch (err) {
    setStatus("Bağlantı kurulamadı: " + err.message, "disconnected");
    return;
  }

  setStatus("Bağlandı: " + nc.getServer(), "connected");

  // Bağlantı kopma / yeniden bağlanma durumunu izle
  (async () => {
    for await (const s of nc.status()) {
      if (s.type === "disconnect") {
        setStatus("Bağlantı koptu, yeniden deneniyor...", "disconnected");
      } else if (s.type === "reconnect") {
        setStatus("Yeniden bağlandı: " + nc.getServer(), "connected");
      }
    }
  })();

  const sub = nc.subscribe("logs.size.>");
  for await (const msg of sub) {
    const data = JSON.parse(sc.decode(msg.data));
    if (msg.subject.endsWith(".summary")) {
      updateSummary(data);
    } else {
      updateFileRow(data);
    }
  }
}

main();