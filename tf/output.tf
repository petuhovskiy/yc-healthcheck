output "internal_ip_address_vm_nat" {
  value = yandex_compute_instance.vm-nat.network_interface.0.ip_address
}

output "external_ip_address_vm_nat" {
  value = yandex_compute_instance.vm-nat.network_interface.0.nat_ip_address
}

output "internal_ip_address_vm_postgres" {
  value = yandex_compute_instance.vm-postgres.network_interface.0.ip_address
}

output "internal_ip_address_vm_app_1" {
  value = yandex_compute_instance.vm-app-1.network_interface.0.ip_address
}

output "internal_ip_address_vm_app_2" {
  value = yandex_compute_instance.vm-app-2.network_interface.0.ip_address
}

output "external_ip_address_vm_nginx" {
  value = yandex_compute_instance.vm-nginx.network_interface.0.nat_ip_address
}

output "internal_ip_address_vm_nginx" {
  value = yandex_compute_instance.vm-nginx.network_interface.0.ip_address
}
