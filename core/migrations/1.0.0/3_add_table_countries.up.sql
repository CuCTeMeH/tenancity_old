CREATE TABLE `countries` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `fips_104` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `iso2` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `iso3` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `ison` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `capital` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `nationality_singular` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `nationality_plural` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `languages` text COLLATE utf8_unicode_ci DEFAULT NULL,
  `default` tinyint(1) NOT NULL DEFAULT 0,
  `sort_order` bigint(20) NOT NULL DEFAULT 0,
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=276 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;