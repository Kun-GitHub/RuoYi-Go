@echo off

SET SERVER=RuoYi-Go

cd /d %~dp0
SET BLDIR=%CD%
if EXIST %SERVER%.exe del %SERVER%.exe

cd ..
SET ROOT=%CD%

xcopy "%ROOT%\config" "%BLDIR%\config" /E /I /Q /Y
REM 删除指定文件
echo Removing app_config.go from build directory...
if EXIST "%BLDIR%\config\app_config.go" (
    del "%BLDIR%\config\app_config.go"
) else (
    echo File file_to_remove.txt does not exist.
)

echo Building %SERVER%
cd %ROOT%\cmd\api
go build -o %BLDIR%\%SERVER%.exe .

echo Build done
