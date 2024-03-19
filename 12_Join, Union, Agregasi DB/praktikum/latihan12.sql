-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Mar 19, 2024 at 03:44 PM
-- Server version: 10.4.32-MariaDB
-- PHP Version: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `latihan12`
--

-- --------------------------------------------------------

--
-- Table structure for table `operators`
--

CREATE TABLE `operators` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `operators`
--

INSERT INTO `operators` (`id`, `name`, `created_at`, `updated_at`) VALUES
(1, 'john doe', '2024-03-19 10:42:21', '2024-03-19 10:42:21'),
(2, 'John Doe', '2024-03-19 10:45:41', '2024-03-19 10:45:41'),
(3, 'John Doe2', '2024-03-19 10:46:14', '2024-03-19 10:46:14'),
(4, 'John Doe3', '2024-03-19 10:46:14', '2024-03-19 10:46:14'),
(5, 'John Doe4', '2024-03-19 10:46:14', '2024-03-19 10:46:14');

-- --------------------------------------------------------

--
-- Table structure for table `payment_methods`
--

CREATE TABLE `payment_methods` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `status` smallint(6) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `payment_methods`
--

INSERT INTO `payment_methods` (`id`, `name`, `status`, `created_at`, `updated_at`) VALUES
(1, 'Transfer', 1, '2024-03-19 11:26:30', '2024-03-19 11:26:30'),
(2, 'COD', 1, '2024-03-19 11:26:30', '2024-03-19 11:26:30'),
(3, 'Pay Later', 1, '2024-03-19 11:26:30', '2024-03-19 11:26:30');

-- --------------------------------------------------------

--
-- Table structure for table `products`
--

CREATE TABLE `products` (
  `id` int(11) NOT NULL,
  `product_type_id` int(11) NOT NULL,
  `operator_id` int(11) NOT NULL,
  `code` varchar(50) NOT NULL,
  `name` varchar(100) NOT NULL,
  `status` smallint(6) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `products`
--

INSERT INTO `products` (`id`, `product_type_id`, `operator_id`, `code`, `name`, `status`, `created_at`, `updated_at`) VALUES
(1, 1, 3, 'Asus', 'Rog', 0, '2024-03-19 10:49:51', '2024-03-19 10:49:51'),
(2, 1, 3, 'Asus', 'Rog2', 1, '2024-03-19 11:21:15', '2024-03-19 11:21:15'),
(3, 2, 1, 'Asus', 'TUF', 0, '2024-03-19 11:21:15', '2024-03-19 11:21:15'),
(4, 2, 1, 'Asus', 'TUF2', 1, '2024-03-19 11:21:15', '2024-03-19 11:21:15'),
(5, 2, 1, 'Asus', 'TUF3', 1, '2024-03-19 11:21:15', '2024-03-19 11:21:15'),
(6, 3, 4, 'Lenovo', 'Legion', 1, '2024-03-19 11:21:15', '2024-03-19 11:21:15'),
(7, 3, 4, 'Lenovo', 'Legion2', 0, '2024-03-19 11:21:15', '2024-03-19 11:21:15'),
(8, 3, 4, 'Lenovo', 'Legion3', 0, '2024-03-19 11:21:15', '2024-03-19 11:21:15');

-- --------------------------------------------------------

--
-- Table structure for table `products_decsription`
--

CREATE TABLE `products_decsription` (
  `id` int(11) NOT NULL,
  `description` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `product_types`
--

CREATE TABLE `product_types` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `product_types`
--

INSERT INTO `product_types` (`id`, `name`, `created_at`, `updated_at`) VALUES
(1, 'Laptop', '2024-03-19 10:48:28', '2024-03-19 10:48:28'),
(2, 'Desktop', '2024-03-19 10:48:28', '2024-03-19 10:48:28'),
(3, 'Smartphone', '2024-03-19 10:48:28', '2024-03-19 10:48:28');

-- --------------------------------------------------------

--
-- Table structure for table `transactions`
--

CREATE TABLE `transactions` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `payment_method_id` int(11) NOT NULL,
  `status` varchar(10) NOT NULL,
  `total_qty` int(11) NOT NULL,
  `total_price` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `transactions`
--

INSERT INTO `transactions` (`id`, `user_id`, `payment_method_id`, `status`, `total_qty`, `total_price`, `created_at`, `updated_at`) VALUES
(1, 1, 1, '1', 2, 50, '2024-03-19 13:01:05', '2024-03-19 13:01:05'),
(2, 1, 2, '1', 3, 75, '2024-03-19 13:01:05', '2024-03-19 13:01:05'),
(3, 1, 3, '0', 1, 25, '2024-03-19 13:01:05', '2024-03-19 13:01:05'),
(4, 2, 1, '1', 4, 100, '2024-03-19 13:01:05', '2024-03-19 13:01:05'),
(5, 2, 2, '0', 2, 50, '2024-03-19 13:01:05', '2024-03-19 13:01:05'),
(6, 2, 3, '1', 3, 75, '2024-03-19 13:01:05', '2024-03-19 13:01:05'),
(7, 3, 1, '0', 1, 25, '2024-03-19 13:01:05', '2024-03-19 13:01:05'),
(8, 3, 2, '1', 3, 75, '2024-03-19 13:01:05', '2024-03-19 13:01:05'),
(9, 3, 3, '1', 2, 50, '2024-03-19 13:01:05', '2024-03-19 13:01:05'),
(10, 4, 1, '1', 2, 50, '2024-03-19 13:01:05', '2024-03-19 13:01:05'),
(11, 4, 2, '0', 1, 25, '2024-03-19 13:01:05', '2024-03-19 13:01:05'),
(12, 4, 3, '1', 4, 100, '2024-03-19 13:01:05', '2024-03-19 13:01:05'),
(13, 5, 1, '1', 3, 75, '2024-03-19 13:01:05', '2024-03-19 13:01:05'),
(14, 5, 2, '1', 2, 50, '2024-03-19 13:01:05', '2024-03-19 13:01:05'),
(15, 5, 3, '0', 1, 25, '2024-03-19 13:01:05', '2024-03-19 13:01:05');

-- --------------------------------------------------------

--
-- Table structure for table `transactions_details`
--

CREATE TABLE `transactions_details` (
  `transaction_id` int(11) NOT NULL,
  `product_id` int(11) NOT NULL,
  `status` varchar(10) NOT NULL,
  `qty` int(11) NOT NULL,
  `price` int(25) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `transactions_details`
--

INSERT INTO `transactions_details` (`transaction_id`, `product_id`, `status`, `qty`, `price`, `created_at`, `updated_at`) VALUES
(1, 1, '1', 2, 25, '2024-03-19 13:12:17', '2024-03-19 13:12:17'),
(1, 2, '0', 1, 15, '2024-03-19 13:12:17', '2024-03-19 13:12:17'),
(1, 3, '1', 3, 40, '2024-03-19 13:12:17', '2024-03-19 13:12:17'),
(2, 4, '1', 1, 10, '2024-03-19 13:12:17', '2024-03-19 13:12:17'),
(2, 5, '1', 4, 50, '2024-03-19 13:12:17', '2024-03-19 13:12:17'),
(2, 6, '1', 2, 30, '2024-03-19 13:12:17', '2024-03-19 13:12:17'),
(3, 7, '0', 2, 20, '2024-03-19 13:12:17', '2024-03-19 13:12:17'),
(3, 8, '1', 3, 45, '2024-03-19 13:12:17', '2024-03-19 13:12:17'),
(3, 1, '1', 1, 15, '2024-03-19 13:12:17', '2024-03-19 13:12:17'),
(4, 2, '1', 3, 40, '2024-03-19 13:12:17', '2024-03-19 13:12:17'),
(4, 3, '1', 2, 25, '2024-03-19 13:12:17', '2024-03-19 13:12:17'),
(4, 4, '1', 4, 55, '2024-03-19 13:12:17', '2024-03-19 13:12:17'),
(5, 5, '1', 1, 10, '2024-03-19 13:12:17', '2024-03-19 13:12:17'),
(5, 6, '0', 2, 30, '2024-03-19 13:12:17', '2024-03-19 13:12:17'),
(5, 7, '1', 3, 45, '2024-03-19 13:12:17', '2024-03-19 13:12:17'),
(6, 8, '1', 4, 55, '2024-03-19 13:12:17', '2024-03-19 13:12:17'),
(6, 1, '1', 1, 15, '2024-03-19 13:12:17', '2024-03-19 13:12:17'),
(6, 2, '0', 2, 25, '2024-03-19 13:12:17', '2024-03-19 13:12:17'),
(7, 3, '0', 3, 40, '2024-03-19 13:12:17', '2024-03-19 13:12:17'),
(7, 4, '1', 1, 10, '2024-03-19 13:12:17', '2024-03-19 13:12:17'),
(7, 5, '1', 2, 30, '2024-03-19 13:12:17', '2024-03-19 13:12:17'),
(8, 6, '1', 3, 45, '2024-03-19 13:12:17', '2024-03-19 13:12:17'),
(8, 7, '1', 4, 55, '2024-03-19 13:12:17', '2024-03-19 13:12:17'),
(8, 8, '0', 1, 15, '2024-03-19 13:12:17', '2024-03-19 13:12:17'),
(9, 1, '1', 2, 25, '2024-03-19 13:12:17', '2024-03-19 13:12:17'),
(9, 2, '0', 1, 15, '2024-03-19 13:12:17', '2024-03-19 13:12:17'),
(9, 3, '1', 3, 40, '2024-03-19 13:12:17', '2024-03-19 13:12:17'),
(10, 4, '0', 1, 10, '2024-03-19 13:12:17', '2024-03-19 13:12:17'),
(10, 5, '1', 4, 50, '2024-03-19 13:12:17', '2024-03-19 13:12:17'),
(10, 6, '0', 2, 30, '2024-03-19 13:12:17', '2024-03-19 13:12:17'),
(11, 7, '1', 2, 20, '2024-03-19 13:12:17', '2024-03-19 13:12:17'),
(11, 8, '1', 3, 45, '2024-03-19 13:12:17', '2024-03-19 13:12:17'),
(11, 1, '1', 1, 15, '2024-03-19 13:12:17', '2024-03-19 13:12:17'),
(12, 2, '0', 3, 40, '2024-03-19 13:12:17', '2024-03-19 13:12:17'),
(12, 3, '1', 2, 25, '2024-03-19 13:12:17', '2024-03-19 13:12:17'),
(12, 4, '1', 4, 55, '2024-03-19 13:12:17', '2024-03-19 13:12:17'),
(13, 5, '1', 1, 10, '2024-03-19 13:12:17', '2024-03-19 13:12:17'),
(13, 6, '1', 2, 30, '2024-03-19 13:12:17', '2024-03-19 13:12:17'),
(13, 7, '0', 3, 45, '2024-03-19 13:12:17', '2024-03-19 13:12:17'),
(14, 8, '1', 4, 55, '2024-03-19 13:12:17', '2024-03-19 13:12:17'),
(14, 1, '0', 1, 15, '2024-03-19 13:12:17', '2024-03-19 13:12:17'),
(14, 2, '1', 2, 25, '2024-03-19 13:12:17', '2024-03-19 13:12:17'),
(15, 3, '1', 3, 40, '2024-03-19 13:12:17', '2024-03-19 13:12:17'),
(15, 4, '0', 1, 10, '2024-03-19 13:12:17', '2024-03-19 13:12:17'),
(15, 5, '1', 2, 30, '2024-03-19 13:12:17', '2024-03-19 13:12:17');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `status` smallint(6) NOT NULL,
  `dob` date NOT NULL,
  `gender` char(1) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `status`, `dob`, `gender`, `created_at`, `updated_at`) VALUES
(1, 'Doni', 1, '2017-03-08', '1', '2024-03-19 11:29:52', '2024-03-19 11:29:52'),
(2, 'Rani', 0, '2015-03-19', '0', '2024-03-19 11:29:52', '2024-03-19 11:29:52'),
(3, 'Aldo', 1, '2018-05-16', '1', '2024-03-19 11:29:52', '2024-03-19 11:29:52'),
(4, 'Dimas', 1, '2013-08-21', '1', '2024-03-19 11:29:52', '2024-03-19 11:29:52'),
(5, 'Putri', 0, '2012-12-21', '0', '2024-03-19 11:29:52', '2024-03-19 11:29:52');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `operators`
--
ALTER TABLE `operators`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `payment_methods`
--
ALTER TABLE `payment_methods`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`id`),
  ADD KEY `operator_id` (`operator_id`),
  ADD KEY `product_type_id` (`product_type_id`);

--
-- Indexes for table `products_decsription`
--
ALTER TABLE `products_decsription`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `product_types`
--
ALTER TABLE `product_types`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`id`),
  ADD KEY `payment_method_id` (`payment_method_id`),
  ADD KEY `user_id` (`user_id`);

--
-- Indexes for table `transactions_details`
--
ALTER TABLE `transactions_details`
  ADD KEY `product_id` (`product_id`),
  ADD KEY `transaction_id` (`transaction_id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `operators`
--
ALTER TABLE `operators`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `payment_methods`
--
ALTER TABLE `payment_methods`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `products`
--
ALTER TABLE `products`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT for table `products_decsription`
--
ALTER TABLE `products_decsription`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `product_types`
--
ALTER TABLE `product_types`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `transactions`
--
ALTER TABLE `transactions`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=16;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `products`
--
ALTER TABLE `products`
  ADD CONSTRAINT `products_ibfk_1` FOREIGN KEY (`operator_id`) REFERENCES `operators` (`id`),
  ADD CONSTRAINT `products_ibfk_2` FOREIGN KEY (`product_type_id`) REFERENCES `product_types` (`id`);

--
-- Constraints for table `transactions`
--
ALTER TABLE `transactions`
  ADD CONSTRAINT `transactions_ibfk_1` FOREIGN KEY (`payment_method_id`) REFERENCES `payment_methods` (`id`),
  ADD CONSTRAINT `transactions_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

--
-- Constraints for table `transactions_details`
--
ALTER TABLE `transactions_details`
  ADD CONSTRAINT `transactions_details_ibfk_1` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`),
  ADD CONSTRAINT `transactions_details_ibfk_2` FOREIGN KEY (`transaction_id`) REFERENCES `transactions` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
