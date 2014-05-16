FROM debian
ADD . /src
RUN mkdir /context
RUN cat /src/foo | sed s/foo/bar/g > /context/foo
RUN cp /src/Dockerfile.next /context/Dockerfile
CMD tar cvf - -C /context . | cat
