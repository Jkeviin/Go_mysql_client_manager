/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */
;

/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */
;

/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */
;

/*!40101 SET NAMES utf8mb4 */
;

--
-- Base de datos: `golang_consola`
--
-- --------------------------------------------------------
--
-- Estructura de tabla para la tabla `clientes`
--
CREATE TABLE `clients` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `email` varchar(100) NOT NULL,
  `phone` varchar(20) NOT NULL,
  `active` boolean NOT NULL DEFAULT TRUE,
  `date` datetime NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO `clients` (`name`, `email`, `phone`, `active`, `date`) VALUES ('John Doe', 'john@example.com', '123456789', TRUE, NOW());

INSERT INTO `clients` (`name`, `email`, `phone`, `active`, `date`) VALUES ('Jane Smith', 'jane@example.com', '987654321', TRUE, NOW());

INSERT INTO `clients` (`name`, `email`, `phone`, `active`, `date`) VALUES ('Alice Johnson', 'alice@example.com', '555555555', TRUE, NOW());
