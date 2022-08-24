-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 25 Agu 2022 pada 00.46
-- Versi server: 10.4.20-MariaDB
-- Versi PHP: 8.0.8

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `campaign_backend`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `campaigns`
--

CREATE TABLE `campaigns` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `short_description` varchar(255) NOT NULL,
  `description` text NOT NULL,
  `goal_amount` int(11) NOT NULL,
  `current_amount` int(11) NOT NULL,
  `perks` text NOT NULL,
  `backer_count` int(11) NOT NULL,
  `slug` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `campaigns`
--

INSERT INTO `campaigns` (`id`, `user_id`, `name`, `short_description`, `description`, `goal_amount`, `current_amount`, `perks`, `backer_count`, `slug`, `created_at`, `updated_at`) VALUES
(1, 1, 'Campaign 1', 'Campaign 1Campaign 1Campaign 1Campaign 1Campaign 1', 'Campaign 1Campaign 1Campaign 1Campaign 1Campaign 1Campaign 1Campaign 1Campaign 1Campaign 1Campaign 1Campaign 1Campaign 1Campaign 1Campaign 1Campaign 1Campaign 1Campaign 1Campaign 1Campaign 1Campaign 1Campaign 1Campaign 1Campaign 1Campaign 1Campaign 1Campaign 1Campaign 1Campaign 1Campaign 1Campaign 1Campaign 1Campaign 1Campaign 1Campaign 1', 1000000, 0, 'satu, dua, tiga, empat', 0, 'campaign-satu', '2022-01-16 13:57:16', '2022-01-16 13:57:16'),
(2, 2, 'Campaign 2', 'Campaign 2Campaign 2Campaign 2Campaign 2Campaign 2', 'Campaign 2Campaign 2Campaign 2Campaign 2Campaign 2Campaign 2Campaign 2Campaign 2Campaign 2Campaign 2Campaign 2Campaign 2Campaign 2Campaign 2', 1500000, 0, 'satu, dua, tiga, empat, lima', 0, 'campaign-dua\r\n', '2022-01-16 13:57:16', '2022-01-16 13:57:16'),
(3, 1, 'Campaign 3', 'Campaign 3Campaign 3Campaign 3Campaign 3Campaign 3', 'Campaign 3Campaign 3Campaign 3Campaign 3Campaign 3Campaign 3Campaign 3Campaign 3Campaign 3Campaign 3Campaign 3Campaign 3Campaign 3Campaign 3Campaign 3Campaign 3Campaign 3Campaign 3Campaign 3Campaign 3Campaign 3', 2500000, 0, 'satu, dua, tiga, empat, lima, enam', 0, 'campaign-tiga\r\n', '2022-01-16 13:57:16', '2022-01-16 13:57:16'),
(4, 1, 'Campaign Test', 'ShortShortShortShortShort', 'LongLongLongLongLongLongLongLongLong', 100000000, 0, 'hadiah satu, hadiah dua, hadiah tiga', 0, 'campaign-test-s-int-1', '2022-01-17 23:25:31', '2022-01-17 23:25:31'),
(5, 1, 'Update data Sebuah campaign dari postman', 'short description', 'panjang kali lebar kali tinggi', 250000000, 0, 'keuntungan satu, keuntungan dua, keuntungan tiga, tambah empat', 0, 'campaign-test-1', '2022-01-17 23:26:30', '2022-01-18 08:47:37'),
(6, 2, 'Update data Sebuah campaign dari postman', 'short description', 'panjang kali lebar kali tinggi', 250000000, 0, 'keuntungan satu, keuntungan dua, keuntungan tiga, tambah empat', 0, 'campaign-test-2', '2022-01-17 23:27:16', '2022-01-18 09:03:16'),
(7, 2, 'Sebuah campaign dari postman', 'short description', 'panjang kali lebar kali tinggi', 150000000, 0, 'keuntungan satu, keuntungan dua, keuntungan tiga', 0, 'sebuah-campaign-dari-postman-2', '2022-01-17 23:46:09', '2022-01-17 23:46:09');

-- --------------------------------------------------------

--
-- Struktur dari tabel `campaign_images`
--

CREATE TABLE `campaign_images` (
  `id` int(11) NOT NULL,
  `campaign_id` int(11) NOT NULL,
  `file_name` varchar(255) NOT NULL,
  `is_primary` tinyint(4) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `campaign_images`
--

INSERT INTO `campaign_images` (`id`, `campaign_id`, `file_name`, `is_primary`, `created_at`, `updated_at`) VALUES
(1, 1, 'satu.jpg', 0, '2022-01-16 15:19:34', '2022-01-16 15:19:34'),
(2, 1, 'dua.jpg', 1, '2022-01-16 15:19:34', '2022-01-16 15:19:34'),
(3, 1, 'tiga.jpg', 0, '2022-01-16 15:21:09', '2022-01-16 15:21:09'),
(4, 6, 'campaign-images/2022-01-18-212826-2-repository-open-graph-template.png', 1, '2022-01-18 10:43:07', '2022-01-18 10:43:07'),
(5, 7, 'campaign-images/2-repository-open-graph-template.png', 0, '2022-01-18 11:32:01', '2022-01-18 21:36:03'),
(6, 7, 'campaign-images/2-repository-open-graph-template.png', 0, '2022-01-18 11:34:44', '2022-01-18 21:36:03'),
(7, 7, 'campaign-images/18181-118-18-00-00-2-repository-open-graph-template.png', 0, '2022-01-18 20:43:18', '2022-01-18 21:36:03'),
(8, 7, 'campaign-images/20220118204754-2-repository-open-graph-template.png', 0, '2022-01-18 20:47:54', '2022-01-18 21:36:03'),
(9, 7, 'campaign-images/2022-01-18-204906-2-repository-open-graph-template.png', 0, '2022-01-18 20:49:06', '2022-01-18 21:36:03'),
(10, 7, 'campaign-images/2022-01-18-212812-2-repository-open-graph-template.png', 0, '2022-01-18 21:28:12', '2022-01-18 21:36:03'),
(11, 7, 'campaign-images/2022-01-18-212826-2-repository-open-graph-template.png', 0, '2022-01-18 21:28:26', '2022-01-18 21:36:03'),
(12, 7, 'campaign-images/2022-01-18-213603-2-repository-open-graph-template.png', 1, '2022-01-18 21:36:03', '2022-01-18 21:36:03');

-- --------------------------------------------------------

--
-- Struktur dari tabel `transactions`
--

CREATE TABLE `transactions` (
  `id` int(11) NOT NULL,
  `campaign_id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `amount` int(11) NOT NULL,
  `status` varchar(255) NOT NULL,
  `code` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `transactions`
--

INSERT INTO `transactions` (`id`, `campaign_id`, `user_id`, `amount`, `status`, `code`, `created_at`, `updated_at`) VALUES
(3, 7, 1, 1500000, 'paid', '', '2022-01-18 15:57:40', '2022-01-18 15:57:40'),
(4, 6, 1, 1500000, 'paid', '', '2022-01-18 15:57:40', '2022-01-18 15:57:40');

-- --------------------------------------------------------

--
-- Struktur dari tabel `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `occupation` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password_hash` varchar(255) NOT NULL,
  `avatar_file_name` varchar(255) NOT NULL,
  `role` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `users`
--

INSERT INTO `users` (`id`, `name`, `occupation`, `email`, `password_hash`, `avatar_file_name`, `role`, `created_at`, `updated_at`) VALUES
(1, 'Andini Septian', 'Go Programmer', 'andini@gmail.com', '$2a$04$nQwsTMjho3t7KI3WH35jKehxmoQzgjgVS5nIU5qQD5hwF6NOuRX8a', 'images/1-repository-open-graph-template.png', 'user', '2022-01-13 03:05:27', '2022-01-16 15:21:29'),
(2, 'Agung Hermansyah', 'Frontend Developer', 'agung@gmail.com', '$2a$04$nQwsTMjho3t7KI3WH35jKehxmoQzgjgVS5nIU5qQD5hwF6NOuRX8a', 'images/2-download (1).jfif', 'user', '2022-01-13 03:05:27', '2022-01-16 15:19:41'),
(3, 'Hendrik', 'PHP Programmer', 'hendrik@gmail.com', 'sample', '', 'user', '2022-01-13 03:08:27', '2022-01-13 03:08:27');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `campaigns`
--
ALTER TABLE `campaigns`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `campaign_images`
--
ALTER TABLE `campaign_images`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `campaigns`
--
ALTER TABLE `campaigns`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT untuk tabel `campaign_images`
--
ALTER TABLE `campaign_images`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;

--
-- AUTO_INCREMENT untuk tabel `transactions`
--
ALTER TABLE `transactions`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT untuk tabel `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=16;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
