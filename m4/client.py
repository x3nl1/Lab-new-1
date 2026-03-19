import subprocess
import json

def run(numbers):
    proc = subprocess.Popen(
        ["go", "run", "main.go"],
        stdin=subprocess.PIPE,
        stdout=subprocess.PIPE,
        cwd="m4"
    )
    data = json.dumps({"numbers": numbers}).encode()
    out, _ = proc.communicate(data)
    return json.loads(out)["sum"]