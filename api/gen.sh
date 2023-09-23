#!/usr/bin/env bash

echo -e "$(printf "1 -> Generate entgo.io model")"
echo -e "$(printf "2 -> Generate entgo.io query")"
echo -n "Select your command: "

read -r CMD

case $CMD in
1)
  echo -n "Your model name: "
  read -r NAME
  if [[ $NAME::1 =~ [[:upper:]] ]]; then
    echo -e "======================================="
    echo -e "$(printf "Generating model with name: %s" "$NAME")"

    go run -mod=mod entgo.io/ent/cmd/ent new "$NAME"

    echo -e "======================================="
    echo "Done!!! Please update your model and generate query then"
  else
    echo -e "$(printf "ERROR: Model must start with uppercase character. E.g., DuyNguyen")"
  fi
  ;;
2)
  echo -e "======================================="
  echo -e "$(printf "Generating query")"

  go generate ./ent

  echo -e "======================================="
  echo -e "Done!!! Query generation completed successfully"
  ;;
*)
  echo -e "$(printf "ERROR: Unknown command: %s" "${CMD}")"
  ;;

esac
