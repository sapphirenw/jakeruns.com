import requests
import re
import os
import json
import sys


def parse_md(content):
    # parse the rest of the data
    title_pattern = r"^#\s+(.+)"
    description_pattern = r"## Description\s+(.*?)(?=\n#|$)"

    title_match = re.search(title_pattern, content, re.MULTILINE)
    description_match = re.search(description_pattern, content, re.DOTALL)

    title = title_match.group(1).strip() if title_match else None
    description = description_match.group(1).strip() if description_match else None

    return {
        "title": title,
        "description": description,
        "content": content,
    }


if __name__ == "__main__":
    # parse args
    if len(sys.argv) < 2:
        print("Supply a file to upload")
        exit(1)

    file_path = f"{os.curdir}/{sys.argv[-1]}"

    # read the file
    with open(file_path, "r") as file:
        content = file.read()
        file.close()

    body = {}
    update = False

    # parse the file
    if "---" in content:
        metadata, md = content.split("---")
        metadata = json.loads(metadata)
        if "articleId" in metadata:
            update = True

        # combine fields
        body = {**body, **metadata}
    else:
        md = content

    # parse the markdown
    data = parse_md(md)
    body = {**body, **data}

    # send the request
    response = requests.post(
        "http://localhost:3000/api/articles",
        headers={
            "x-api-key": os.getenv("API_KEY"),
            "Content-Type": "application/json",
        },
        data=json.dumps(body),
    )

    if response.status_code != 200:
        print(f"issue with file: {file_path}")
        print(response.content)
        print("")
    else:
        print(f"successfully uploaded: {file_path}")

        if not update:
            with open(file_path, "w") as file:
                # alter the file to include the generated articleId
                article = json.loads(response.content)
                header_data = {"articleId": article["articleId"]}
                new_content = f"{json.dumps(header_data)}\n---\n{content}"
                file.write(new_content)
                file.close()
