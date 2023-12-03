#!/usr/bin/env bash

# Script taken from
# https://github.com/Stogas/aoc2023-go/blob/main/create-day.sh

if [[ ! -n $1 ]]; then
  echo "No day num provided" >&2
	exit 2
fi

dayDir="day${1}"

if [[ ! -d "${dayDir}" ]]; then
	mkdir "${dayDir}"
  cp template.go "${dayDir}/day${1}.go"
	touch "${dayDir}/in"
else
	echo "Day directory already exists" >&2
	exit 3
fi