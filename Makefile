run:
	go run main.go
# Сборка exe файла
package-windows:
	fyne package -os windows
mod-tidy:
	go mod tidy
# Сборка проекта в папку build
build:
	call build.bat "build" "Запуск скриптов.exe"
build-windows: package-windows build