FROM dpokidov/imagemagick AS BASE

RUN apt update && apt install python3 -y
