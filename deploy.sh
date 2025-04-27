# Stop old service
docker compose -f docker-compose.tmp.yml --env-file project.env down

# Start new service
docker compose -f docker-compose.tmp.yml --env-file project.env up --build -d