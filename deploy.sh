# Stop old service
docker compose --env-file project.env down

# Start new service
docker compose --env-file project.env up --build -d