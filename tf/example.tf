terraform {
  required_providers {
    yandex = {
      source = "yandex-cloud/yandex"
      version = "0.49.0"
    }
  }
}

provider "yandex" {
  zone      = "ru-central1-a"
}

data "yandex_compute_image" "container-optimized-image" {
  family = "container-optimized-image"
}
