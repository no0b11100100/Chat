import os
import sys

def format():
    formatted_data = ""
    if len(sys.argv) != 2:
        return

    path = sys.argv[1]
    with open(path, "r") as f:
        for line in f:
            line = line.replace(",)", ")")
            formatted_data += line.replace(",}", "}")

    with open(path, "w") as f:
        f.write(formatted_data)


if __name__ == '__main__':
    format()
