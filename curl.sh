if [ $# -ne 2 ]; then
    echo "Usage: $0 <IP_ADDRESS> <PORT>"
    exit 1
fi

GW_IP="$1"
GW_PORT="$2"
BASE_URL="http://server.example.com:$GW_PORT"

output_file="curl_output.txt" > "$output_file"

for ((i=1; i<=100; i++)); do    
    REQUEST_URL="$BASE_URL"
    echo "server.example.com:$GW_PORT:$GW_IP" "$REQUEST_URL"
    curl --resolve "server.example.com:$GW_PORT:$GW_IP" "$REQUEST_URL/"  >> "$output_file"
done

echo "Finished sending 200 random requests."

count_v1=$(cat "$output_file" | grep hi | wc -l)
count_v2=$(cat "$output_file" | grep hello | wc -l)

echo "Number of 'v1' occurrences: $count_v1"
echo "Number of 'v2' occurrences: $count_v2"
rm curl_output.txt