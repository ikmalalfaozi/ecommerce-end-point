-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Nov 05, 2022 at 11:26 AM
-- Server version: 10.4.22-MariaDB
-- PHP Version: 8.1.2

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `ecommerce-end-point`
--

-- --------------------------------------------------------

--
-- Table structure for table `carts`
--

CREATE TABLE `carts` (
  `customer_id` int(11) NOT NULL,
  `product_id` int(11) NOT NULL,
  `quantity` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `carts`
--

INSERT INTO `carts` (`customer_id`, `product_id`, `quantity`) VALUES
(1, 5, 3),
(2, 6, 1),
(2, 7, 1);

-- --------------------------------------------------------

--
-- Table structure for table `customers`
--

CREATE TABLE `customers` (
  `customer_id` int(11) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `first_name` varchar(50) NOT NULL,
  `last_name` varchar(50) NOT NULL,
  `address` varchar(255) NOT NULL,
  `zip` varchar(15) NOT NULL,
  `phone` varchar(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `customers`
--

INSERT INTO `customers` (`customer_id`, `email`, `password`, `first_name`, `last_name`, `address`, `zip`, `phone`) VALUES
(1, 'user@gmail.com', '$2a$10$rJZH4vqctPv0XSKaQj/lA.Aue5XdiRhMAEsyRzMCl9UgpcaRuq08a', 'user', 'user', 'Jl. Ahmad Yani No 58 ...', '53262', '085158887776'),
(2, 'user2@gmail.com', '$2a$10$Nc7qOgMZEShs.53sXlqOh.hCoqxCqObn1Tv.hz.AYJPk90bPgHUjq', 'user2', 'user2', 'Jl. Ahmad Yani No 58 ...', '53262', '085158887776'),
(3, 'user3@gmail.com', '$2a$10$gjK98HkeF5VPqxvWdKHTFOZzDLw0XHvJjdsjIatZxI0zHM8g6u0cm', 'user3', 'user3', 'Jl. Ahmad Yani No 58 ...', '53262', '085158887776'),
(4, 'user4@gmail.com', '$2a$10$uQ.y0u8VUvXMTCAjufO8UeBH3Lm2XbI11P8qmX7p2JZ0eE29zu7bW', 'user4', 'user4', 'Jl. Ahmad Yani No 58 ...', '53262', '085158887776'),
(5, 'user5@gmail.com', '$2a$10$exqXyCCO42AGvO09o0ppGe.r8egfQho39qee.EX7Ar/0SDvyIGlfS', 'user5', 'user5', 'Jl. Ahmad Yani No 58 ...', '53262', '085158887776'),
(6, 'user6@gmail.com', '$2a$10$7vjArrKPGxS2IKN5kUDfxOXkzR4.HzneinWaHtHD.Wbu6siZhNFcq', 'user6', 'user6', 'Jl. Ahmad Yani No 58 ...', '53262', '085158887776'),
(7, 'user7@gmail.com', '$2a$10$HA/Ep4JGWUTLtewRVd8XguKeonxQP8wEv5G.ZOxpQHfPKyQXV7kC.', 'user7', 'user7', 'Jl. Ahmad Yani No 58 ...', '53262', '085158887776'),
(8, 'user8@gmail.com', '$2a$10$h1w4IH5di63iWBdhB6FqpulWBetGoBT.kIU8BgI6/07774tPy1Kga', 'user8', 'user8', 'Jl. Ahmad Yani No 58 ...', '53262', '085158887776'),
(9, 'user9@gmail.com', '$2a$10$7IFQqv1AwEZMaW6Mw.DZVeGV8fG6g1F7wWjvlWHMG1lO5P/ZFW4tW', 'user9', 'user9', 'Jl. Ahmad Yani No 58 ...', '53262', '085158887776'),
(10, 'user10@gmail.com', '$2a$10$X0FCr9vPAlhzJ/Q4beM2bOHljeJqHTjba1ItPxcHUiepno/IlZvZu', 'user10', 'user10', 'Jl. Ahmad Yani No 58 ...', '53262', '085158887776'),
(11, 'user11@gmail.com', '$2a$10$vxxltrg/KUTGtH1OrFOIdenHT44YbHGEAtGumuyoLQ7oAvYmRk3zu', 'user11', 'user11', 'Jl. Ahmad Yani No 58 ...', '53262', '085158887776'),
(12, 'user12@gmail.com', '$2a$10$ZT/ViTkKCJIE3kuPCfobyOz.AS1jso19JzEm1v1FhQoGUBt5j5jf.', 'user12', 'user12', 'Jl. Ahmad Yani No 58 ...', '53262', '085158887776'),
(13, 'user13@gmail.com', '$2a$10$qgu.lW4cZ9B4NjanIy/To.YSIt4FHz03XEqoGUYaDbwFziyipmUUW', 'user13', 'user13', 'Jl. Ahmad Yani No 58 ...', '53262', '085158887776'),
(14, 'user14@gmail.com', '$2a$10$eLhqmRFQwQ/hsfPUiNDB6./tDSSl2zrQZVHY/uMc0KNqaQ0H8QmdO', 'user14', 'user14', 'Jl. Ahmad Yani No 58 ...', '53262', '085158887776'),
(16, 'user15@gmail.com', '$2a$10$Xhvo7l8inGMFzA4nSZulZ.AQOxEeQhaHZ37.PcVoD17R.iHJ4umNS', 'user15', 'user15', 'Jl. Ahmad Yani No 58 ...', '53262', '085158887776');

-- --------------------------------------------------------

--
-- Table structure for table `orders`
--

CREATE TABLE `orders` (
  `order_id` int(11) NOT NULL,
  `customer_id` int(11) NOT NULL,
  `product_id` int(11) NOT NULL,
  `quantity` int(11) NOT NULL,
  `status` varchar(20) NOT NULL,
  `order_date` datetime NOT NULL DEFAULT current_timestamp(),
  `shipped_date` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `orders`
--

INSERT INTO `orders` (`order_id`, `customer_id`, `product_id`, `quantity`, `status`, `order_date`, `shipped_date`) VALUES
(1, 3, 2, 1, 'Sent', '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(2, 3, 3, 1, 'Sent', '2022-11-04 17:15:47', '0000-00-00 00:00:00'),
(3, 1, 2, 4, 'Processed', '2022-11-04 17:27:33', '0000-00-00 00:00:00'),
(4, 1, 4, 1, 'Processed', '2022-11-04 17:36:48', '0000-00-00 00:00:00'),
(5, 1, 4, 1, 'Processed', '2022-11-04 17:37:59', '0000-00-00 00:00:00'),
(6, 1, 4, 1, 'Processed', '2022-11-04 17:42:58', '0000-00-00 00:00:00'),
(7, 3, 3, 1, 'Processed', '2022-11-05 16:46:31', '0000-00-00 00:00:00'),
(8, 1, 5, 4, 'Processed', '2022-11-05 16:47:22', '0000-00-00 00:00:00');

-- --------------------------------------------------------

--
-- Table structure for table `products`
--

CREATE TABLE `products` (
  `product_id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `price` bigint(20) NOT NULL,
  `quantity` int(11) NOT NULL,
  `image_url` varchar(255) NOT NULL,
  `description` text NOT NULL,
  `category` varchar(255) NOT NULL,
  `sku` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `products`
--

INSERT INTO `products` (`product_id`, `name`, `price`, `quantity`, `image_url`, `description`, `category`, `sku`) VALUES
(2, 'Adidas', 200000, 5, 'https://source.unsplash.com/random/900x700/?shoes', 'Ini adalah sepatu Adidas', 'Fashion', 'Adidas Shoes Hitam Ukuran 42'),
(3, 'Adidas', 200000, 8, 'https://source.unsplash.com/random/900x700/?shoes', 'Ini adalah sepatu Adidas', 'Fashion', 'Adidas Shoes Hitam Ukuran 42'),
(4, 'Adidas', 200000, 7, 'https://source.unsplash.com/random/900x700/?shoes', 'Ini adalah sepatu Adidas', 'Fashion', 'Adidas Shoes Hitam Ukuran 42'),
(5, 'Adidas', 200000, 6, 'https://source.unsplash.com/random/900x700/?shoes', 'Ini adalah sepatu Adidas', 'Fashion', 'Adidas Shoes Hitam Ukuran 42'),
(6, 'Adidas', 200000, 10, 'https://source.unsplash.com/random/900x700/?shoes', 'Ini adalah sepatu Adidas', 'Fashion', 'Adidas Shoes Hitam Ukuran 42'),
(7, 'Adidas', 200000, 10, 'https://source.unsplash.com/random/900x700/?shoes', 'Ini adalah sepatu Adidas', 'Fashion', 'Adidas Shoes Hitam Ukuran 42'),
(8, 'Adidas', 200000, 10, 'https://source.unsplash.com/random/900x700/?shoes', 'Ini adalah sepatu Adidas', 'Fashion', 'Adidas Shoes Hitam Ukuran 42'),
(9, 'Adidas', 200000, 10, 'https://source.unsplash.com/random/900x700/?shoes', 'Ini adalah sepatu Adidas', 'Fashion', 'Adidas Shoes Hitam Ukuran 42'),
(10, 'Adidas', 200000, 10, 'https://source.unsplash.com/random/900x700/?shoes', 'Ini adalah sepatu Adidas', 'Fashion', 'Adidas Shoes Hitam Ukuran 42'),
(11, 'Adidas', 200000, 10, 'https://source.unsplash.com/random/900x700/?shoes', 'Ini adalah sepatu Adidas', 'Fashion', 'Adidas Shoes Hitam Ukuran 42'),
(12, 'Adidas', 200000, 10, 'https://source.unsplash.com/random/900x700/?shoes', 'Ini adalah sepatu Adidas', 'Fashion', 'Adidas Shoes Hitam Ukuran 42'),
(13, 'Adidas', 200000, 10, 'https://source.unsplash.com/random/900x700/?shoes', 'Ini adalah sepatu Adidas', 'Fashion', 'Adidas Shoes Hitam Ukuran 42'),
(14, 'Adidas', 200000, 10, 'https://source.unsplash.com/random/900x700/?shoes', 'Ini adalah sepatu Adidas', 'Fashion', 'Adidas Shoes Hitam Ukuran 42'),
(15, 'Redmi Note 10', 200000, 10, 'https://source.unsplash.com/random/900x700/?handphone', 'Ini adalah HP Redmi Note 10 ...', 'Elektronik', 'Redmi Note 10'),
(16, 'Redmi Note 9', 200000, 10, 'https://source.unsplash.com/random/900x700/?handphone', 'Ini adalah HP Redmi Note 9 ...', 'Elektronik', 'Redmi Note 9');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `carts`
--
ALTER TABLE `carts`
  ADD PRIMARY KEY (`customer_id`,`product_id`),
  ADD KEY `fk_chekouts_product_id` (`product_id`);

--
-- Indexes for table `customers`
--
ALTER TABLE `customers`
  ADD PRIMARY KEY (`customer_id`),
  ADD UNIQUE KEY `email` (`email`);

--
-- Indexes for table `orders`
--
ALTER TABLE `orders`
  ADD PRIMARY KEY (`order_id`),
  ADD KEY `fk_orders_customer_id` (`customer_id`),
  ADD KEY `fk_orders_product_id` (`product_id`);

--
-- Indexes for table `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`product_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `customers`
--
ALTER TABLE `customers`
  MODIFY `customer_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=17;

--
-- AUTO_INCREMENT for table `orders`
--
ALTER TABLE `orders`
  MODIFY `order_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT for table `products`
--
ALTER TABLE `products`
  MODIFY `product_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=17;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `carts`
--
ALTER TABLE `carts`
  ADD CONSTRAINT `fk_chekouts_customer_id` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`customer_id`),
  ADD CONSTRAINT `fk_chekouts_product_id` FOREIGN KEY (`product_id`) REFERENCES `products` (`product_id`);

--
-- Constraints for table `orders`
--
ALTER TABLE `orders`
  ADD CONSTRAINT `fk_orders_customer_id` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`customer_id`),
  ADD CONSTRAINT `fk_orders_product_id` FOREIGN KEY (`product_id`) REFERENCES `products` (`product_id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
