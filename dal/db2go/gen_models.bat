@echo off


set OUT_DIR=..
set PACK_NAME=models
set SUFFIX_NAME="do"
set READ_ONLY="created_time,updated_time"
set DB_NAME="coin"
set WITH_OUT=""
set TAGS="bson"
set DSN_URL="mysql://root:123456@127.0.0.1:3306/coin?charset=utf8"
set JSON_PROPERTIES="omitempty"
set TINYINT_TO_BOOL="deleted,disabled,banned"
set TABLE_NAME=""

db2go --url %DSN_URL% --out %OUT_DIR% --db %DB_NAME% --table %TABLE_NAME% --enable-decimal --json-properties %JSON_PROPERTIES% ^
--suffix %SUFFIX_NAME% --package %PACK_NAME% --readonly %READ_ONLY% --without %WITH_OUT% --tag %TAGS% --tinyint-as-bool %TINYINT_TO_BOOL%


If "%errorlevel%" == "0" (
echo generate go file ok, formatting...
gofmt -w %OUT_DIR%/%PACK_NAME%
) else (
echo if there is no db2go.exe, please download from https://github.com/civet148/release/tree/master/db2go
)

pause