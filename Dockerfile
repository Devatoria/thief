FROM bash

COPY bin/thief /bin/thief

ENTRYPOINT ["/bin/thief"]
