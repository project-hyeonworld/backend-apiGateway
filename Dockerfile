FROM eclipse-temurin:21-jre-alpine

WORKDIR /app

COPY gradlew .
COPY build.gradle .
COPY settings.gradle .

COPY gradle/ ./gradle/
COPY src/ ./src/

RUN ./gradlew build -x test

WORKDIR /app

COPY /app/build/libs/*SNAPSHOT.jar app.jar

EXPOSE {APPLICATION_PORT}

ENTRYPOINT ["java", "-jar", "app.jar"]