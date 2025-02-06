import subprocess 

output = subprocess.run(['git', '--version'], capture_output=True, text=True)
print(output.stdout)
print("Hello world!!!")