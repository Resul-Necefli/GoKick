# GoKick 

GoKick — Golang proqramlaşdırma dilində **Layered Architecture (Laylı Arxitektura)** və **Dependency Injection** prinsiplərini öyrənmək məqsədilə yazılmış sadə RESTful API layihəsidir.

Bu layihə "Crowdfunding" (Kütləvi Maliyyələşdirmə) platformasının sadələşdirilmiş backend modelidir.

## Arxitektura

Layihə klassik 3-laylı (3-Tier) arxitektura əsasında qurulub. Asılılıqlar (Dependencies) yuxarıdan aşağıya doğru **Dependency Injection** vasitəsilə ötürülür:

1.  **Handler Layer (`internal/handler`)**: HTTP sorğularını qəbul edir, JSON emal edir və Service layına yönləndirir.
2.  **Service Layer (`internal/service`)**: Biznes məntiqini (validasiya, qaydalar) icra edir.
3.  **Repository Layer (`internal/repository`)**: Məlumatların saxlanmasına cavabdehdir (hazırda In-Memory).

## Texnologiyalar

* **Dil:** Go (v1.24.8)
* **Kitabxanalar:** Go Standard Library (`net/http`, `encoding/json` və s.)
* **Verilənlər Bazası:** In-Memory (`map` strukturu)

## İşə Salma

Layihəni lokal kompüterdə işə salmaq üçün:

```bash
# 1. Repozitoriyanı klonlayın
git clone [https://github.com/Resul-Necefli/GoKick.git](https://github.com/Resul-Necefli/GoKick.git)

# 2. Qovluğa keçin
cd GoKick

# 3. Serveri başladın
go run cmd/api/main.go
``` 

## API Endpoints
POST /campaigns - Yeni kampaniya yarat

GET /campaigns - Bütün kampaniyaları gətir

GET /campaigns/{id} - ID-yə görə kampaniyanı gətir

PUT /campaigns/{id} - Kampaniya məlumatlarını yenilə

POST /campaigns/{id}/donate - Kampaniyaya ianə et


