CREATE TABLE `estate_paid_bills` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `bill_id` bigint(20) unsigned DEFAULT NULL,
  `estate_id` bigint(20) unsigned DEFAULT NULL,
  `rent_id` bigint(20) unsigned DEFAULT NULL,
  `payer_id` bigint(20) unsigned DEFAULT NULL,
  `amount` decimal(50,4) NOT NULL DEFAULT 0.0000,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `estate_paid_bills_bill_id_foreign` (`bill_id`),
  KEY `estate_paid_bills_estate_id_foreign` (`estate_id`),
  KEY `estate_paid_bills_rent_id_foreign` (`rent_id`),
  KEY `estate_paid_bills_payer_id_foreign` (`payer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;