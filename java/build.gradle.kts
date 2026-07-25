import org.gradle.api.tasks.testing.logging.TestLogEvent

plugins {
    id("java-library")
    id("maven-publish")
    id("com.diffplug.spotless") version "7.2.1"
}

group = "dev.pcvolkmer.mv64e"
version = "0.4.3" // x-release-please-version

java {
    toolchain {
        languageVersion = JavaLanguageVersion.of(11)
    }
    withJavadocJar()
    withSourcesJar()
}

repositories {
    mavenCentral()
}

dependencies {
    api(libs.jackson.databind)
    implementation(libs.jspecify)

    testImplementation(platform(libs.junit.bom))
    testImplementation(libs.bundles.tests)
}

tasks.test {
    testLogging {
        events = setOf(TestLogEvent.PASSED, TestLogEvent.SKIPPED, TestLogEvent.FAILED)
    }
    useJUnitPlatform()
    dependsOn(tasks.spotlessApply)
}

tasks.withType<JavaCompile> {
    options.encoding = "UTF-8"
}

spotless {
    java {
        importOrder()
        removeUnusedImports()
        googleJavaFormat()
    }
}

publishing {
    repositories {
        mavenLocal()
        maven {
            name = "GitDnpmDev"

            val releasesRepoUrl = uri("https://git.dnpm.dev/api/packages/public/maven")
            val snapshotsRepoUrl = uri("https://git.dnpm.dev/api/packages/public-snapshots/maven")
            url = if (version.toString().endsWith("SNAPSHOT"))
                snapshotsRepoUrl
            else
                releasesRepoUrl

            credentials(HttpHeaderCredentials::class) {
                name = "Authorization"
                value = "token ${properties["dnpm_dev_token"] ?: ""}"
            }

            authentication {
                create<HttpHeaderAuthentication>("header")
            }
        }
    }

    publications {
        create<MavenPublication>("maven") {
            from(components["java"])
        }
    }
}
