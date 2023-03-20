docker-compose down &&
 docker-compose rm &&
 docker-compose pull &&
 docker-compose build --no-cache &&
 docker-compose up -d --force-recreate
 sleep 10
docker exec myarticles_db  mongo myarticles ./create-user.js -u rootadmin -p rootadminpassword --authenticationDatabase admin