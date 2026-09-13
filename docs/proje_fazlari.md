# NexusNode P2P Depolama Ağı - Proje Geliştirme Fazları

### FAZ 1: Çekirdek Ağ ve Parçalama (Core & Sharding)
*   Merkezi sunucu olmadan cihazların P2P iletişim kurması (WebRTC / libp2p).
*   Dosyanın matematiksel parçalara bölünmesi (Sharding) ve ağa dağıtılması.
*   Parçaların geri çağrılarak dosyanın birleştirilmesi.

### FAZ 2: Kriptografi ve Yedeklilik (Security & Redundancy)
*   AES-256 şifreleme entegrasyonu.
*   Reed-Solomon algoritması entegrasyonu (Düğüm kopmalarına karşı tolerans).
*   Self-Healing (Kendi Kendini Onarma) mekanizmasının yazılması.

### FAZ 3: Kullanıcı Arayüzü ve Masaüstü İstemcisi
*   React/Electron.js ile masaüstü uygulaması geliştirilmesi.
*   Sistem sağlığını gösteren Dashboard yapımı.

### FAZ 4: İleri Seviye Genişleme
*   Merkeziyetsiz Kimlik (DID) Entegrasyonu.
*   Token Ekonomisi (Kapasite paylaşanlara kredi verilmesi).
*   Federe Versiyon Kontrol sistemi altyapısına dönüştürme.
