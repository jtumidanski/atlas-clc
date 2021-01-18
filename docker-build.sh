if [[ "$1" = "NO-CACHE" ]]
then
   docker build -f Dockerfile.dev --no-cache --tag atlas-clc:latest .
else
   docker build -f Dockerfile.dev --tag atlas-clc:latest .
fi
