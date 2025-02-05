ARG BASE_IMAGE={APPLICATION_NAME}:{DOCKER_IMAGE_TAG}

FROM ${BASE_IMAGE} AS base_check
FROM base_check AS build

FROM eclipse-temurin:21-jdk-alpine AS fallback
FROM fallback AS build

WORKDIR /app

COPY gradlew .
COPY build.gradle .
COPY settings.gradle .

COPY gradle/ ./gradle/
COPY configuration/ ./configuration/
COPY src/ ./src/

RUN ./gradlew build -x test

FROM eclipse-temurin:21-jre-alpine

WORKDIR /app

COPY --from=build /app/build/libs/*SNAPSHOT.jar app.jar

EXPOSE 8888

ENTRYPOINT ["java", "-jar", "app.jar"]