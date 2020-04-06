CREATE TABLE `estate_bills` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `type` enum('bill','rent') COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'bill',
  `description` longtext COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `recurring` tinyint(1) NOT NULL DEFAULT 0,
  `amount` decimal(50,4) NOT NULL DEFAULT 0.0000,
  `estate_id` bigint(20) unsigned DEFAULT NULL,
  `rent_id` bigint(20) unsigned DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `estate_bills_estate_id_foreign` (`estate_id`),
  KEY `estate_bills_rent_id_foreign` (`rent_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;