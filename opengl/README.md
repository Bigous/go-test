# Necessário:
1. Instale o GCC (caso já não esteja instalado)
	- No windows 10:
		1. Abrir o powershell em adm
		- `Set-ExecutionPolicy Unrestricted`
		- `Install-Package mingw`
			- Se este passo falhar após o download:
				- Descompactar o arquivo gerado em %TMP%\chocolatey\mingw\mingwinstall.zip para a pasta C:\Chocolatey\bin
				- Adicionar C:\Chocolatey\bin\mingw\bin ao path
		- Congratulations :D
	- No Windows anterior ao 10:
		1. Instale o chocolatey (caso já não esteja instalado)
			1. Abra o prompt de comando como administrador
			- `@powershell -NoProfile -ExecutionPolicy Bypass -Command "iex ((new-object net.webclient).DownloadString('https://chocolatey.org/install.ps1'))" && SET PATH=%PATH%;%ALLUSERSPROFILE%\chocolatey\bin`
		- Instale o mingw
			1. Abra o prompt de comando como administrador
			- `choco install mingw`
- Instale ou atualize a versão do GO para 1.6 (ou superior - por causa da versão do cgo necessário para linkar go com dlls em c)
	- Nome do pacote no Chocolatey: `golang`
