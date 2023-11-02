@echo off

set folder=%1
set exeFile=%2
rmdir %folder% /s /q

if not exist %folder% (
    mkdir %folder%
)

move %exeFile% .\%folder%\%exeFile%
xcopy .\shellScripts .\%folder%\shellScripts\ /yq