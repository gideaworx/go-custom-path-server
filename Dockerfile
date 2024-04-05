ARG build_image=chainguard/go:latest
ARG run_image=chainguard/glibc-dynamic:latest

FROM ${build_image} AS build

USER nonroot
WORKDIR /build

COPY --chown=nonroot:nonroot . .
RUN ["go", "mod", "download"]
RUN ["go", "build"]

FROM ${run_image} AS run

USER nonroot
WORKDIR /app

COPY --chown=nonroot:nonroot --from=build /build/go-custom-path-server .
CMD [ "/app/go-custom-path-server" ]