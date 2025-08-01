plugins {
    id("java")
    id("org.jetbrains.intellij") version "1.17.2"
}

group = "com.rzork"
version = "1.0.3"

repositories {
    mavenCentral()
}

intellij {
    version.set("2023.3.5")
    type.set("IC") 

    plugins.set(listOf(
        "com.intellij.java",
        "org.jetbrains.kotlin",
        "org.jetbrains.plugins.terminal"
    ))
}

dependencies {
    testImplementation("junit:junit:4.13.2")
}

tasks {
    withType<JavaCompile> {
        sourceCompatibility = "17"
        targetCompatibility = "17"
    }

    patchPluginXml {
        sinceBuild.set("233")
        untilBuild.set("241.*")
    }

    signPlugin {
        certificateChain.set(System.getenv("CERTIFICATE_CHAIN"))
        privateKey.set(System.getenv("PRIVATE_KEY"))
        password.set(System.getenv("PRIVATE_KEY_PASSWORD"))
    }

    publishPlugin {
        token.set(System.getenv("PUBLISH_TOKEN"))
    }
} 