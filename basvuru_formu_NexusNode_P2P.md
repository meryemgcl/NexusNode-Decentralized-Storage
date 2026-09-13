# KASTAMONU ÜNİVERSİTESİ 1. AR-GE PROJE PAZARI BAŞVURU FORMU

**Proje Başlığı:** NexusNode: Kurumsal Veri Egemenliği İçin Eşler Arası (P2P) Merkeziyetsiz Depolama Ağı
**Başvuru Alanı:** Fen Bilimleri / Yazılım ve Bilişim Teknolojileri

### Proje Özeti ve Anahtar Kelimeler
Günümüzde kurumlar ve bireyler, kritik verilerini merkezi bulut sağlayıcılarında barındırmakta olup; bu durum veri egemenliği ihlallerine ve tek nokta hatalarına neden olmaktadır. Bu projenin amacı, internete bağlı sıradan cihazların birer sunucu gibi çalışmasını sağlayan Eşler Arası (P2P) merkeziyetsiz bir dosya depolama ve paylaşım ağı geliştirmektir. Sistem, yüklenen dosyaları "Silinti Kodlaması (Erasure Coding)" algoritmaları ile matematiksel parçalara böler, şifreler ve ağdaki uç düğümlere (node) dağıtır. Donanım ve yapay zeka kullanılmayan bu saf yazılım projesi, merkezi bir sunucu çöktüğünde bile verinin %100 kayıpsız geri çağrılabilmesini sağlar.
**Anahtar Kelimeler:** Merkeziyetsiz Sistemler, P2P Ağlar, Veri Egemenliği, Silinti Kodlaması, Siber Güvenlik.

### Projenin Amaç ve Hedefleri
*   **Amaç:** Kurumların kapalı veya açık, veri kaybı riski sıfır olan merkeziyetsiz bir depolama ağı kurmalarını sağlayacak yazılım protokolünü geliştirmek.
*   **Hedefler:**
    1. İstemciler arası doğrudan iletişimi sağlayacak WebRTC/libp2p tabanlı çekirdek ağ mimarisini kurmak.
    2. AES-256 ve Reed-Solomon hata düzeltme algoritmalarını entegre etmek.
    3. Kullanıcıların veri dağılımını izleyebileceği bir arayüz geliştirmek.

### Projenin Yenilikçi Yönü (Özgün Değeri)
1.  **Sunucusuz Yedeklilik:** Düğümlerin %30'u çevrimdışı olsa dahi veri kayıpsız kurtarılır.
2.  **Sıfır Güven Mimarisi:** Düğümler verinin tamamına veya şifreleme anahtarına sahip değildir.
3.  **Ölçeklenebilirlik:** Ağa katılan her yeni cihaz kapasiteyi artırır.

### Projenin Yöntemi
Go (Golang) veya Rust dilleri kullanılarak geliştirilecektir. 
1. **Veri Parçalama:** "Erasure Coding" ile N adet veri ve M adet kurtarma parçasına bölünür.
2. **Kriptografi:** Her parça cihazda AES-256 ile şifrelenir.
3. **P2P Dağıtım:** Dağıtık Hash Tablosu (DHT) algoritması ile parçalar dağıtılır.
4. **Kurtarma:** Açık düğümlerden parçalar paralel çekilerek birleştirilir.

### Tahmini Proje Bütçesi
*   Geliştirici Lisansları: 10.000 TL
*   Test Bulut Sunucuları (Node Kiralama): 15.000 TL
*   Sızma Testi Yazılımları: 10.000 TL
*   Toplam: 35.000 TL
