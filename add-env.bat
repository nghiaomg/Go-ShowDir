@echo off
setlocal

REM Get the current directory
set "CURRENT_DIR=%~dp0"
if "%CURRENT_DIR:~-1%"=="\" set "CURRENT_DIR=%CURRENT_DIR:~0,-1%"

echo Adding %CURRENT_DIR% to PATH...

echo %PATH% | find /I "%CURRENT_DIR%" >nul
if %errorlevel%==0 (
    echo Already exists in PATH
) else (
    setx PATH "%PATH%;%CURRENT_DIR%"
    echo Added successfully!
)

echo.
echo >>> Close this window and open a NEW CMD/PowerShell to use "showdir"
pause