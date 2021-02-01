USMDEX_ENABLE_TLS=true; \
USMDEX_MONGODB_SHARD_URL_1=udexcluster0-shard-00-00-xzynf.mongodb.net:27017; \
USMDEX_MONGODB_SHARD_URL_2=udexcluster0-shard-00-01-xzynf.mongodb.net:27017; \
USMDEX_MONGODB_SHARD_URL_3=udexcluster0-shard-00-02-xzynf.mongodb.net:27017; \
fresh

# go run --race main.go