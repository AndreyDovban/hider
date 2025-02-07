import subprocess
import os

output = subprocess.run(['git', '--version'], capture_output=True, text=True)
print(output.stdout)
print("Hello world!!!")
print(os.curdir)

try:
    file = open('myfile.dat', 'r+')
    print("task 1")
except IOError:
    file = open('myfile.dat', 'w+')
    print("task 2")
