validate-number(){
  re='^[0-9]+$'
  if [[ ! $1 =~ $re ]]; then
    echo "Error: The value must be a number." >&2
    exit 1
