run:
	go run main.go
# Сборка exe файла
package-windows:
	fyne package -os windows
mod-tidy:
	go mod tidy