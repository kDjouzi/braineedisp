-- phpMyAdmin SQL Dump
-- version 4.6.4
-- https://www.phpmyadmin.net/
--
-- Client :  127.0.0.1
-- Généré le :  Dim 08 Octobre 2017 à 11:33
-- Version du serveur :  5.7.14
-- Version de PHP :  5.6.25

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de données :  `brain`
--

-- --------------------------------------------------------

--
-- Structure de la table `brainee`
--

CREATE TABLE `brainee` (
  `id` int(11) NOT NULL,
  `author` text NOT NULL,
  `text` text NOT NULL,
  `brand` text NOT NULL
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

--
-- Contenu de la table `brainee`
--

INSERT INTO `brainee` (`id`, `author`, `text`, `brand`) VALUES
(1, 'Khémys Djouzi', 'Plus de Téraoctets ! Toujours plus !', 'Western Digital'),
(2, 'Jon', 'Des toits plus colorés sur vos voitures, pour une ville plus joyeuse.', 'TaxiFou'),
(3, 'Mark', 'Du kérosène pour ma fusée de ville. Aussi, plus de points de recharge pour les voitures électriques.', 'Total');

--
-- Index pour les tables exportées
--

--
-- Index pour la table `brainee`
--
ALTER TABLE `brainee`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT pour les tables exportées
--

--
-- AUTO_INCREMENT pour la table `brainee`
--
ALTER TABLE `brainee`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
