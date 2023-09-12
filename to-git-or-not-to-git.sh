curl -s https://01.alem.school/assets/superhero/all.json | jq -r " .[] | select( .id == 170).name"
curl -s https://01.alem.school/assets/superhero/all.json | jq -r " .[] | select( .id == 170).powerstats.power"
curl -s https://01.alem.school/assets/superhero/all.json | jq -r " .[] | select( .id == 170).appearance.gender"

