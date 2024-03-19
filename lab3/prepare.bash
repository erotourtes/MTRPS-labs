#!/bin/bash

# Check if the number of arguments passed is exactly one
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <number>"
    exit 1
fi

file=starter-project-python/requirements/backend.in
api_file=starter-project-python/spaceship/routers/api.py

if [ "$1" -lt 5 ]; then
cat > $file << EOF
fastapi==0.110.0
pydantic==1.10.0
starlette==0.36.3
uvicorn==0.28.0
## The following requirements were added by pip freeze:
annotated-types==0.6.0
anyio==4.3.0
click==8.1.7
h11==0.14.0
httptools==0.6.1
idna==3.6
pydantic_core==2.14.1
python-dotenv==1.0.1
PyYAML==6.0.1
sniffio==1.3.1
typing_extensions==4.10.0
uvloop==0.19.0
watchfiles==0.21.0
websockets==12.0
EOF

elif [ "$1" -eq 5 ]; then
cat > $file << EOF
annotated-types==0.6.0
anyio==4.3.0
click==8.1.7
fastapi==0.110.0
h11==0.14.0
httptools==0.6.1
idna==3.6
numpy==1.26.4
pydantic==1.10.0
pydantic_core==2.14.1
python-dotenv==1.0.1
PyYAML==6.0.1
sniffio==1.3.1
starlette==0.36.3
typing_extensions==4.10.0
uvicorn==0.28.0
uvloop==0.19.0
watchfiles==0.21.0
websockets==12.0
EOF

cat > $api_file << EOF
from fastapi import APIRouter
from numpy import random

router = APIRouter()


@router.get('')
def hello_world() -> dict:
    return {'msg': 'Hello, World!'}

@router.get('/matrix')
def matrix() -> dict:
    matrix_a = random.randint(1, 100, (10, 10))
    matrix_b = random.randint(1, 100, (10, 10))
    product = matrix_a * matrix_b
    
    return {
        'matrix_a': matrix_a.tolist(),
        'matrix_b': matrix_b.tolist(),
        'product': product.tolist()
    }
EOF
fi

#pip install -r $file --upgrade
