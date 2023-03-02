#!/bin/bash
if [ $# -eq 0 ]
then
    echo "Module name not provided."
    exit 1
fi

MODULE_NAME=$1
if ! [[ "$MODULE_NAME" =~ ^github.com\/.*\/conduit-connector-(.*)$ ]]
then
  echo "Module name ${MODULE_NAME} not in recommended format \"github.com/repository/conduit-connector-connectorname\"."
  echo
  echo "Certain things (such as pull request templates) will not work correctly."
  while true; do
      read -n1 -p "Are you sure you want to continue? [y/n] " yn
      echo
      case $yn in
          [Yy]* ) break;;
          [Nn]* ) exit;;
          * ) echo "Please answer yes or no.";;
      esac
  done
fi

CONNECTOR_NAME=${BASH_REMATCH[1]}

if [[ "$OSTYPE" == "darwin"* ]]; then
  LC_ALL=C find . -type f ! -name "setup.sh" -exec sed -i "" "s~github.com/conduitio/conduit-connector-connectorname~$MODULE_NAME~g" {} +
  LC_ALL=C find . -type f ! -name "setup.sh" -exec sed -i "" "s~connectorname~$CONNECTOR_NAME~g" {} +
  LC_ALL=C sed -i "" "s~*       @ConduitIO/conduit-core~ ~g" .github/CODEOWNERS
else
  find . -type f ! -name "setup.sh" -exec sed -i "s~github.com/conduitio/conduit-connector-connectorname~$MODULE_NAME~g" {} +
  find . -type f ! -name "setup.sh" -exec sed -i "s~connectorname~$CONNECTOR_NAME~g" {} +
  sed -i "s~*       @ConduitIO/conduit-core~ ~g" .github/CODEOWNERS
fi

# Remove this script
rm "$0"
rm README.md
mv README_TEMPLATE.md README.md
