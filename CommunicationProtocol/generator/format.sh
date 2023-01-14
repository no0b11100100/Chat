for file_name in ./gen/**/*.{cpp,h,hpp}; do
	if [ -f "$file_name" ]; then
		printf '%s\n' "$file_name"
		python3 format_cpp.py $file_name
        clang-format -i $file_name
	fi
done