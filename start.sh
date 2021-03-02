#!/bin/bash

echo "===== BUILDING CLIENT ====="
cd client && npm run build && cd ../

echo "===== START SERVER ====="
go run backend/server/main.go

echo "===== DONE! ====="  