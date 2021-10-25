-- phpMyAdmin SQL Dump
-- version 4.9.5deb2
-- https://www.phpmyadmin.net/
--
-- Хост: localhost:3306
-- Время создания: Окт 25 2021 г., 20:01
-- Версия сервера: 8.0.26-0ubuntu0.20.04.3
-- Версия PHP: 7.4.3

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- База данных: `sosedi`
--

-- --------------------------------------------------------

--
-- Структура таблицы `liked_users`
--

CREATE TABLE `liked_users` (
  `user_id` int UNSIGNED NOT NULL,
  `liked_user_id` int UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

--
-- Дамп данных таблицы `liked_users`
--

INSERT INTO `liked_users` (`user_id`, `liked_user_id`) VALUES
(49, 19),
(49, 27),
(49, 38);

-- --------------------------------------------------------

--
-- Структура таблицы `users`
--

CREATE TABLE `users` (
  `id` int UNSIGNED NOT NULL,
  `name` varchar(50) NOT NULL,
  `surname` varchar(50) NOT NULL,
  `birthday` date NOT NULL,
  `price_min` int UNSIGNED NOT NULL,
  `price_max` int UNSIGNED NOT NULL,
  `password` varchar(100) NOT NULL,
  `sault` varchar(10) NOT NULL,
  `email` varchar(50) NOT NULL,
  `sex` varchar(1) NOT NULL,
  `photo_key` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

--
-- Дамп данных таблицы `users`
--

INSERT INTO `users` (`id`, `name`, `surname`, `birthday`, `price_min`, `price_max`, `password`, `sault`, `email`, `sex`, `photo_key`) VALUES
(19, 'Нина', 'ниночка', '2000-05-05', 34000, 23000, '$2a$05$SLUFmSuykVSnzVot7GjFLurB8.GOhPcEnIzCkWiknVDzYmbE5O6B.', 'ticJgdiEk2', 'gendina.nina@mail.ru', 'w', ''),
(26, 'Нина', 'ниночка', '2001-06-17', 34000, 23000, '$2a$05$by3swHnpOnx20DuW1HQqgeP8Z2f5PosoyNQbaB7E/IGmvfjhgg152', 'dTUbtERbH2', 'df@mail.ru', 'm', ''),
(27, 'werewrewre', 'wearwewae', '2001-06-17', 34000, 23000, '$2a$05$3C1gSMGEwZSI3aMwvqRke..LnT5m2biS4qjtvQNQamv3Xy3EJJ9O2', 'Bb1LoknCn7', 'dfeef', 'm', ''),
(31, 'fff', 'fff', '2001-06-17', 34000, 23000, '$2a$05$QP7EwaQ4Oly2ngekhTJ/iOamWBJQPn/kdselAtu3loMK6cMbxoz9i', '5B2Kjw7skk', 'rr@mail.ru', 'm', ''),
(35, 'd', 'ниночка', '2001-06-17', 34000, 23000, '$2a$05$TyX8Gh.AB92kPWdjgVyHxeFC/2tanZH1btauw4h0OCKcDxa./WeQi', 'cK2DfeGJrt', 'nina@mail.ru', 'm', ''),
(36, '2wqwqw', 'qwqwqw', '2001-06-17', 34000, 23000, '$2a$05$vTs3JHxma4g7iHvQlJhLnOMeHa7T2in7zjTFBTQiqzqj3I4V64TYi', '1DurM36Bsp', 'new@mail.ru', 'm', ''),
(37, 'Нина', 'ниночка', '2001-06-17', 34000, 23000, '$2a$05$v8I61ib7sCjYWfdibW0Bm.L17VBp5VZAlkxhOaGKm9SBAPLnfuOXm', '0FNv8rL83p', 'gendina.ninaee@mail.ru', 'm', '614f7d112e9e9f7110f25cb9.jpg'),
(38, 'Нина', 'ниночка', '2001-06-17', 34000, 23000, '$2a$05$LCLWVDx5IHP0.vZcpsuvjenU5k.AWI9mUUky8QNyyH7Jfty/Ss9Jm', 'qkAmj8d0T2', 'gendinas.nina@mail.ru', 'm', '614f8b832e9e9f7972de7d1e.jpg'),
(39, 'Нина', 'ниночка', '2001-06-17', 34000, 23000, '$2a$05$tMa0IMT8dMS0Z4p3Ajqese6OHUHpgVv9/uJc4ZpYmS3JMjcrmm1Ma', 'paAcbmHB0s', 'gendinfffa.nina@mail.ru', 'm', '614fa07d2e9e9f844c05567e.jpeg'),
(40, 'нина', 'САМАЯВРЕДНАЯНАСВЕТЕ', '1987-07-16', 34000, 23000, '$2a$05$/tkTi./cOG00ujoCgbPNyu.xlH8ALEqxk7xB2fmJFnULfSk20.VDa', 'XAAgACeAFW', 'holl.er@mail.com', 'w', '61508a2df5356bca7da78fe2.png'),
(42, 'нина', 'САМАЯВРЕДНАЯНАСВЕТЕ', '1987-07-16', 34000, 23000, '$2a$05$/tkTi./cOG00ujoCgbPNyu.xlH8ALEqxk7xB2fmJFnULfSk20.VDa', 'XAAgACeAFW', 'holll.er@mail.com', 'w', '61508a2df5356bca7da78fe2.png'),
(43, 'нина', 'САМАЯВРЕДНАЯНАСВЕТЕ', '1987-07-16', 34000, 23000, '$2a$05$/tkTi./cOG00ujoCgbPNyu.xlH8ALEqxk7xB2fmJFnULfSk20.VDa', 'XAAgACeAFW', 'hollll.er@mail.com', 'w', '61508a2df5356bca7da78fe2.png'),
(44, 'нина', 'САМАЯВРЕДНАЯНАСВЕТЕ', '1987-07-16', 34000, 23000, '$2a$05$/tkTi./cOG00ujoCgbPNyu.xlH8ALEqxk7xB2fmJFnULfSk20.VDa', 'XAAgACeAFW', 'holllll.er@mail.com', 'w', '61508a2df5356bca7da78fe2.png'),
(45, 'нина', 'САМАЯВРЕДНАЯНАСВЕТЕ', '1987-07-16', 34000, 23000, '$2a$05$/tkTi./cOG00ujoCgbPNyu.xlH8ALEqxk7xB2fmJFnULfSk20.VDa', 'XAAgACeAFW', 'holllll.eer@mail.com', 'w', '61508a2df5356bca7da78fe2.png'),
(46, 'нина', 'САМАЯВРЕДНАЯНАСВЕТЕ', '1987-07-16', 34000, 23000, '$2a$05$/tkTi./cOG00ujoCgbPNyu.xlH8ALEqxk7xB2fmJFnULfSk20.VDa', 'XAAgACeAFW', 'holllll.eeer@mail.com', 'w', '61508a2df5356bca7da78fe2.png'),
(47, 'нина', 'САМАЯВРЕДНАЯНАСВЕТЕ', '1987-07-16', 34000, 23000, '$2a$05$/tkTi./cOG00ujoCgbPNyu.xlH8ALEqxk7xB2fmJFnULfSk20.VDa', 'XAAgACeAFW', 'holllll.eeeer@mail.com', 'w', '61508a2df5356bca7da78fe2.png'),
(48, 'нина', 'САМАЯВРЕДНАЯНАСВЕТЕ', '1987-07-16', 34000, 23000, '$2a$05$/tkTi./cOG00ujoCgbPNyu.xlH8ALEqxk7xB2fmJFnULfSk20.VDa', 'XAAgACeAFW', 'hoolllll.eeeer@mail.com', 'w', '61508a2df5356bca7da78fe2.png'),
(49, 'Roman', 'Bakanov', '2001-03-05', 34000, 23000, '$2a$05$ZD978xx0n7IFa99DK0sVwem81aSyOHz4kfj.4CsqWw0S/VqxuiGSS', 'OPC6RcMLrs', 'roma@mail.ru', 'm', '6172cc54f5356b1b9fdeb868.png');

--
-- Индексы сохранённых таблиц
--

--
-- Индексы таблицы `liked_users`
--
ALTER TABLE `liked_users`
  ADD PRIMARY KEY (`user_id`,`liked_user_id`),
  ADD KEY `FK_liked_user_id_id` (`liked_user_id`);

--
-- Индексы таблицы `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`);

--
-- AUTO_INCREMENT для сохранённых таблиц
--

--
-- AUTO_INCREMENT для таблицы `users`
--
ALTER TABLE `users`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=50;

--
-- Ограничения внешнего ключа сохраненных таблиц
--

--
-- Ограничения внешнего ключа таблицы `liked_users`
--
ALTER TABLE `liked_users`
  ADD CONSTRAINT `FK_liked_user_id_id` FOREIGN KEY (`liked_user_id`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `FK_user_id_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
