ARG BASE_IMAGE={APPLICATION_NAME}:{DOCKER_IMAGE_TAG}
FROM ${BASE_IMAGE} AS build

WORKDIR /app

COPY gradlew .
COPY build.gradle .
COPY settings.gradle .

COPY gradle/ ./gradle/
COPY src/ ./src/

RUN ./gradlew build -x test

FROM eclipse-temurin:21-jre-alpine

WORKDIR /app

COPY --from=build /app/build/libs/*SNAPSHOT.jar app.jar

EXPOSE {APPLICATION_PORT}

ENTRYPOINT ["java", "-jar", "app.jar"]