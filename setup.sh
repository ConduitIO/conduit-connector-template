#!/bin/bash
if [ $# -eq 0 ]
then
    echo "Module name not provided."
fi

MODULE_NAME=$1
if ! [[ "$MODULE_NAME" =~ ^.*\/conduit-connector-(.*)$ ]]
then
  echo "Module name '$MODULE_NAME' not in required format 'example/conduit-connector-connectorName'"
  exit 1
fi

CONNECTOR_NAME=${BASH_REMATCH[1]}

if [[ "$OSTYPE" == "darwin"* ]]; then
  LC_ALL=C find . -type f ! -name "setup.sh" -exec sed -i "" "s~github.com/conduitio/conduit-connector-connectorName~$MODULE_NAME~g" {} +
  LC_ALL=C find . -type f ! -name "setup.sh" -exec sed -i "" "s~connectorName~$CONNECTOR_NAME~g" {} +
  LC_ALL=C sed -i "" "s~*       @ConduitIO/conduit-core~ ~g" .github/CODEOWNERS
else
  find . -type f ! -name "setup.sh" -exec sed -i "s~github.com/conduitio/conduit-connector-connectorName~$MODULE_NAME~g" {} +
  find . -type f ! -name "setup.sh" -exec sed -i "s~connectorName~$CONNECTOR_NAME~g" {} +
  sed -i "s~*       @ConduitIO/conduit-core~ ~g" .github/CODEOWNERS
fi

# Remove this script
rm "$0"
rm README.md
mv README_TEMPLATE.md README.md
