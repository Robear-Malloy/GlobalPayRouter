#!/usr/bin/env bash
# scripts/healthcheck.sh
# Run after `docker compose up` to confirm every service is alive.
# Usage: ./scripts/healthcheck.sh

set -euo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

SERVICES=(
  "api-gateway        http://localhost:8080/health"
  "routing-engine     http://localhost:8081/health"
  "fraud-detection    http://localhost:8082/health"
  "compliance         http://localhost:8083/health"
  "fx-rate            http://localhost:8084/health"
  "transaction-ledger http://localhost:8085/health"
  "notification       http://localhost:8086/health"
  "rail-visa          http://localhost:8091/health"
  "rail-swift         http://localhost:8092/health"
  "rail-alipay        http://localhost:8093/health"
  "rail-wechat        http://localhost:8094/health"
  "rail-jcb           http://localhost:8095/health"
)

ALL_OK=true

echo ""
echo "GlobalRoute Pay — service health check"
echo "────────────────────────────────────────"

for entry in "${SERVICES[@]}"; do
  name=$(echo "$entry" | awk '{print $1}')
  url=$(echo "$entry"  | awk '{print $2}')

  if curl -sf "$url" > /dev/null 2>&1; then
    echo -e "  ${GREEN}✓${NC}  $name"
  else
    echo -e "  ${RED}✗${NC}  $name  ($url unreachable)"
    ALL_OK=false
  fi
done

echo ""
if $ALL_OK; then
  echo -e "${GREEN}All services healthy.${NC}"
  sleep 10
else
  echo -e "${RED}One or more services failed — check docker compose logs.${NC}"
  sleep 10
  exit 1
fi
