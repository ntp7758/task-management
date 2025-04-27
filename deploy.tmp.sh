# Stop old service
docker compose -f docker-compose.tmp.yml down

# Start new service
docker compose -f docker-compose.tmp.yml up --build -d