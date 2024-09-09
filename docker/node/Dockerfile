FROM node:14-alpine

# Install necessary packages
RUN apk add --no-cache curl jq

# Install openapi-to-postman tool
RUN npm install -g openapi-to-postmanv2

# Set working directory
WORKDIR /app

# Copy the Swagger file and the replace script into the container
COPY ./internal/docs /app/docs
COPY ./script/replace_parameters.sh /app/replace_parameters.sh

# Make the replace script executable
RUN chmod +x /app/replace_parameters.sh

# Default command
CMD ["sh", "-c", "openapi2postmanv2 -s /app/docs/swagger.json -o /app/docs/postman_collection.json -p && /app/replace_parameters.sh /app/docs/postman_collection.json && jq '{collection: .}' /app/docs/postman_collection.json > /app/docs/postman_collection_wrapped.json && curl --location --request PUT https://api.getpostman.com/collections/${POSTMAN_COLLECTION_ID} --header \"X-Api-Key: ${POSTMAN_API_KEY}\" --header \"Content-Type: application/json\" --data @/app/docs/postman_collection_wrapped.json"]
