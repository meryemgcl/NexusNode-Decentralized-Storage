# NexusNode P2P Depolama Ağı - Proje Geliştirme Fazları

Bu belge, projenin basit bir prototipten başlayarak devasa bir ekosisteme dönüşmesini sağlayacak genişletilebilir geliştirme fazlarını içermektedir.

### FAZ 1: Çekirdek Ağ ve Parçalama (Core & Sharding) - [Mevcut Odak]
*   İki bilgisayarın (Node) merkezi bir sunucu olmadan birbiriyle doğrudan (P2P) iletişim kurabilmesi (WebRTC veya libp2p entegrasyonu).
*   Yüklenen bir metin dosyasının matematiksel olarak 3 parçaya bölünmesi (Sharding) ve ağa dağıtılması.
*   Dağıtılan parçaların geri çağrılarak orijinal dosyanın kayıpsız birleştirilmesi.

### FAZ 2: Kriptografi ve Yedeklilik (Security & Redundancy)
*   AES-256 uçtan uca şifreleme modülünün yazılması. Verinin ağa şifreli dağıtılması.
*   Reed-Solomon (Silinti Kodlaması) algoritmasının entegrasyonu. Örneğin veri 5 parçaya bölünür, 3 parça ile dosya kurtarılabilir hale getirilir (Düğüm kopmalarına karşı tolerans).
*   Çevrimdışı olan düğümlerin (Node) tespit edilip eksilen veri parçalarının diğer aktif düğümlerde yeniden kopyalanması (Self-Healing / Kendi Kendini Onarma mekanizması).

### FAZ 3: Kullanıcı Arayüzü ve Masaüstü İstemcisi
*   Terminalden çıkan sistemin React veya Electron.js ile kullanıcı dostu bir masaüstü uygulamasına dönüştürülmesi.
*   Kullanıcıların ağda ne kadar alan paylaştığını ve sistem sağlığını gösteren bir "Dashboard" yapılması.

### FAZ 4: İleri Seviye Genişleme (İnovasyon Kapıları)
*   **Merkeziyetsiz Kimlik (DID) Entegrasyonu:** Sisteme giriş yapanların şifre yerine kriptografik cüzdanları ile bağlanması.
*   **Token Ekonomisi:** Kendi diskini ağa kiralayan kullanıcıların sistem içi "Kredi/Token" kazanması.
*   **Federe Versiyon Kontrol:** Bu ağın üzerine inşa edilmiş, yazılımcılar veya 3D tasarımcılar için devasa dosyaları tutabilen merkeziyetsiz bir Git sistemi.
