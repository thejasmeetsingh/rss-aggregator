USERNAME=$1
PASSWORD=$2
DB_NAME=$3
COMMAND=$4

cd sql/schema
goose postgres postgres://$USERNAME:$PASSWORD@db/$DB_NAME $COMMAND