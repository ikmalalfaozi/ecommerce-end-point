# **Ecommerce End Point**

**Implementasi End Point Ecommerce Sederhana Menggunakan Golang (Gin dan Gorm)**

## Table of Contents

-   [Introduction](#introduction)
-   [General Information](#general-information)
-   [Technologies Used](#technologies-used)
-   [Installation and Requirement](#installation-and-requirement)
-   [How to Run](#how-to-run)
-   [How to Use](#how-to-use)
-   [Project Status](#project-status)
-   [Room for Improvement](#room-for-improvement)
-   [Contact](#contact)

## Introduction

Hai, Selamat datang di Repository Github saya!

Proyek ini dibuat oleh saya sendiri, Ikmal Alfaozi

## General Information

Dalam project ini, saya mengimplementasikan End Point Ecommerce sederhana dengan menggunakan bahasa pemrograman Golang khususnya menggunakan framework gin dan gorm. End point yang dibuat di antaraynya adalah

-   GET list of product
-   GET specific product by id
-   POST new product and add to database
-   GET list of customers
-   GET specific customer by id
-   POST new customers and add to database
-   Checkout product by product id
-   Buy products directly without checkout
-   GET list of product in the cart
-   GET product in the chart by id
-   Buy the product in the cart by id
-   Update the number of products in the cart
-   Update order status
-   GET list of orders
-   GET specific order by id

## Technologies Used

Implementasi program ditulis dalam bahasa Go dan database MySQL.

## Installation and Requirement

-   Unduh seluruh folder dan file pada repository ini atau clone repository
-   Unduh dan pasang [Go](https://go.dev/dl/)
-   Unduh dan pasang [MySQL](https://dev.mysql.com/downloads/)

## How to Run

-   Pastikan telah melakukan prosedur [Installation and Requirement](#installation-and-requirement)
-   Buka terminal dan change directory ke folder `src`
    ```
    cd src
    ```
-   Lakukan instalasi terhadap requirement
    ```
    go mod tidy
    ```
-   Buat database dengan nama `ecommerce-end-point`
-   Import file `ecommerce-end-point.sql` ke database `ecommerce-end-point`
-   Pastikan dsn yang ada di file `src/internal/delivery/router.go` sesuai dengan user, password, dan nama database yang dibuat
    ```
    dsn := user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
    ```
-   Jalankan program dengan perintah
    ```
    go run main.go
    ```

## How to Use

Contoh penggunaan end point ecommerce ini dapat dilihat pada folder documentation

## Project Status

Project status: _finish_

## Room for Improvement

Pengembangan end point ecommerce ini perlu menerapkan middleware untuk meningkatkan keamanan siapa yang boleh mengakses end point tersebut. Selain itu, end point ecommerce ini juga perlu menambahkan end point untuk _payment_ (pembayaran) agar customer dapat melakukan pembayaran product yang telah dipesan.

## Contact

[ikmalalfaozi@gmail.com](mailto:ikmalalfaozi@gmail.com)
