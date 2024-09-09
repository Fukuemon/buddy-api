#!/bin/sh

# Postman collection file path
POSTMAN_COLLECTION_FILE=$1

# List of parameters to replace
PARAMETERS="facility_id position_id department_id team_id position_id"

# Loop through each parameter and replace :parameter with {{parameter}}
for PARAM in $PARAMETERS; do
  sed -i "s/:${PARAM}/{{${PARAM}}}/g" "$POSTMAN_COLLECTION_FILE"
done
