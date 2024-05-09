@echo off

SET SERVER=RuoYi-Go

cd /d %~dp0
SET BLDIR=%CD%
if EXIST %SERVER%.exe del %SERVER%.exe

cd ..
SET ROOT=%CD%

echo Building %SERVER%
cd %ROOT%\cmd
go build -o %BLDIR%\%SERVER%.exe .

echo Build done
