@echo off

SET SERVER=time-machine

cd /d %~dp0
SET BLDIR=%CD%
if EXIST %SERVER%.exe del %SERVER%.exe

cd ..
SET ROOT=%CD%

echo Building %SERVER%
cd %ROOT%\cmd\backend
go build -o %BLDIR%\%SERVER%.exe .

echo Build done
