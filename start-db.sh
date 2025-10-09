# Configuration
COMPOSE_FILE="docker-compose.yml"  # Change this to your compose file path
PROJECT_NAME=""                     # Optional: set project name, leave empty to use directory name

# Color output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Build the docker compose command
COMPOSE_CMD="docker compose"
if [ -n "$COMPOSE_FILE" ]; then
    COMPOSE_CMD="$COMPOSE_CMD -f $COMPOSE_FILE"
fi
if [ -n "$PROJECT_NAME" ]; then
    COMPOSE_CMD="$COMPOSE_CMD -p $PROJECT_NAME"
fi

# Check if docker compose file exists
if [ ! -f "$COMPOSE_FILE" ]; then
    echo -e "${RED}Error: Docker Compose file '$COMPOSE_FILE' not found${NC}"
    exit 1
fi

# Check if any containers from this compose file are running
RUNNING_CONTAINERS=$($COMPOSE_CMD ps -q)

if [ -z "$RUNNING_CONTAINERS" ]; then
    echo -e "${YELLOW}No containers running. Starting Docker Compose...${NC}"
    $COMPOSE_CMD up -d
    
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}Docker Compose started successfully${NC}"
        $COMPOSE_CMD ps
    else
        echo -e "${RED}Failed to start Docker Compose${NC}"
        exit 1
    fi
else
    echo -e "${GREEN}Docker Compose is already running${NC}"
    $COMPOSE_CMD ps
fi