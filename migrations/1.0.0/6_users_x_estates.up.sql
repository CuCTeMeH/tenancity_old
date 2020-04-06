CREATE TABLE IF NOT EXISTS `users_x_estates` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  `estate_id` bigint(20) unsigned DEFAULT NULL,
  `rent_id` bigint(20) unsigned DEFAULT NULL,
  `is_tenant` tinyint(1) NOT NULL DEFAULT 0,
  `is_landlord` tinyint(1) NOT NULL DEFAULT 0,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `users_x_estates_user_id_foreign` (`user_id`),
  KEY `users_x_estates_estate_id_foreign` (`estate_id`),
  KEY `users_x_rents_estate_id_foreign` (`rent_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;