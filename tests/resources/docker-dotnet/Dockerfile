## This dockerfile is just a test resource and is not buildable
ARG runtime_version=2.0.6
FROM microsoft/aspnetcore:${runtime_version} AS base
WORKDIR /app
EXPOSE 80

FROM base AS secondary
RUN nothing

ARG build_image
ARG build_version=2.0
FROM microsoft/$build_image:$build_version AS builder
WORKDIR /src
COPY *.sln ./
COPY Web/Web.csproj Web/
RUN dotnet restore
COPY . .
WORKDIR /src/Web
RUN dotnet build -c Release -o /app

FROM builder AS publish
RUN dotnet publish -c Release -o /app

FROM imaginary/cert-generator:1.0
RUN generate-cert.sh

FROM secondary AS production
WORKDIR /app
COPY --from=publish /app .
COPY --from=3 /cert /app
ENTRYPOINT ["dotnet", "Web.dll"]
