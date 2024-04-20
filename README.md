# Wifi Ağı Bilgilerini Alma Servisi

Bu basit Go programı, bir Windows işletim sistemi üzerinde çalışarak mevcut Wifi ağlarının SSID'lerini alır ve istemcilere geri döndürür.

## Nasıl Çalışır

1. **Kurulum**: Projeyi klonlayın ve Go'nun yüklü olduğundan emin olun.
2. **Windows İşletim Sistemi**: Bu program yalnızca Windows işletim sistemi üzerinde çalışır.
3. **Sunucuyu Başlatın**: `go run main.go` komutuyla sunucuyu başlatın.
4. **Test Edin**: Tarayıcıdan veya bir API test aracından `http://localhost:8080/getWifi` adresine bir GET isteği yaparak mevcut Wifi ağlarının SSID'lerini alın.

## API Dökümantasyonu

### `/getWifi` Endpoint'i

- **Method**: GET
- **Parametreler**: Herhangi bir parametre gerektirmez.
- **Cevap**: JSON formatında mevcut Wifi ağlarının SSID'lerini içeren bir dizi döndürür.

Örnek istek:
GET http://localhost:8080/getWifi
