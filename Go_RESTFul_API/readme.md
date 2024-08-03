# Catatan Pembelajaran RESTFul API 

## Step-by-Step to Create RESTFul API
1. Membuat API Specification dengan standar OPEN API berformat JSON
2. Membuat Database
3. Membuat Direktori model 
    - Berisi Domain dan web
4. Membuat Repository
5. Membuat Service
6. 


## Arsitektur direktori terdiri dari:
1. Model adalah,direktori ini  terdiri dari:
    - Domain, merupakan representasi dari table database
    - Web, merupakan representasi dari request dan response

2. Repository adalah data akses layer ke domain, direktori ini  terdiri dari:
    - file kontrak dalam bentuk interface,
    - file implementasi dalam bentuk struct

3.  Service adalah business logic pada aplikasi, direktori ini terdiri dari:
    - file kontrak dalam bentuk interface,
    - file implementasi

N.B, Pembuatan interface setiap parameter selalu diawali dengan context.Context