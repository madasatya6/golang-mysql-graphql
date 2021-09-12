-- phpMyAdmin SQL Dump
-- version 5.0.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 12 Sep 2021 pada 02.16
-- Versi server: 10.4.11-MariaDB
-- Versi PHP: 7.4.2

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `golang-mysql-graphql`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `products`
--

CREATE TABLE `products` (
  `ID_PRO` varchar(16) NOT NULL,
  `PRO_NAME` varchar(50) DEFAULT NULL,
  `QTE_IN_STOCK` int(11) DEFAULT NULL,
  `PRO_IMAGE` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `products`
--

INSERT INTO `products` (`ID_PRO`, `PRO_NAME`, `QTE_IN_STOCK`, `PRO_IMAGE`) VALUES
('lagu', 'ketan', 16, 'aslkfl'),
('pisang12', 'pisang bakar', 12, 'balisojdoeidjdl'),
('pisang13', 'sebakul nasi', 14, 'aslkdlfl');

-- --------------------------------------------------------

--
-- Struktur dari tabel `products_attribute`
--

CREATE TABLE `products_attribute` (
  `id` int(220) NOT NULL,
  `ID_PRO` varchar(225) DEFAULT NULL,
  `color` varchar(220) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `products_attribute`
--

INSERT INTO `products_attribute` (`id`, `ID_PRO`, `color`) VALUES
(1, 'pisang12', 'green'),
(2, 'pisang13', 'RED'),
(3, 'pisang12', 'blue');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`ID_PRO`);

--
-- Indeks untuk tabel `products_attribute`
--
ALTER TABLE `products_attribute`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `products_attribute`
--
ALTER TABLE `products_attribute`
  MODIFY `id` int(220) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
