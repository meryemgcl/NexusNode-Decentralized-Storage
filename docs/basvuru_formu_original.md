# KASTAMONU ÜNİVERSİTESİ 1. AR-GE PROJE PAZARI BAŞVURU FORMU

*(Aşağıdaki metinleri doğrudan yüklediğiniz PDF/Word formundaki ilgili kutucuklara kopyalayabilirsiniz.)*

---

**Proje Başlığı:** NexusNode: Kurumsal Veri Egemenliği İçin Eşler Arası (P2P) Merkeziyetsiz Depolama Ağı
**Başvuru Alanı:** Fen Bilimleri / Yazılım ve Bilişim Teknolojileri

---

### Proje Özeti ve Anahtar Kelimeler
Günümüzde kurumlar ve bireyler, kritik verilerini merkezi bulut sağlayıcılarında (Google, AWS vb.) barındırmakta olup; bu durum veri egemenliği ihlallerine, tek nokta hatalarına (single point of failure) ve yüksek abonelik maliyetlerine neden olmaktadır. Bu projenin amacı, internete bağlı sıradan cihazların birer sunucu gibi çalışmasını sağlayan Eşler Arası (P2P) merkeziyetsiz bir dosya depolama ve paylaşım ağı geliştirmektir. Sistem, yüklenen dosyaları "Silinti Kodlaması (Erasure Coding)" algortimaları ile matematiksel parçalara böler, şifreler ve ağdaki uç düğümlere (node) dağıtır. Donanım ve yapay zeka kullanılmayan bu saf yazılım projesi, merkezi bir sunucu çöktüğünde bile verinin %100 kayıpsız ve yüksek güvenlikli olarak geri çağrılabilmesini sağlar.
**Anahtar Kelimeler:** Merkeziyetsiz Sistemler, P2P Ağlar, Veri Egemenliği, Silinti Kodlaması, Siber Güvenlik.

### Projenin Amaç ve Hedefleri
*   **Amaç:** Kurumların kendi donanım altyapılarını kullanarak kapalı veya açık, sansürlenemez ve veri kaybı riski sıfır olan merkeziyetsiz bir depolama ağı kurmalarını sağlayacak yazılım protokolünü geliştirmek.
*   **Hedefler:**
    1. İstemciler arası doğrudan iletişimi sağlayacak WebRTC/libp2p tabanlı çekirdek ağ mimarisini kurmak.
    2. Dosya parçalanması ve şifrelenmesi için AES-256 ve Reed-Solomon hata düzeltme algoritmalarını entegre etmek.
    3. Kullanıcıların ağ durumunu ve veri dağılımını izleyebileceği bir arayüz geliştirmek.

### Projenin Yenilikçi Yönü (Özgün Değeri)
1.  **Sunucusuz (Serverless) Yedeklilik:** Veri tek bir veri merkezinde değil, ağdaki tüm cihazlara parçalanmış olarak dağıtılır. Düğümlerin %30'u çevrimdışı olsa dahi veri kayıpsız kurtarılır.
2.  **Sıfır Güven (Zero Trust) Mimarisi:** Parçaları barındıran hiçbir cihaz, verinin tamamına veya şifreleme anahtarına sahip değildir.
3.  **Ölçeklenebilirlik:** Sisteme katılan her yeni cihaz, ağın depolama kapasitesini ve hızını artırır (Bant genişliği paylaşımı).

### Ticari Potansiyeli ve İhtiyaç Durumu
KVKK (Kişisel Verilerin Korunması Kanunu) ve ticari veri gizliliği yasaları, kamu kurumlarının ve hastanelerin verilerini yurtdışı bulut sistemlerinde tutmasını zorlaştırmaktadır. Bu yazılım, kurumlara "kendi donanımlarıyla kendi özel bulutlarını yaratma" imkanı sunarak devasa bulut faturalarını ortadan kaldırır. B2B (Kurumdan Kuruma) lisanslama modeli ile satış potansiyeli son derece yüksektir.

### Projenin Hedef Pazarı veya Etki Alanı
*   **Kamu ve Sağlık Sektörü:** KVKK gereği hasta verilerini yerel tutmak zorunda olan hastaneler ve kamu idareleri.
*   **Özel Sektör:** Ticari sırlarını merkezi sunucularda tutmak istemeyen savunma sanayii ve teknoloji firmaları.
*   **Etki Alanı:** Bulut bilişim maliyetlerini düşürmek isteyen tüm Bilişim (IT) ekosistemi.

### Projenin Yöntemi
Proje, yüksek eşzamanlılık (concurrency) sağlayan Go (Golang) veya Rust dilleri kullanılarak geliştirilecektir. 
1. **Veri Parçalama (Sharding):** Yüklenen veri "Erasure Coding" ile N adet veri ve M adet kurtarma parçasına bölünür.
2. **Kriptografi:** Her parça cihazda izole olarak AES-256 ile şifrelenir.
3. **P2P Dağıtım:** Dağıtık Hash Tablosu (DHT) algoritması ile parçalar en uygun düğümlere iletilir.
4. **Kurtarma:** Kullanıcı indirme isteği yolladığında, açık düğümlerden parçalar eşzamanlı (paralel) çekilerek birleştirilir.

### Projenin Yapılabilirliği ve Sürdürülebilirliği
Proje fiziksel donanım üretimi içermediği için tedarik zinciri veya üretim riskleri taşımaz. Saf yazılım mühendisliği prensipleriyle açık kaynaklı kriptografi kütüphaneleri (libp2p) üzerine inşa edilecektir. Çekirdek algoritma yazıldıktan sonra platform bağımsız (Windows, Linux, Mac) olarak tüm cihazlarda çalıştırılabilir, bu da yüksek sürdürülebilirlik sağlar.

### Tahmini Proje Bütçesi ve Gerekçesi
*   **Geliştirici Lisansları ve Geliştirme Ortamı Giderleri:** 10.000 TL
*   **Testler İçin Bulut Sunucu (Node) Kiralama:** 15.000 TL (Ağ simülasyonu ve stres testleri için sanal makineler)
*   **Güvenlik/Sızma Testi (Pentest) Yazılımları:** 10.000 TL
*   **Toplam Tahmini Bütçe:** 35.000 TL (Tamamen Ar-Ge ve ağ stres testlerine yönelik bütçedir).

### Proje Çıktıları ve Kazanımlar
*   **Bilimsel Çıktı:** Merkeziyetsiz ağlarda veri aktarım hızını optimize eden algoritmalar üzerine akademik yayın potansiyeli.
*   **Ekonomik Fayda:** Yabancı bulut şirketlerine ödenen döviz çıkışını engelleyecek yerli bir ağ protokolü.
*   **Sosyal Fayda:** Veri güvenliğini artırarak siber şantaj (Ransomware) olaylarına karşı kesin bir çözüm sunar.

### Referanslar
1. Benet, J. (2014). IPFS - Content Addressed, Versioned, P2P File System.
2. Plank, J. S. (2013). Erasure Codes for Storage Systems: A Brief Primer. USENIX.
