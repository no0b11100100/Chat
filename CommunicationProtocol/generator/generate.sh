for file_name in ./qface/*.qface; do
	if [ -f "$file_name" ]; then
		printf '##%s\n' "$file_name"
        python3 codegen.py --input $file_name --output ./gen/
	fi
done
